# Scrappy Ops - GitOps Manifests

This repository defines the desired state of the Scrappy Ops infrastructure and application workloads. It is the single source of truth for the cluster configuration.

## Directory Structure
- /kubernetes: Contains Kubernetes resource definitions.
    - /base: Standard manifests for deployments, services, and networking.
    - /overlays: Environment-specific configurations and patches.
- /crossplane: Defines managed infrastructure and cloud resources.

## Automation
- Tooling: ArgoCD
- Infrastructure Provider: Crossplane
- Sync Strategy: Automated reconciliation of git state to the live cluster.

## Maintenance
Manual changes to the cluster should be avoided. All modifications to the environment must be committed to this repository to be applied by the GitOps controller.
