// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SovereignAISearchStorage

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SovereignAISearchStorageSearchResult is an auto generated low-level Go binding around an user-defined struct.
type SovereignAISearchStorageSearchResult struct {
	Query         string
	ResultHashCID string
	Timestamp     *big.Int
	IsEncrypted   bool
	EncryptionKey string
	SearchScore   *big.Int
}

// SovereignAISearchStorageMetaData contains all meta data concerning the SovereignAISearchStorage contract.
var SovereignAISearchStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SearchResultStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UserProfileCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createUserProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getRecentSearches\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"query\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultHashCID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEncrypted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"encryptionKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"searchScore\",\"type\":\"uint256\"}],\"internalType\":\"structSovereignAISearchStorage.SearchResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageNumber\",\"type\":\"uint256\"}],\"name\":\"getSearchHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"query\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultHashCID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEncrypted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"encryptionKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"searchScore\",\"type\":\"uint256\"}],\"internalType\":\"structSovereignAISearchStorage.SearchResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSearches\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSearchTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"keyword\",\"type\":\"string\"}],\"name\":\"searchHistoryByKeyword\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"query\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultHashCID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEncrypted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"encryptionKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"searchScore\",\"type\":\"uint256\"}],\"internalType\":\"structSovereignAISearchStorage.SearchResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"query\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultHashCID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isEncrypted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"encryptionKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"searchScore\",\"type\":\"uint256\"}],\"name\":\"storeSearchResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userProfiles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"searchCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastSearchTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506115c98061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c806353f6ad371161005857806353f6ad371461010457806383f166a214610121578063c765da8414610134578063cc083f4f1461013c575f80fd5b8063332d56d71461007e5780633aef702e146100cf5780634817ee01146100ef575b5f80fd5b6100af61008c36600461107e565b5f60208190529081526040902080546001820154600290920154909160ff169083565b604080519384529115156020840152908201526060015b60405180910390f35b6100e26100dd3660046110ab565b61014f565b6040516100c6919061110e565b6101026100fd36600461126f565b61050a565b005b61010c610664565b604080519283526020830191909152016100c6565b6100e261012f366004611311565b6106b3565b6101026109f9565b6100e261014a366004611328565b610ab1565b335f9081526020819052604090206001015460609060ff1661018c5760405162461bcd60e51b815260040161018390611362565b60405180910390fd5b5f83116101db5760405162461bcd60e51b815260206004820181905260248201527f506167652073697a65206d7573742062652067726561746572207468616e20306044820152606401610183565b335f908152600160205260408120906101f4848661139f565b825490915081106102475760405162461bcd60e51b815260206004820152601860248201527f50616765206e756d626572206f7574206f662072616e676500000000000000006044820152606401610183565b5f61025286836113b6565b8354909150811115610262575081545b5f61026d83836113c9565b90505f8167ffffffffffffffff811115610289576102896111d2565b6040519080825280602002602001820160405280156102c257816020015b6102af611049565b8152602001906001900390816102a75790505b5090505f5b828110156104fc57856102da82876113b6565b815481106102ea576102ea6113dc565b905f5260205f2090600602016040518060c00160405290815f82018054610310906113f0565b80601f016020809104026020016040519081016040528092919081815260200182805461033c906113f0565b80156103875780601f1061035e57610100808354040283529160200191610387565b820191905f5260205f20905b81548152906001019060200180831161036a57829003601f168201915b505050505081526020016001820180546103a0906113f0565b80601f01602080910402602001604051908101604052809291908181526020018280546103cc906113f0565b80156104175780601f106103ee57610100808354040283529160200191610417565b820191905f5260205f20905b8154815290600101906020018083116103fa57829003601f168201915b505050918352505060028201546020820152600382015460ff161515604082015260048201805460609092019161044d906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610479906113f0565b80156104c45780601f1061049b576101008083540402835291602001916104c4565b820191905f5260205f20905b8154815290600101906020018083116104a757829003601f168201915b505050505081526020016005820154815250508282815181106104e9576104e96113dc565b60209081029190910101526001016102c7565b509450505050505b92915050565b335f9081526020819052604090206001015460ff1661053b5760405162461bcd60e51b815260040161018390611362565b6040805160c0810182528681526020808201879052428284015285151560608301526080820185905260a08201849052335f908152600180835293812080549485018155815220815191928392600690910290910190819061059d9082611474565b50602082015160018201906105b29082611474565b5060408201516002820155606082015160038201805460ff1916911515919091179055608082015160048201906105e99082611474565b5060a09190910151600590910155335f90815260208190526040812080549091829061061483611534565b9091555050426002820181905560405133917f2779bd197c7c61762a8b9f56c4c8f29ff0ebe7d27f2cbc3caba9af873631b79891610653918b9161154c565b60405180910390a250505050505050565b335f90815260208190526040812060010154819060ff166106975760405162461bcd60e51b815260040161018390611362565b5050335f90815260208190526040902080546002909101549091565b335f9081526020819052604090206001015460609060ff166106e75760405162461bcd60e51b815260040161018390611362565b5f82116107365760405162461bcd60e51b815260206004820152601c60248201527f436f756e74206d7573742062652067726561746572207468616e2030000000006044820152606401610183565b335f90815260016020526040902080548390811115610753575080545b5f8167ffffffffffffffff81111561076d5761076d6111d2565b6040519080825280602002602001820160405280156107a657816020015b610793611049565b81526020019060019003908161078b5790505b5090505f5b828110156109f0578354849082906107c49086906113c9565b6107ce91906113b6565b815481106107de576107de6113dc565b905f5260205f2090600602016040518060c00160405290815f82018054610804906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610830906113f0565b801561087b5780601f106108525761010080835404028352916020019161087b565b820191905f5260205f20905b81548152906001019060200180831161085e57829003601f168201915b50505050508152602001600182018054610894906113f0565b80601f01602080910402602001604051908101604052809291908181526020018280546108c0906113f0565b801561090b5780601f106108e25761010080835404028352916020019161090b565b820191905f5260205f20905b8154815290600101906020018083116108ee57829003601f168201915b505050918352505060028201546020820152600382015460ff1615156040820152600482018054606090920191610941906113f0565b80601f016020809104026020016040519081016040528092919081815260200182805461096d906113f0565b80156109b85780601f1061098f576101008083540402835291602001916109b8565b820191905f5260205f20905b81548152906001019060200180831161099b57829003601f168201915b505050505081526020016005820154815250508282815181106109dd576109dd6113dc565b60209081029190910101526001016107ab565b50949350505050565b335f9081526020819052604090206001015460ff1615610a515760405162461bcd60e51b81526020600482015260136024820152725573657220616c72656164792065786973747360681b6044820152606401610183565b335f818152602081815260408083206001808201805460ff19169091179055928355426002840181905590519081529192917f558f6a4c7dd60a47054b0f83e1b809be6f15b9cfb9340787599e85f5e2752ef5910160405180910390a250565b335f9081526020819052604090206001015460609060ff16610ae55760405162461bcd60e51b815260040161018390611362565b5f825111610b355760405162461bcd60e51b815260206004820152601760248201527f4b6579776f72642063616e6e6f7420626520656d7074790000000000000000006044820152606401610183565b335f90815260016020526040812090805b8254811015610c1a57610bff838281548110610b6457610b646113dc565b905f5260205f2090600602015f018054610b7d906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ba9906113f0565b8015610bf45780601f10610bcb57610100808354040283529160200191610bf4565b820191905f5260205f20905b815481529060010190602001808311610bd757829003601f168201915b505050505086610f70565b15610c125781610c0e81611534565b9250505b600101610b46565b505f8167ffffffffffffffff811115610c3557610c356111d2565b604051908082528060200260200182016040528015610c6e57816020015b610c5b611049565b815260200190600190039081610c535790505b5090505f805b8454811015610f6557610d2d858281548110610c9257610c926113dc565b905f5260205f2090600602015f018054610cab906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd7906113f0565b8015610d225780601f10610cf957610100808354040283529160200191610d22565b820191905f5260205f20905b815481529060010190602001808311610d0557829003601f168201915b505050505088610f70565b15610f5d57848181548110610d4457610d446113dc565b905f5260205f2090600602016040518060c00160405290815f82018054610d6a906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610d96906113f0565b8015610de15780601f10610db857610100808354040283529160200191610de1565b820191905f5260205f20905b815481529060010190602001808311610dc457829003601f168201915b50505050508152602001600182018054610dfa906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610e26906113f0565b8015610e715780601f10610e4857610100808354040283529160200191610e71565b820191905f5260205f20905b815481529060010190602001808311610e5457829003601f168201915b505050918352505060028201546020820152600382015460ff1615156040820152600482018054606090920191610ea7906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed3906113f0565b8015610f1e5780601f10610ef557610100808354040283529160200191610f1e565b820191905f5260205f20905b815481529060010190602001808311610f0157829003601f168201915b50505050508152602001600582015481525050838381518110610f4357610f436113dc565b60200260200101819052508180610f5990611534565b9250505b600101610c74565b509095945050505050565b80515f90839083901580610f85575080518251105b15610f94575f92505050610504565b5f5b81518351610fa491906113c9565b811161103e5760015f5b835181101561101757838181518110610fc957610fc96113dc565b01602001516001600160f81b03191685610fe383866113b6565b81518110610ff357610ff36113dc565b01602001516001600160f81b0319161461100f575f9150611017565b600101610fae565b50801561102b576001945050505050610504565b508061103681611534565b915050610f96565b505f95945050505050565b6040518060c0016040528060608152602001606081526020015f81526020015f15158152602001606081526020015f81525090565b5f6020828403121561108e575f80fd5b81356001600160a01b03811681146110a4575f80fd5b9392505050565b5f80604083850312156110bc575f80fd5b50508035926020909101359150565b5f81518084525f5b818110156110ef576020818501810151868301820152016110d3565b505f602082860101526020601f19601f83011685010191505092915050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b838110156111c457603f19898403018552815160c0815181865261115b828701826110cb565b915050888201518582038a87015261117382826110cb565b915050878201518886015260608083015115158187015250608080830151868303828801526111a283826110cb565b60a0948501519790940196909652505094870194925090860190600101611135565b509098975050505050505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126111f5575f80fd5b813567ffffffffffffffff80821115611210576112106111d2565b604051601f8301601f19908116603f01168101908282118183101715611238576112386111d2565b81604052838152866020858801011115611250575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a08688031215611283575f80fd5b853567ffffffffffffffff8082111561129a575f80fd5b6112a689838a016111e6565b965060208801359150808211156112bb575f80fd5b6112c789838a016111e6565b95506040880135915081151582146112dd575f80fd5b909350606087013590808211156112f2575f80fd5b506112ff888289016111e6565b95989497509295608001359392505050565b5f60208284031215611321575f80fd5b5035919050565b5f60208284031215611338575f80fd5b813567ffffffffffffffff81111561134e575f80fd5b61135a848285016111e6565b949350505050565b6020808252600f908201526e55736572206e6f742061637469766560881b604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176105045761050461138b565b808201808211156105045761050461138b565b818103818111156105045761050461138b565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061140457607f821691505b60208210810361142257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561146f57805f5260205f20601f840160051c8101602085101561144d5750805b601f840160051c820191505b8181101561146c575f8155600101611459565b50505b505050565b815167ffffffffffffffff81111561148e5761148e6111d2565b6114a28161149c84546113f0565b84611428565b602080601f8311600181146114d5575f84156114be5750858301515b5f19600386901b1c1916600185901b17855561152c565b5f85815260208120601f198616915b82811015611503578886015182559484019460019091019084016114e4565b508582101561152057878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f600182016115455761154561138b565b5060010190565b604081525f61155e60408301856110cb565b9050826020830152939250505056fea2646970667358221220f8557f61c34fdfd7a13c62aad2a9df5a48edbb74e04e9d83d7481dd41654ed2364736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// SovereignAISearchStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SovereignAISearchStorageMetaData.ABI instead.
var SovereignAISearchStorageABI = SovereignAISearchStorageMetaData.ABI

// SovereignAISearchStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SovereignAISearchStorageMetaData.Bin instead.
var SovereignAISearchStorageBin = SovereignAISearchStorageMetaData.Bin

// DeploySovereignAISearchStorage deploys a new Ethereum contract, binding an instance of SovereignAISearchStorage to it.
func DeploySovereignAISearchStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SovereignAISearchStorage, error) {
	parsed, err := SovereignAISearchStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SovereignAISearchStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SovereignAISearchStorage{SovereignAISearchStorageCaller: SovereignAISearchStorageCaller{contract: contract}, SovereignAISearchStorageTransactor: SovereignAISearchStorageTransactor{contract: contract}, SovereignAISearchStorageFilterer: SovereignAISearchStorageFilterer{contract: contract}}, nil
}

// SovereignAISearchStorage is an auto generated Go binding around an Ethereum contract.
type SovereignAISearchStorage struct {
	SovereignAISearchStorageCaller     // Read-only binding to the contract
	SovereignAISearchStorageTransactor // Write-only binding to the contract
	SovereignAISearchStorageFilterer   // Log filterer for contract events
}

// SovereignAISearchStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SovereignAISearchStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SovereignAISearchStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SovereignAISearchStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SovereignAISearchStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SovereignAISearchStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SovereignAISearchStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SovereignAISearchStorageSession struct {
	Contract     *SovereignAISearchStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SovereignAISearchStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SovereignAISearchStorageCallerSession struct {
	Contract *SovereignAISearchStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SovereignAISearchStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SovereignAISearchStorageTransactorSession struct {
	Contract     *SovereignAISearchStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SovereignAISearchStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SovereignAISearchStorageRaw struct {
	Contract *SovereignAISearchStorage // Generic contract binding to access the raw methods on
}

// SovereignAISearchStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SovereignAISearchStorageCallerRaw struct {
	Contract *SovereignAISearchStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SovereignAISearchStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SovereignAISearchStorageTransactorRaw struct {
	Contract *SovereignAISearchStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSovereignAISearchStorage creates a new instance of SovereignAISearchStorage, bound to a specific deployed contract.
func NewSovereignAISearchStorage(address common.Address, backend bind.ContractBackend) (*SovereignAISearchStorage, error) {
	contract, err := bindSovereignAISearchStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SovereignAISearchStorage{SovereignAISearchStorageCaller: SovereignAISearchStorageCaller{contract: contract}, SovereignAISearchStorageTransactor: SovereignAISearchStorageTransactor{contract: contract}, SovereignAISearchStorageFilterer: SovereignAISearchStorageFilterer{contract: contract}}, nil
}

// NewSovereignAISearchStorageCaller creates a new read-only instance of SovereignAISearchStorage, bound to a specific deployed contract.
func NewSovereignAISearchStorageCaller(address common.Address, caller bind.ContractCaller) (*SovereignAISearchStorageCaller, error) {
	contract, err := bindSovereignAISearchStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SovereignAISearchStorageCaller{contract: contract}, nil
}

// NewSovereignAISearchStorageTransactor creates a new write-only instance of SovereignAISearchStorage, bound to a specific deployed contract.
func NewSovereignAISearchStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SovereignAISearchStorageTransactor, error) {
	contract, err := bindSovereignAISearchStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SovereignAISearchStorageTransactor{contract: contract}, nil
}

// NewSovereignAISearchStorageFilterer creates a new log filterer instance of SovereignAISearchStorage, bound to a specific deployed contract.
func NewSovereignAISearchStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SovereignAISearchStorageFilterer, error) {
	contract, err := bindSovereignAISearchStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SovereignAISearchStorageFilterer{contract: contract}, nil
}

// bindSovereignAISearchStorage binds a generic wrapper to an already deployed contract.
func bindSovereignAISearchStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SovereignAISearchStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SovereignAISearchStorage *SovereignAISearchStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SovereignAISearchStorage.Contract.SovereignAISearchStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SovereignAISearchStorage *SovereignAISearchStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.SovereignAISearchStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SovereignAISearchStorage *SovereignAISearchStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.SovereignAISearchStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SovereignAISearchStorage *SovereignAISearchStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SovereignAISearchStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SovereignAISearchStorage *SovereignAISearchStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SovereignAISearchStorage *SovereignAISearchStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.contract.Transact(opts, method, params...)
}

// GetRecentSearches is a free data retrieval call binding the contract method 0x83f166a2.
//
// Solidity: function getRecentSearches(uint256 count) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageCaller) GetRecentSearches(opts *bind.CallOpts, count *big.Int) ([]SovereignAISearchStorageSearchResult, error) {
	var out []interface{}
	err := _SovereignAISearchStorage.contract.Call(opts, &out, "getRecentSearches", count)

	if err != nil {
		return *new([]SovereignAISearchStorageSearchResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]SovereignAISearchStorageSearchResult)).(*[]SovereignAISearchStorageSearchResult)

	return out0, err

}

// GetRecentSearches is a free data retrieval call binding the contract method 0x83f166a2.
//
// Solidity: function getRecentSearches(uint256 count) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) GetRecentSearches(count *big.Int) ([]SovereignAISearchStorageSearchResult, error) {
	return _SovereignAISearchStorage.Contract.GetRecentSearches(&_SovereignAISearchStorage.CallOpts, count)
}

// GetRecentSearches is a free data retrieval call binding the contract method 0x83f166a2.
//
// Solidity: function getRecentSearches(uint256 count) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageCallerSession) GetRecentSearches(count *big.Int) ([]SovereignAISearchStorageSearchResult, error) {
	return _SovereignAISearchStorage.Contract.GetRecentSearches(&_SovereignAISearchStorage.CallOpts, count)
}

// GetSearchHistory is a free data retrieval call binding the contract method 0x3aef702e.
//
// Solidity: function getSearchHistory(uint256 pageSize, uint256 pageNumber) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageCaller) GetSearchHistory(opts *bind.CallOpts, pageSize *big.Int, pageNumber *big.Int) ([]SovereignAISearchStorageSearchResult, error) {
	var out []interface{}
	err := _SovereignAISearchStorage.contract.Call(opts, &out, "getSearchHistory", pageSize, pageNumber)

	if err != nil {
		return *new([]SovereignAISearchStorageSearchResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]SovereignAISearchStorageSearchResult)).(*[]SovereignAISearchStorageSearchResult)

	return out0, err

}

// GetSearchHistory is a free data retrieval call binding the contract method 0x3aef702e.
//
// Solidity: function getSearchHistory(uint256 pageSize, uint256 pageNumber) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) GetSearchHistory(pageSize *big.Int, pageNumber *big.Int) ([]SovereignAISearchStorageSearchResult, error) {
	return _SovereignAISearchStorage.Contract.GetSearchHistory(&_SovereignAISearchStorage.CallOpts, pageSize, pageNumber)
}

// GetSearchHistory is a free data retrieval call binding the contract method 0x3aef702e.
//
// Solidity: function getSearchHistory(uint256 pageSize, uint256 pageNumber) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageCallerSession) GetSearchHistory(pageSize *big.Int, pageNumber *big.Int) ([]SovereignAISearchStorageSearchResult, error) {
	return _SovereignAISearchStorage.Contract.GetSearchHistory(&_SovereignAISearchStorage.CallOpts, pageSize, pageNumber)
}

// GetUserStats is a free data retrieval call binding the contract method 0x53f6ad37.
//
// Solidity: function getUserStats() view returns(uint256 totalSearches, uint256 lastSearchTime)
func (_SovereignAISearchStorage *SovereignAISearchStorageCaller) GetUserStats(opts *bind.CallOpts) (struct {
	TotalSearches  *big.Int
	LastSearchTime *big.Int
}, error) {
	var out []interface{}
	err := _SovereignAISearchStorage.contract.Call(opts, &out, "getUserStats")

	outstruct := new(struct {
		TotalSearches  *big.Int
		LastSearchTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSearches = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastSearchTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserStats is a free data retrieval call binding the contract method 0x53f6ad37.
//
// Solidity: function getUserStats() view returns(uint256 totalSearches, uint256 lastSearchTime)
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) GetUserStats() (struct {
	TotalSearches  *big.Int
	LastSearchTime *big.Int
}, error) {
	return _SovereignAISearchStorage.Contract.GetUserStats(&_SovereignAISearchStorage.CallOpts)
}

// GetUserStats is a free data retrieval call binding the contract method 0x53f6ad37.
//
// Solidity: function getUserStats() view returns(uint256 totalSearches, uint256 lastSearchTime)
func (_SovereignAISearchStorage *SovereignAISearchStorageCallerSession) GetUserStats() (struct {
	TotalSearches  *big.Int
	LastSearchTime *big.Int
}, error) {
	return _SovereignAISearchStorage.Contract.GetUserStats(&_SovereignAISearchStorage.CallOpts)
}

// SearchHistoryByKeyword is a free data retrieval call binding the contract method 0xcc083f4f.
//
// Solidity: function searchHistoryByKeyword(string keyword) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageCaller) SearchHistoryByKeyword(opts *bind.CallOpts, keyword string) ([]SovereignAISearchStorageSearchResult, error) {
	var out []interface{}
	err := _SovereignAISearchStorage.contract.Call(opts, &out, "searchHistoryByKeyword", keyword)

	if err != nil {
		return *new([]SovereignAISearchStorageSearchResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]SovereignAISearchStorageSearchResult)).(*[]SovereignAISearchStorageSearchResult)

	return out0, err

}

// SearchHistoryByKeyword is a free data retrieval call binding the contract method 0xcc083f4f.
//
// Solidity: function searchHistoryByKeyword(string keyword) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) SearchHistoryByKeyword(keyword string) ([]SovereignAISearchStorageSearchResult, error) {
	return _SovereignAISearchStorage.Contract.SearchHistoryByKeyword(&_SovereignAISearchStorage.CallOpts, keyword)
}

// SearchHistoryByKeyword is a free data retrieval call binding the contract method 0xcc083f4f.
//
// Solidity: function searchHistoryByKeyword(string keyword) view returns((string,string,uint256,bool,string,uint256)[])
func (_SovereignAISearchStorage *SovereignAISearchStorageCallerSession) SearchHistoryByKeyword(keyword string) ([]SovereignAISearchStorageSearchResult, error) {
	return _SovereignAISearchStorage.Contract.SearchHistoryByKeyword(&_SovereignAISearchStorage.CallOpts, keyword)
}

// UserProfiles is a free data retrieval call binding the contract method 0x332d56d7.
//
// Solidity: function userProfiles(address ) view returns(uint256 searchCount, bool isActive, uint256 lastSearchTime)
func (_SovereignAISearchStorage *SovereignAISearchStorageCaller) UserProfiles(opts *bind.CallOpts, arg0 common.Address) (struct {
	SearchCount    *big.Int
	IsActive       bool
	LastSearchTime *big.Int
}, error) {
	var out []interface{}
	err := _SovereignAISearchStorage.contract.Call(opts, &out, "userProfiles", arg0)

	outstruct := new(struct {
		SearchCount    *big.Int
		IsActive       bool
		LastSearchTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SearchCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.LastSearchTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserProfiles is a free data retrieval call binding the contract method 0x332d56d7.
//
// Solidity: function userProfiles(address ) view returns(uint256 searchCount, bool isActive, uint256 lastSearchTime)
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) UserProfiles(arg0 common.Address) (struct {
	SearchCount    *big.Int
	IsActive       bool
	LastSearchTime *big.Int
}, error) {
	return _SovereignAISearchStorage.Contract.UserProfiles(&_SovereignAISearchStorage.CallOpts, arg0)
}

// UserProfiles is a free data retrieval call binding the contract method 0x332d56d7.
//
// Solidity: function userProfiles(address ) view returns(uint256 searchCount, bool isActive, uint256 lastSearchTime)
func (_SovereignAISearchStorage *SovereignAISearchStorageCallerSession) UserProfiles(arg0 common.Address) (struct {
	SearchCount    *big.Int
	IsActive       bool
	LastSearchTime *big.Int
}, error) {
	return _SovereignAISearchStorage.Contract.UserProfiles(&_SovereignAISearchStorage.CallOpts, arg0)
}

// CreateUserProfile is a paid mutator transaction binding the contract method 0xc765da84.
//
// Solidity: function createUserProfile() returns()
func (_SovereignAISearchStorage *SovereignAISearchStorageTransactor) CreateUserProfile(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SovereignAISearchStorage.contract.Transact(opts, "createUserProfile")
}

// CreateUserProfile is a paid mutator transaction binding the contract method 0xc765da84.
//
// Solidity: function createUserProfile() returns()
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) CreateUserProfile() (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.CreateUserProfile(&_SovereignAISearchStorage.TransactOpts)
}

// CreateUserProfile is a paid mutator transaction binding the contract method 0xc765da84.
//
// Solidity: function createUserProfile() returns()
func (_SovereignAISearchStorage *SovereignAISearchStorageTransactorSession) CreateUserProfile() (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.CreateUserProfile(&_SovereignAISearchStorage.TransactOpts)
}

// StoreSearchResult is a paid mutator transaction binding the contract method 0x4817ee01.
//
// Solidity: function storeSearchResult(string query, string resultHashCID, bool isEncrypted, string encryptionKey, uint256 searchScore) returns()
func (_SovereignAISearchStorage *SovereignAISearchStorageTransactor) StoreSearchResult(opts *bind.TransactOpts, query string, resultHashCID string, isEncrypted bool, encryptionKey string, searchScore *big.Int) (*types.Transaction, error) {
	return _SovereignAISearchStorage.contract.Transact(opts, "storeSearchResult", query, resultHashCID, isEncrypted, encryptionKey, searchScore)
}

// StoreSearchResult is a paid mutator transaction binding the contract method 0x4817ee01.
//
// Solidity: function storeSearchResult(string query, string resultHashCID, bool isEncrypted, string encryptionKey, uint256 searchScore) returns()
func (_SovereignAISearchStorage *SovereignAISearchStorageSession) StoreSearchResult(query string, resultHashCID string, isEncrypted bool, encryptionKey string, searchScore *big.Int) (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.StoreSearchResult(&_SovereignAISearchStorage.TransactOpts, query, resultHashCID, isEncrypted, encryptionKey, searchScore)
}

// StoreSearchResult is a paid mutator transaction binding the contract method 0x4817ee01.
//
// Solidity: function storeSearchResult(string query, string resultHashCID, bool isEncrypted, string encryptionKey, uint256 searchScore) returns()
func (_SovereignAISearchStorage *SovereignAISearchStorageTransactorSession) StoreSearchResult(query string, resultHashCID string, isEncrypted bool, encryptionKey string, searchScore *big.Int) (*types.Transaction, error) {
	return _SovereignAISearchStorage.Contract.StoreSearchResult(&_SovereignAISearchStorage.TransactOpts, query, resultHashCID, isEncrypted, encryptionKey, searchScore)
}

// SovereignAISearchStorageSearchResultStoredIterator is returned from FilterSearchResultStored and is used to iterate over the raw logs and unpacked data for SearchResultStored events raised by the SovereignAISearchStorage contract.
type SovereignAISearchStorageSearchResultStoredIterator struct {
	Event *SovereignAISearchStorageSearchResultStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SovereignAISearchStorageSearchResultStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SovereignAISearchStorageSearchResultStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SovereignAISearchStorageSearchResultStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SovereignAISearchStorageSearchResultStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SovereignAISearchStorageSearchResultStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SovereignAISearchStorageSearchResultStored represents a SearchResultStored event raised by the SovereignAISearchStorage contract.
type SovereignAISearchStorageSearchResultStored struct {
	User      common.Address
	Query     string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSearchResultStored is a free log retrieval operation binding the contract event 0x2779bd197c7c61762a8b9f56c4c8f29ff0ebe7d27f2cbc3caba9af873631b798.
//
// Solidity: event SearchResultStored(address indexed user, string query, uint256 timestamp)
func (_SovereignAISearchStorage *SovereignAISearchStorageFilterer) FilterSearchResultStored(opts *bind.FilterOpts, user []common.Address) (*SovereignAISearchStorageSearchResultStoredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SovereignAISearchStorage.contract.FilterLogs(opts, "SearchResultStored", userRule)
	if err != nil {
		return nil, err
	}
	return &SovereignAISearchStorageSearchResultStoredIterator{contract: _SovereignAISearchStorage.contract, event: "SearchResultStored", logs: logs, sub: sub}, nil
}

// WatchSearchResultStored is a free log subscription operation binding the contract event 0x2779bd197c7c61762a8b9f56c4c8f29ff0ebe7d27f2cbc3caba9af873631b798.
//
// Solidity: event SearchResultStored(address indexed user, string query, uint256 timestamp)
func (_SovereignAISearchStorage *SovereignAISearchStorageFilterer) WatchSearchResultStored(opts *bind.WatchOpts, sink chan<- *SovereignAISearchStorageSearchResultStored, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SovereignAISearchStorage.contract.WatchLogs(opts, "SearchResultStored", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SovereignAISearchStorageSearchResultStored)
				if err := _SovereignAISearchStorage.contract.UnpackLog(event, "SearchResultStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSearchResultStored is a log parse operation binding the contract event 0x2779bd197c7c61762a8b9f56c4c8f29ff0ebe7d27f2cbc3caba9af873631b798.
//
// Solidity: event SearchResultStored(address indexed user, string query, uint256 timestamp)
func (_SovereignAISearchStorage *SovereignAISearchStorageFilterer) ParseSearchResultStored(log types.Log) (*SovereignAISearchStorageSearchResultStored, error) {
	event := new(SovereignAISearchStorageSearchResultStored)
	if err := _SovereignAISearchStorage.contract.UnpackLog(event, "SearchResultStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SovereignAISearchStorageUserProfileCreatedIterator is returned from FilterUserProfileCreated and is used to iterate over the raw logs and unpacked data for UserProfileCreated events raised by the SovereignAISearchStorage contract.
type SovereignAISearchStorageUserProfileCreatedIterator struct {
	Event *SovereignAISearchStorageUserProfileCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SovereignAISearchStorageUserProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SovereignAISearchStorageUserProfileCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SovereignAISearchStorageUserProfileCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SovereignAISearchStorageUserProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SovereignAISearchStorageUserProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SovereignAISearchStorageUserProfileCreated represents a UserProfileCreated event raised by the SovereignAISearchStorage contract.
type SovereignAISearchStorageUserProfileCreated struct {
	User      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserProfileCreated is a free log retrieval operation binding the contract event 0x558f6a4c7dd60a47054b0f83e1b809be6f15b9cfb9340787599e85f5e2752ef5.
//
// Solidity: event UserProfileCreated(address indexed user, uint256 timestamp)
func (_SovereignAISearchStorage *SovereignAISearchStorageFilterer) FilterUserProfileCreated(opts *bind.FilterOpts, user []common.Address) (*SovereignAISearchStorageUserProfileCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SovereignAISearchStorage.contract.FilterLogs(opts, "UserProfileCreated", userRule)
	if err != nil {
		return nil, err
	}
	return &SovereignAISearchStorageUserProfileCreatedIterator{contract: _SovereignAISearchStorage.contract, event: "UserProfileCreated", logs: logs, sub: sub}, nil
}

// WatchUserProfileCreated is a free log subscription operation binding the contract event 0x558f6a4c7dd60a47054b0f83e1b809be6f15b9cfb9340787599e85f5e2752ef5.
//
// Solidity: event UserProfileCreated(address indexed user, uint256 timestamp)
func (_SovereignAISearchStorage *SovereignAISearchStorageFilterer) WatchUserProfileCreated(opts *bind.WatchOpts, sink chan<- *SovereignAISearchStorageUserProfileCreated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SovereignAISearchStorage.contract.WatchLogs(opts, "UserProfileCreated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SovereignAISearchStorageUserProfileCreated)
				if err := _SovereignAISearchStorage.contract.UnpackLog(event, "UserProfileCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserProfileCreated is a log parse operation binding the contract event 0x558f6a4c7dd60a47054b0f83e1b809be6f15b9cfb9340787599e85f5e2752ef5.
//
// Solidity: event UserProfileCreated(address indexed user, uint256 timestamp)
func (_SovereignAISearchStorage *SovereignAISearchStorageFilterer) ParseUserProfileCreated(log types.Log) (*SovereignAISearchStorageUserProfileCreated, error) {
	event := new(SovereignAISearchStorageUserProfileCreated)
	if err := _SovereignAISearchStorage.contract.UnpackLog(event, "UserProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
