# Infrastructure

This directory contains the infrastructure code and modules for provisioning and managing resources in the capstone
project.

## Prerequisites

- [OpenTofu](https://opentofu.org/docs/intro/install/) installed and connected to cloud provider
- [Docker](https://www.docker.com/get-started) installed and running.
- Appropriate permissions and access configured for infrastructure provisioning (e.g., AWS credentials for ECR and EKS
  modules).

## Amazon Elastic Kubernetes Service

> Amazon Elastic Kubernetes Service (Amazon EKS) is a managed Kubernetes service to run Kubernetes in the AWS cloud and
> on-premises data centers. In the cloud, Amazon EKS automatically manages the availability and scalability of the
> Kubernetes control plane nodes responsible for scheduling containers, managing application availability, storing
> cluster data, and other key tasks.

## Project Structure

```plaintext
.
├── build
│   └── root
│       ├── config.mk               # Configuration variables for Makefile
│       └── Makefile                # Core Makefile for infrastructure tasks
├── main.tofu                       # Main infrastructure configuration file
├── Makefile -> ./build/root/Makefile # Symlink to core Makefile
├── modules
│   ├── ecr
│   │   ├── ecr.tofu                # ECR resource definitions
│   │   ├── outputs.tofu            # ECR output variables
│   │   └── variables.tofu          # ECR input variables
│   └── eks
│       ├── eks.tofu                # EKS cluster definitions
│       ├── iam.tofu                # IAM roles and policies for EKS
│       ├── outputs.tofu            # EKS output variables
│       └── vpc.tofu                # VPC configuration for EKS
└── README.md                       # Project documentation
```

## Workflow

- [Initialize Terraform](#initialize-terraform)
- [Install or Update Dependencies](#install-or-update-dependencies)
- [Format Terraform Files](#format-terraform-files)
- [Validate Terraform Configuration](#validate-terraform-configuration)
- [Plan Infrastructure Changes](#plan-infrastructure-changes)
- [Apply Infrastructure Changes](#apply-infrastructure-changes)
- [Clean Up](#clean-up)

### Initialize Terraform

Initializes the Terraform environment and prepares it for running.

```bash
make infrastructure-init
```

### Install or Update Dependencies

Fetches and updates any required dependencies for the Terraform configuration.

```bash
make infrastructure-deps
```

### Format Terraform Files

Formats the Terraform files according to standard conventions.

```bash
make infrastructure-fmt
```

### Validate Terraform Configuration

Validates the syntax and usage of variables in the Terraform configuration files.

```bash
make infrastructure-validate
```

### Plan Infrastructure Changes

Generates a plan showing any changes to be made to the infrastructure, compared to the existing state.

```bash
make infrastructure-plan
```

### Apply Infrastructure Changes

Applies the changes defined in the Terraform configuration files, provisioning or modifying resources as needed.

```bash
make infrastructure-apply
```

### Clean Up

Cleans up any generated files, such as plan outputs.

```bash
make clean
```

## Makefile Details

The `Makefile` is located in the `build/root` directory and is symlinked to the project root for easy access. It
includes various targets for infrastructure tasks, which are executed using Terraform commands. The following targets
are available:

- `infrastructure-init`: Initializes the Terraform backend.
- `infrastructure-deps`: Downloads or updates module dependencies.
- `infrastructure-fmt`: Formats Terraform files.
- `infrastructure-validate`: Validates Terraform configurations.
- `infrastructure-plan`: Generates a Terraform execution plan.
- `infrastructure-apply`: Applies Terraform configurations.
- `clean`: Removes any generated files from previous runs.
