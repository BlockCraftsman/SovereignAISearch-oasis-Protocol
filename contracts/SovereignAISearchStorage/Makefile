#!/usr/bin/env gmake

include ../../common.mk

.PHONY: SovereignAISearchStorage clean

SovereignAISearchStorage:
	@printf "$(CYAN)*** Compiling $(BLUE)$@$(CYAN)...$(OFF)\n"
	@solc --bin --abi --optimize --overwrite -o . $@.sol
	@mv $@.bin $@.hex
	@printf "$(CYAN)*** Generating Go bindings for $(BLUE)$@$(CYAN)...$(OFF)\n"
	@abigen --abi=$@.abi --bin=$@.hex --pkg SovereignAISearchStorage --out=$@.go

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@-rm -f SovereignAISearchStorage.abi SovereignAISearchStorage.hex SovereignAISearchStorage.go

