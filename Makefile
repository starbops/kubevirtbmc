# VERSION and COMMIT are set by the CI/CD pipeline. If not set, they are set to
# the current branch and commit.
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || echo "$(shell git rev-parse --abbrev-ref HEAD)-head")
COMMIT ?= $(shell git rev-parse HEAD)

DIRTY :=
ifneq ($(shell git status --porcelain --untracked-files=no),)
DIRTY := -dirty
endif
VERSION := $(VERSION)$(DIRTY)
# Export the tag to be used in the e2e tests
export TAG = $(VERSION)

REPO ?= starbops

# Image URL to use all building/pushing image targets
MGR_IMG ?= $(REPO)/virtbmc-controller:$(TAG)
AGT_IMG ?= $(REPO)/virtbmc:$(TAG)

K8S_VERSION = 1.28.13
KIND_K8S_VERSION = v$(shell echo $(K8S_VERSION))
# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.28.x
export CERT_MANAGER_VERSION = v1.14.2

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# CONTAINER_TOOL defines the container tool to be used for building images.
# Be aware that the target commands are only tested with Docker which is
# scaffolded by default. However, you might want to replace it to use other
# tools. (i.e. podman)
CONTAINER_TOOL ?= docker

# KUBEVIRT API version to use
KUBEVIRT_API_VERSION = v1.2.0

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook \
		paths="./api/...;./internal/...;./pkg/...;./cmd/..." \
		output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" \
		paths="./api/...;./internal/...;./pkg/...;./cmd/..."

.PHONY: generate-kubevirt-crd
generate-kubevirt-crd: controller-gen ## Clone KubeVirt API and generate CustomResourceDefinition objects for integration testing purposes.
	TMP_DIR=$$(mktemp -d -p /tmp/); \
	KUBEVIRT_API_DIR=$$TMP_DIR/kubevirt-api; \
	git clone --depth 1 --branch $(KUBEVIRT_API_VERSION) https://github.com/kubevirt/api $$KUBEVIRT_API_DIR; \
	$(CONTROLLER_GEN) crd:allowDangerousTypes=true paths="$$KUBEVIRT_API_DIR/core/v1/..." output:crd:artifacts:config=config/kubevirt-crd; \
	rm -rvf $$TMP_DIR; \
	rm -vf config/kubevirt-crd/kubevirt.io_datavolumetemplatespecs.yaml

.PHONY: generate-mock
generate-mock: mockgen ## Generate mocks for interfaces.
	# $(MOCKGEN) -destination=mocks/mock_resourcemanager.go -package=mocks kubevirt.io/kubevirtbmc/pkg/resourcemanager ResourceManager
	$(MOCKGEN) -source=pkg/resourcemanager/resource_manager.go -destination=pkg/resourcemanager/mock_resource_manager.go -package=resourcemanager

REDFISH_SCHEMA_BUNDLE ?= DSP8010_2024.3
.PHONY: download-redfish-schema
download-schema: ## Download the Redfish schema.
	test -d ./hack/$(REDFISH_SCHEMA_BUNDLE) || \
	( curl -sSL https://www.dmtf.org/sites/default/files/standards/documents/$(REDFISH_SCHEMA_BUNDLE).zip -o ./hack/$(REDFISH_SCHEMA_BUNDLE).zip && \
	unzip -q -d ./hack/ ./hack/$(REDFISH_SCHEMA_BUNDLE).zip && \
	rm -f ./hack/$(REDFISH_SCHEMA_BUNDLE).zip )

.PHONY: generate-redfish-api
generate-redfish-api: ## Generate Redfish API server.
	./hack/redfish/generate.sh

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: manifests generate generate-kubevirt-crd fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" go test $(shell go list ./... | grep -v /test/) -coverprofile cover.out

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.61.0
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

.PHONY: e2e-setup
e2e-setup: kind ## Setup end-to-end test environment.
	$(KIND) create cluster --name kvbmc-e2e --config test/e2e/kind-config.yaml --image=kindest/node:$(KIND_K8S_VERSION)

.PHONY: e2e-teardown
e2e-teardown: kind ## Teardown end-to-end test environment.
	$(KIND) delete cluster --name kvbmc-e2e

.PHONY: e2e-test
e2e-test: generate fmt vet kind ## Run end-to-end tests.
	go test -v ./test/...

.PHONY: local-e2e-test
local-e2e-test: e2e-setup e2e-test e2e-teardown ## Run end-to-end tests locally.

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter & yamllint
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix

##@ Build

LINKFLAGS := "-X main.AppVersion=$(VERSION) -X main.GitCommit=$(COMMIT)"

.PHONY: build
build: manifests generate fmt vet ## Build manager binary.
	go build -ldflags $(LINKFLAGS) -o bin/manager cmd/controller/main.go
	go build -ldflags $(LINKFLAGS) -o bin/virtbmc cmd/virtbmc/main.go

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	go run ./cmd/controller/main.go

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker images with the manager and the agent respectively.
	$(CONTAINER_TOOL) build -t $(MGR_IMG) --build-arg LINKFLAGS=$(LINKFLAGS) .
	$(CONTAINER_TOOL) build -t $(AGT_IMG) --build-arg LINKFLAGS=$(LINKFLAGS) --build-arg TARGETARCH=amd64 -f Dockerfile.virtbmc .
ifeq ($(PUSH),true)
	$(CONTAINER_TOOL) push $(MGR_IMG)
	$(CONTAINER_TOOL) push $(AGT_IMG)
endif

# PLATFORMS defines the target platforms for the manager image be built to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - be able to use docker buildx. More info: https://docs.docker.com/build/buildx/
# - have enabled BuildKit. More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image to your registry (i.e. if you do not set a valid value via IMG=<myregistry/image:<tag>> then the export will fail)
# To adequately provide solutions that are compatible with multiple platforms, you should consider using this option.
PLATFORMS ?= linux/arm64,linux/amd64,linux/s390x,linux/ppc64le
.PHONY: docker-buildx
docker-buildx: ## Build and push docker image for the manager for cross-platform support
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile.virtbmc > Dockerfile.virtbmc.cross
	- $(CONTAINER_TOOL) buildx create --name project-v3-builder
	$(CONTAINER_TOOL) buildx use project-v3-builder
	- $(CONTAINER_TOOL) buildx build --push --platform=$(PLATFORMS) --build-arg LINKFLAGS=$(LINKFLAGS) --tag $(MGR_IMG) -f Dockerfile.cross .
	- $(CONTAINER_TOOL) buildx build --push --platform=$(PLATFORMS) --build-arg LINKFLAGS=$(LINKFLAGS) --tag $(AGT_IMG) -f Dockerfile.virtbmc.cross .
	- $(CONTAINER_TOOL) buildx rm project-v3-builder
	rm Dockerfile.cross
	rm Dockerfile.virtbmc.cross

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

.PHONY: install
install: manifests kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | $(KUBECTL) apply -f -

.PHONY: uninstall
uninstall: manifests kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/crd | $(KUBECTL) delete --ignore-not-found=$(ignore-not-found) -f -

include .config.env
export

MGR_IMG := $(REGISTRY_HOST):$(REGISTRY_PORT)/$(CONTROLLER_IMAGE_NAME):$(CONTROLLER_IMAGE_TAG)
IMG ?= $(MGR_IMG)

.PHONY: deploy
deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | $(KUBECTL) apply -f -

.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/default | $(KUBECTL) delete --ignore-not-found=$(ignore-not-found) -f -

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
KUBECTL ?= kubectl
KUSTOMIZE ?= $(LOCALBIN)/kustomize
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
ENVTEST ?= $(LOCALBIN)/setup-envtest
KIND ?= $(LOCALBIN)/kind
MOCKGEN ?= $(LOCALBIN)/mockgen

## Tool Versions
KUSTOMIZE_VERSION ?= v5.2.1
CONTROLLER_TOOLS_VERSION ?= v0.15.0
KIND_VERSION ?= v0.24.0
MOCKGEN_VERSION ?= v0.5.0

.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary. If wrong version is installed, it will be removed before downloading.
$(KUSTOMIZE): $(LOCALBIN)
	@if test -x $(LOCALBIN)/kustomize && ! $(LOCALBIN)/kustomize version | grep -q $(KUSTOMIZE_VERSION); then \
		echo "$(LOCALBIN)/kustomize version is not expected $(KUSTOMIZE_VERSION). Removing it before installing."; \
		rm -rf $(LOCALBIN)/kustomize; \
	fi
	test -s $(LOCALBIN)/kustomize || GOBIN=$(LOCALBIN) GO111MODULE=on go install sigs.k8s.io/kustomize/kustomize/v5@$(KUSTOMIZE_VERSION)

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary. If wrong version is installed, it will be overwritten.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen && $(LOCALBIN)/controller-gen --version | grep -q $(CONTROLLER_TOOLS_VERSION) || \
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	test -s $(LOCALBIN)/setup-envtest || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest

.PHONY: kind
kind: $(KIND) ## Download kind locally if necessary.
$(KIND): $(LOCALBIN)
	test -s $(LOCALBIN)/kind || GOBIN=$(LOCALBIN) GO111MODULE=on go install sigs.k8s.io/kind@$(KIND_VERSION)

.PHONY: mockgen
mockgen: $(MOCKGEN) ## Download mockgen locally if necessary.
$(MOCKGEN): $(LOCALBIN)
	test -s $(LOCALBIN)/mockgen || GOBIN=$(LOCALBIN) go install go.uber.org/mock/mockgen@$(MOCKGEN_VERSION)

