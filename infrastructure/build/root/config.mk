TERRAFORM_CMD ?= tofu

# Used internally.  Users should pass GOOS and/or GOARCH.
help-vars: # @HELP prints Build Variables
help-vars:
	@echo "VARIABLES:"
	@echo "    TERRAFORM_CMD = $(TERRAFORM_CMD)"

help: # @HELP prints this help message
help: help-vars
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '

.PHONY: help help-vars
