# Contributing to KubeVirtBMC

Thank you for helping improve KubeVirtBMC! This document distills what new contributors need in order to be productive, based on the conventions we share with the wider [KubeVirt](https://kubevirt.io/) ecosystem.

- [Ways to Contribute](#ways-to-contribute)
- [Development Environment](#development-environment)
- [Project Components](#project-components)
- [Daily Development Workflow](#daily-development-workflow)
- [Getting KubeVirtBMC Running End-to-End](#getting-kubevirtbmc-running-end-to-end)
- [Pull Request Requirements](#pull-request-requirements)
- [Resources & Support](#resources--support)

## Ways to Contribute

- **Issues first**: Every code or documentation change must reference an existing GitHub issue. Open one if it does not exist yet, and describe the user impact plus a proposed approach.
- **Documentation**: Improvements to README sections, Helm chart values, or this file are welcome—especially when they make the developer story clearer.
- **Code changes**: Bug fixes, new features, refactors, and test coverage. Please start discussions in the relevant issue before implementing larger changes.

We follow the [Kubernetes Community Code of Conduct](https://kubernetes.io/community/code-of-conduct/).

## Development Environment

### Prerequisites

| Tool | Minimum version | Notes |
| --- | --- | --- |
| Go | 1.20 | Matches `go.mod`; install via your OS package manager or `golang.org/dl` |
| Docker or Podman | 23.x | Needed for image builds (`make docker-build`) |
| kubectl | 1.27+ | Should match the target cluster |
| kind | 0.19+ | Only required if you run the supplied kind-based environment |
| Helm | 3.12+ | Handy for installing cert-manager |
| Access to a Kubernetes cluster with KubeVirt | v1.2.0 API | kind/Minikube/real clusters are all fine |

> Tip: `make help` lists all supported targets with short descriptions.

### 1. Clone and bootstrap the repo

```sh
mkdir -p "$GOPATH/src/github.com/starbops"
cd "$GOPATH/src/github.com/starbops"
git clone https://github.com/starbops/kubevirtbmc.git
cd kubevirtbmc
make localbin  # optional, but creates ./bin for helper tools
```

### 2. Create a minimal Kubernetes + KubeVirt environment

The quickest path is kind, and the repository already ships a config tailored for KubeVirt e2e tests. The `e2e-setup` target installs the `kind` binary (via the `kind` prerequisite) and boots a cluster using `test/e2e/kind-config.yaml`.

```sh
make e2e-setup
```

Install KubeVirt using the operator flow documented in the official installation guide for Kubernetes clusters. This pulls the latest stable release and waits for it to become `Available`. [^kubevirt-install]

```sh
export RELEASE=$(curl -L https://storage.googleapis.com/kubevirt-prow/release/kubevirt/kubevirt/stable.txt)
kubectl apply -f https://github.com/kubevirt/kubevirt/releases/download/${RELEASE}/kubevirt-operator.yaml
kubectl apply -f https://github.com/kubevirt/kubevirt/releases/download/${RELEASE}/kubevirt-cr.yaml
kubectl -n kubevirt wait kv kubevirt --for=condition=Available --timeout=10m
```

Install cert-manager with the default static manifest as described in the upstream docs. This single manifest bundles the CRDs, webhook, and controller manager. [^cert-manager-install]

```sh
export CERT_MANAGER_VERSION=v1.17.1
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/${CERT_MANAGER_VERSION}/cert-manager.yaml
kubectl -n cert-manager wait --for=condition=Available deployment/cert-manager-webhook --timeout=5m
```

### 3. Optional: load KubeVirt CRDs for integration tests

When running unit tests with envtest, the repo can generate the necessary KubeVirt CRDs automatically:

```sh
make generate-kubevirt-crd
```

## Project Components

KubeVirtBMC ships two cooperating binaries:

- **`virtbmc-controller`** (`cmd/controller/`): A Kubebuilder-based controller that watches `VirtualMachine`, `VirtualMachineInstance`, `Service`, and the custom `VirtualMachineBMC` resources. It creates/updates `virtbmc` pods and services per VM.
- **`virtbmc` agent** (`cmd/virtbmc/`): An IPMI and Redfish emulator that translates management commands into Kubernetes API actions on the target VM.

The Helm chart (`deploy/charts/kubevirtbmc`) deploys the controller, the CRDs, and RBAC artifacts. Most contributors interact with the controller binary during development and rely on the agent image when testing end to end.

## Daily Development Workflow

1. **Install CRDs locally** (so kubectl understands them): `make install`.
2. **Run unit tests**: `make test`. This target also runs `go fmt`, `go vet`, and generates deepcopy code.
3. **Run static checks**: `make lint` (installs `golangci-lint` the first time).
4. **Exercise end-to-end coverage**: `make local-e2e-test` will create a kind cluster, run the test suite under `test/e2e`, and tear everything down. Run it before you rely on the CI workflow, so regressions are caught locally.
5. **Iterate on the controller locally**:
   ```sh
   export ENABLE_WEBHOOKS=false  # speeds up local runs
   make run
   ```
   Keep a terminal tab for logs, use another for `kubectl` actions.
6. **Build container images** when you need to test in-cluster:
   ```sh
   export REPO=<your-registry-namespace>
   export PUSH=true            # optional, push automatically after build
   make docker-build
   ```
   This produces `virtbmc-controller` and `virtbmc` images tagged with the current branch (`<branch>-head`).

## Getting KubeVirtBMC Running End-to-End

1. **Deploy CRDs & controller** to your cluster:
   ```sh
   make install
   make deploy IMG=${REPO:-starbops}/virtbmc-controller:main-head
   ```
2. **Verify the pods**:
   ```sh
   kubectl -n kubevirtbmc-system get pods
   ```
3. **Create a sample VirtualMachine** (any KubeVirt VM works) and the corresponding `VirtualMachineBMC`. A starter manifest lives in `config/samples/virtualmachine_v1alpha1_virtualmachinebmc.yaml`.
4. **Check the BMC service**:
   ```sh
   kubectl -n kubevirtbmc-system get svc
   ```
5. **Access via IPMI or Redfish** from inside the cluster (see the README’s examples). Use `kubectl run -it --rm ipmitool ...` or `curl` for Redfish requests.
6. **Expose Redfish externally** using the Ingress example in `README.md` if you need off-cluster access.
7. **Tear down** when done:
   ```sh
   kubectl delete -k config/samples/
   make undeploy
   make uninstall
   ```

## Pull Request Requirements

- **Keep generated code fresh**: `make manifests generate generate-kubevirt-crd` must run cleanly with no diff. The `pull-request` workflow double-checks this (`.github/workflows/pull-request.yaml`, lines 36-43), so make sure the command exits without changes before you push.
- **Issue reference**: Mention the issue number in the PR description and include `Fixes #<id>` (or `Refs #<id>` if it only partially addresses it).
- **Developer Certificate of Origin (DCO)**: All commits must be signed. Use `git commit -s ...` and verify by running `git log --format=%B -n 1 HEAD` to ensure each commit contains `Signed-off-by: Your Name <email>`.
- **Style & checks**:
  - `make fmt`, `make lint`, and `make test` must pass locally.
  - Add or update tests when you change behavior; e2e tests live under `test/e2e`.
  - Run `make docker-build` if your change touches Dockerfiles or deployment artifacts.
- **Documentation**: Update `README.md`, Helm chart values, or other docs when you add/modify user-facing behavior. Small fixes go in the same PR; large docs-only efforts can be separate.
- **Review readiness**: Squash purely “fixup” commits before requesting review, keep PRs focused, and explain how you validated the change (commands, screenshots, etc.).

> We align with KubeVirt’s [contribution guidelines](https://kubevirt.io/community/contributing/) for expectations around testing, review responsiveness, and release note hygiene.

## Resources & Support

- **GitHub Discussions / Issues**: Preferred place for feature ideas or blockers.
- **KubeVirt Community**: SIG-Compute meetings and mailing lists often cover virt components relevant to this project.
- **Security disclosures**: Do **not** open a public issue. Email the maintainers listed in `OWNERS` (or the repository admins) with an encrypted report.

Thanks for making KubeVirtBMC better!

[^kubevirt-install]: [Installing KubeVirt on Kubernetes](https://kubevirt.io/user-guide/cluster_admin/installation/#installing-kubevirt-on-kubernetes)
[^cert-manager-install]: [cert-manager default static install](https://cert-manager.io/docs/installation/#default-static-install)
