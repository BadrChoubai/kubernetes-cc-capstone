# Services

This directory contains the source code for the application and
Container/Image artifacts created while building the capstone project.

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed and running.
- Go installed on your local machine (if not using Docker exclusively).
- A configured Docker registry where images will be pushed.

This project uses a `Makefile` to build Go binaries and create Docker images for different platforms. The Makefile
automates the process of building and packaging your applications into containers, making it easier to manage
dependencies and deployment.

## Workflow

- [Run the Build Process](#run-the-build-process)
- [Build Container Images](#build-docker-images)
- [Push Images to Container Registry](#push-images-to-your-registry)

> Set Up Environment Variables:
> 
>  ```shell
>  export OS=linux
>  export ARCH=amd64
>  export TAG=latest
>  export REGISTRY=localhost:5000
> ```

### Run the Build Process

This is the default target and will build all binaries defined in the `BINS` variable for the specified platform.

```shell
make build
```
### Build Docker images:

Builds Docker images for each binary specified in the `BINS` variable. Images are tagged with `$(REGISTRY)/$(BIN):$(TAG)`.

```bash
make image
```

### Push images to your registry:

Pushes the Docker images for the specified binaries to the configured Docker registry.

```bash
make push
```

### Clean up:

Cleans up build artifacts, removing binaries and temporary files.

```bash
make clean
```

#### Directory Structure

- `cmd/`: Contains the main application source code for each binary.
- `build/`: Contains the Dockerfile template (`Dockerfile.in`) used to create the images.
- **Generated Files and Directories**:
   - `.dist/`: Output directory for compiled binaries.
   - `.go/`: Caching directory for Go modules and builds.
   - `.dockerfile/`: Output directory for dynamically generated Dockerfile for building specific binaries.
   - `.image/`: Output directory for marker file indicating successful build of a Docker image for a binary.
