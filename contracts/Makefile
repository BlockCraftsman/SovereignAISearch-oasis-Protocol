#!/usr/bin/env gmake

include ../common.mk

.PHONY: all SovereignAISearchStorage clean

all: SovereignAISearchStorage
	@printf "$(CYAN)*** Contracts built successfully!$(OFF)\n"

SovereignAISearchStorage:
	@printf "$(CYAN)*** Building contract: $(BLUE)SovereignAISearchStorage$(CYAN)...$(OFF)\n"
	@$(MAKE) -C SovereignAISearchStorage

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@$(MAKE) -C SovereignAISearchStorage clean

