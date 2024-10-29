package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"aiquery/api/model"
	"aiquery/api/reqmodel"
	"aiquery/api/shared"
	"aiquery/llm"
	"aiquery/rag"
	"aiquery/rag/searxng"
	"aiquery/utils"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func Search(c *gin.Context) {
	var query reqmodel.Query
	if err := c.BindJSON(&query); err != nil {
		utils.GinErrorMsg(c, err)
		return
	}
	query = query.Regularize()
	if query.Question == "" {
		c.String(200, "")
		return
	}

	ans := model.Answer{
		Question: query.Question,
	}

	searched := ""

	// Create a SearxNGConfig instance
	config := searxng.SearxNGConfig{
		Endpoint: "http://127.0.0.1:18080/search", // Replace with your actual endpoint
	}

	queryStr := query.Question

	searchRaw, err := searxng.SearchRawFromSearxNG(config, query.Country(), query.Locale(), queryStr)
	if err != nil {
		log.Println("SearchRawFromSearxNG error:", err)
	} else {
		serpResults := ConvertSearxNGResultsToSERPResults(searchRaw.SpiderResults)
		searched = rag.SearchResultsToText(serpResults)
	}

	ragRaw := ConvertSearxNGRawItemToRagRawItem(searchRaw)

	meta := reqmodel.NewMeta(ragRaw)
	meta.ID = ans.ID
	c.String(200, "%s\r\n", meta.Json())
	ans.Evidence = utils.Json(meta.Evidences)

	if query.ProgLang == "" {
		query.ProgLang = "golang"
	}
	guideStr := query.ProgLang

	if str, err := GenerateProgramGuideStr(query.ProgLang); err == nil {
		guideStr = str
	}

	content := llm.Promptc(query.Language, query.Field, query.Question, guideStr, searched)

	req := utils.ModelGPT35Request([]openai.ChatCompletionMessage{
		utils.ChatMsgFromSystem(content),
		utils.ChatMsgFromUser(query.Question),
	})

	ans.FirstAnswer = chatStreamToGin(c, req)
	ans.IsOk = true

	//store ai search infos to web3 storage

	// create json from data
	data := map[string]interface{}{
		"Question":    query.Question,
		"FirstAnswer": ans.FirstAnswer,
		"RagRaw":      ragRaw,
		"IsOk":        ans.IsOk,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal JSON data:", err)
		return
	}

	// save data to json
	filePath := fmt.Sprintf("/tmp/search_results_%s.json", ans.ID)
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Println("Failed to write JSON data to file:", err)
		return
	}

	// us w3 up to upload file
	cmd := exec.Command("w3", "up", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Failed to upload file using w3 up:", err)
		log.Println("Output:", string(output))
		return
	}

	//
	uploadURL := extractUploadURL(string(output))
	if uploadURL == "" {
		log.Println("Failed to extract upload URL from output:", string(output))
		return
	}
	log.Println("File uploaded successfully. URL:", uploadURL)

	// Extract the CID from the upload URL
	cid := extractCID(uploadURL)
	if cid == "" {
		log.Println("Failed to extract CID from upload URL:", uploadURL)
		return
	}

	// TODO submit blockchain
	// Get contract address and network from environment variables
	contractAddr := os.Getenv("CONTRACT_ADDRESS")
	network := os.Getenv("NETWORK")
	sovereignProtocolPath := os.Getenv("SOVEREIGN_PROTOCOL_PATH")
	if contractAddr == "" || network == "" || sovereignProtocolPath == "" {
		log.Println("Contract address, network, or sovereign protocol path not set in environment variables")
		return
	}

	// Store the CID information inside the deployed contract.
	storeCIDCmd := exec.Command(sovereignProtocolPath, "storeCID", "--network", network, contractAddr, query.Question, cid, "false", "", "100")
	storeCIDOutput, err := storeCIDCmd.CombinedOutput()
	if err != nil {
		log.Println("Failed to store CID in contract:", err)
		log.Println("Output:", string(storeCIDOutput))
		return
	}
	log.Println("CID stored successfully in contract")

	webStorageInfo := fmt.Sprintf(`Web3 Storage URL:%s
`, uploadURL)

	ans.FirstAnswer += webStorageInfo

	c.String(200, webStorageInfo)

}

func extractCID(uploadURL string) string {
	// 使用正则表达式匹配CID
	re := regexp.MustCompile(`/ipfs/([a-zA-Z0-9]+)$`)
	matches := re.FindStringSubmatch(uploadURL)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractUploadURL(output string) string {
	// 使用正则表达式匹配链接地址
	re := regexp.MustCompile(`https://w3s\.link/ipfs/[a-zA-Z0-9]+`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func chatStreamToGin(c *gin.Context, req openai.ChatCompletionRequest) (completeAns string) {
	var resp *openai.ChatCompletionStream

	resp, err := shared.Cli.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		utils.GinErrorMsgTxt(c, "LLM backend broken")
		log.Println(err)
		return
	}
	defer resp.Close()

	sb := strings.Builder{}
	buf := strings.Builder{}

	c.Stream(func(w io.Writer) bool {
		msg, err := resp.Recv()
		if errors.Is(err, io.EOF) || err != nil {
			bufS := buf.String()
			if bufS != "" {
				w.Write([]byte(bufS))
				sb.WriteString(bufS)
			}
			return false
		}
		delta := msg.Choices[0].Delta.Content
		delta = utils.CleanStr(delta)
		sb.WriteString(delta)

		if utils.WriteSplitByRune(w, &buf, delta, '\n', '.', ';', '。', '？', '?') {
			return true
		}
		return true
	})
	completeAns = sb.String()
	return
}
