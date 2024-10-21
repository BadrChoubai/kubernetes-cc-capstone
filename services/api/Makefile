####################################################################################################
# This Makefile contains all targets related to the API documentation
####################################################################################################

# Directories
SERVICES_DIRECTORY = ../internal/services
SERVICES =  auth \
			users

SWAG_CMD = swag init \
	--instanceName $(1) \
	--output $(CURDIR)/$(1) \
	--outputTypes=yaml,json \
	-g http_routes.go

# Rule to generate documentation if swagger.yaml is missing
define SERVICE_DOCS_RULE
$(CURDIR)/$(1)/swagger.yaml: $(SERVICES_DIRECTORY)/$(1)/http_routes.go
	@echo "Generating Swagger docs for $1..."
	cd $(SERVICES_DIRECTORY)/$(1) && $(SWAG_CMD)
endef

define DOCS_CLEAN_RULE
clean/$(1):
	@echo "removing $(1) documentation"
	rm -rf $(1)
endef

# Apply the rule for each service
$(foreach service,$(SERVICES),$(eval $(call SERVICE_DOCS_RULE,$(service))))
$(foreach service,$(SERVICES),$(eval $(call DOCS_CLEAN_RULE,$(service))))

# Generate documentation for all services
all-docs: # @HELP Generate all API documentation
all-docs: $(foreach service,$(SERVICES),$(CURDIR)/$(service)/swagger.yaml)

clean: # @HELP clean up artifacts from all-docs
clean: $(foreach service,$(SERVICES),clean/$(service))

all: # @HELP default job
all: all-docs

####################################################################################################
# HELP
####################################################################################################

help-vars: # @HELP prints build variables
help-vars:
	@echo "  SERVICES = $(SERVICES)"

help: help-vars
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '

SHELL := /usr/bin/env bash -o errexit -o pipefail
.DEFAULT_GOAL = all
.PHONY: all all-docs clean lint
