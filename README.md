# Capstone Project: AWS Elastic Kubernetes Service

> Related Course Notes and Code: [Docker/Kubernetes Course](https://www.github.com/badrchoubai/docker-kubernetes-course)

For the capstone project in the course, we focused on moving away from our local development environment (`minikube`) and start to learn
about the workflow, deployment options, etc. involved in deploying software to a production cloud environment: AWS (Amazon Web
Services).

> ## EKS vs ECS
>
>  | **AWS EKS (Elastic Kubernetes Service)**         | **AWS ECS (Elastic Container Service)** |
>  |:-----------------------------------------------------|:--------------------------------------------|
>  | Managed service for Kubernetes deployments           | Managed service for Container deployments   |
>  | No AWS-specific syntax or philosophy required        | AWS-specific syntax and philosophy applies  |
>  | Use standard Kubernetes configurations and resources | Use AWS-specific configuration and concepts |

**Sections**:

- [Stretch Goals](#stretch-goals)
  - [OpenTofu](#opentofu)

---

## Stretch Goals

For myself, I wanted to add two stretch goals:

1. Build the applications in a different programming language (Go)
2. Use an infrastructure-as-code tool to manage resources deployed to AWS ([OpenTofu](#opentofu))

### Prerequisites for `services` Project

- [Docker](https://www.docker.com/get-started) installed and running
- Go installed on your local machine (if not using Docker exclusively)
- A configured Docker registry where images will be pushed

The project uses a `Makefile` to build Go binaries inside of `services/cmd/` and create Docker images for different platforms. The Makefile
automates the process of building and packaging your applications into containers, making it easier to manage
dependencies and deployment.

[`services` README](./project/services/README.md)

### Prerequisites for `infrastructure` Project

- [OpenTofu](https://opentofu.org/docs/intro/install/) installed and connected to cloud provider

[`infrastructure` README](./project/infrastructure/README.md)

--- 

### OpenTofu

OpenTofu is an open-source fork of Terraform managed by the Linux Foundation. From
the [Manifesto](https://opentofu.org/manifesto/),
it was created in response to Hashicorp's decision to change the license on the Terraform source code from MPL (Mozilla
Public License) to a non-open source license.

The maintainers of OpenTofu and its users believe that the license change ultimately harms the open-source community and
the ecosystem that Terraform developed over the nine years leading up to the change.

> Install OpenTofu on Ubuntu using `snap`:
>
> ```shell
> snap install --classic opentofu 
> ```

