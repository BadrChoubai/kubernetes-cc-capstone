include build/root/config.mk

infrastructure-init: # @HELP initialize infrastructure-as-code tool
infrastructure-init:
	$(TERRAFORM_CMD) init


infrastructure-deps: # @HELP install or update infrastructure-as-code tool dependencies
infrastructure-deps:
	$(TERRAFORM_CMD) get


infrastructure-fmt: # @HELP apply standard formatting for all infrastructure-as-code files
infrastructure-fmt:
	$(TERRAFORM_CMD) fmt


infrastructure-validate: # @HELP validate infrastructure-as-code files for correct usage of variables, syntax, etc.
infrastructure-validate:
	$(TERRAFORM_CMD) validate


infrastructure-plan: # @HELP compare existing infrastructure to infrastructure-as-code plan
infrastructure-plan:
	$(TERRAFORM_CMD) plan -out=tfplan | tee plan_output.txt
	@if grep -q "No changes" plan_output.txt; then \
		echo "No changes detected."; \
		touch no_changes; \
	else \
		rm -rf no_changes; \
	fi

infrastructure-apply: # @HELP validate configuration, if plan can be applied, apply infrastructure-as-code configuration
infrastructure-apply: infrastructure-validate infrastructure-plan
	@if [ -f no_changes ]; then \
		echo "Skipping apply due to no changes."; \
		exit 0; \
	else \
		$(TERRAFORM_CMD) apply tfplan; \
	fi

all: # @HELP
all: infrastructure-apply

clean: # @HELP clean up files
clean:
	rm plan_output.txt

SHELL := /usr/bin/env bash -o errexit -o pipefail
.DEFAULT_GOAL = all
.PHONY: all