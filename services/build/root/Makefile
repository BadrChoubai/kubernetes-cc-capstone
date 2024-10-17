include build/root/config.mk

BINS ?= users auth
IMAGE_DOTFILES = $(foreach bin,$(BINS),.image-$(bin)-$(TAG))

BUILD_DIRS := .dist/$(OS)_$(ARCH)	\
			  .images				\
			  .dockerfile			\
              .go/cache            	\
              .go/pkg

all: # @HELP builds binaries for one platform ($OS/$ARCH)
all: build

build: # @HELP for each BIN, build binary and place into /bin/$(OS)_$(ARCH)
build: $(BINS)

# Build each binary using Docker
$(BINS): $(BUILD_DIRS)
	docker run                                                  \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
		-v $$(pwd)/.dist/$(OS)_$(ARCH):/.dist                   \
		-v $$(pwd)/.go/cache:/.cache								\
	    --env GOCACHE="/.cache/gocache"                         \
	    --env GOMODCACHE="/.cache/gomodcache"                   \
	    --env ARCH="$(ARCH)"                                    \
	    --env OS="$(OS)"                                        \
	    --env VERSION="$(VERSION)"                              \
	    --env DEBUG="$(DBG)"                                    \
	    --env GOFLAGS="$(GOFLAGS)"                              \
	    --env HTTP_PROXY="$(HTTP_PROXY)"                        \
	    --env HTTPS_PROXY="$(HTTPS_PROXY)"                      \
	    $(CONTAINER_IMAGE)                                      \
	    go build -o .dist/$(OS)_$(ARCH)/"$@" ./cmd/"$@"

$(BUILD_DIRS):
	mkdir -p $@

clean: # @HELP clean artifacts from build
clean:
	rm -rf .go .dist .images .dockerfile

image images: # Builds container images
image images: $(IMAGE_DOTFILES)
	for bin in $(BINS); do \
	    echo "image: $(REGISTRY)/$$bin:$(TAG)";  \
	done
	echo

# Pattern rule to create .image files based on each BIN
.image-%-$(TAG): BIN = $(patsubst .image-%-$(TAG),%,$@)
.image-%-$(TAG):
	echo $(BIN)
	sed                                            \
		-e 's|{ARG_BIN}|$(BIN)$(BIN_EXTENSION)|g'  \
		-e 's|{ARG_ARCH}|$(ARCH)|g'                \
		-e 's|{ARG_OS}|$(OS)|g'                    \
		-e 's|{ARG_FROM}|$(CONTAINER_IMAGE)|g'     \
		./build/Dockerfile.in > ./.dockerfile/.dockerfile-$(BIN)-$(OS)_$(ARCH)
	docker build \
		--progress=plain \
		-t $(REGISTRY)/$(BIN):$(TAG) \
	    -f ./.dockerfile/.dockerfile-$(BIN)-$(OS)_$(ARCH) \
		.
	docker images -q $(REGISTRY)/$(BIN):$(TAG) > ./.images/$@
	echo

push: # @HELP pushes the image for one platform ($OS/$ARCH) to the defined registry
push: image
	for bin in $(BINS); do                     \
	    docker push $(REGISTRY)/$$bin:$(TAG);  \
	done
	echo

lint: # HELP runs golangci-lint linters
lint:
	golangci-lint run

SHELL := /usr/bin/env bash -o errexit -o pipefail
.DEFAULT_GOAL = all
.PHONY: all clean