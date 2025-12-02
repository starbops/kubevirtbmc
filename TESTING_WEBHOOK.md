# Testing the v1beta1 Webhook

This guide explains how to test the VirtualMachineBMC webhook validation.

## Prerequisites

1. **Install dependencies** (if not already done):
   ```bash
   make envtest
   make controller-gen
   ```

2. **Generate required CRDs**:
   ```bash
   make generate-kubevirt-crd
   make manifests
   ```

## Running Webhook Tests

### Option 1: Run All Tests (Recommended)

This runs all tests including the webhook tests:

```bash
make test
```

This command will:
- Generate manifests and CRDs
- Run `go fmt` and `go vet`
- Download envtest binaries if needed
- Run all tests (excluding e2e tests in `/test` directory)
- Generate a coverage report

### Option 2: Run Only Webhook Tests

Run tests for the specific webhook package:

```bash
cd internal/webhook/bmc/v1beta1
go test -v .
```

Or from the project root:

```bash
go test -v ./internal/webhook/bmc/v1beta1/...
```

### Option 3: Run Specific Test Cases

Run a specific test by name:

```bash
go test -v ./internal/webhook/bmc/v1beta1/... -ginkgo.focus="Should deny creation if VirtualMachineRef is nil"
```

### Option 4: Run Tests with Coverage

Generate a coverage report for webhook tests:

```bash
go test -v ./internal/webhook/bmc/v1beta1/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Environment

The webhook tests use `envtest` which:
- Starts a local Kubernetes API server
- Loads CRDs from `config/crd/bases` and `config/kubevirt-crd`
- Sets up webhook server on a local port
- Provides a real Kubernetes client for testing

## What the Tests Cover

The webhook tests validate:

### Create Operations:
- ✅ Rejects if `VirtualMachineRef` is nil
- ✅ Rejects if `VirtualMachineRef.Name` is empty
- ✅ Rejects if referenced VirtualMachine doesn't exist
- ✅ Rejects if `AuthSecretRef` is provided but Secret doesn't exist
- ✅ Rejects if `AuthSecretRef.Name` is empty when `AuthSecretRef` is provided
- ✅ Accepts if VirtualMachine exists and `AuthSecretRef` is not provided
- ✅ Accepts if both VirtualMachine and Secret exist

### Update Operations:
- ✅ Rejects if `VirtualMachineRef` is nil
- ✅ Rejects if VirtualMachine doesn't exist
- ✅ Rejects if `AuthSecretRef` is provided but Secret doesn't exist
- ✅ Accepts if both VirtualMachine and Secret exist

## Troubleshooting

### Issue: Tests fail with "CRD not found"

**Solution**: Make sure you've generated the KubeVirt CRDs:
```bash
make generate-kubevirt-crd
```

### Issue: Tests fail with "envtest binaries not found"

**Solution**: Install envtest:
```bash
make envtest
```

### Issue: Tests timeout waiting for webhook server

**Solution**: Check that the webhook manifests are generated:
```bash
make manifests
```

### Issue: "VirtualMachine type not found"

**Solution**: Ensure KubeVirt scheme is added. Check `webhook_suite_test.go` includes:
```go
err = kubevirtv1.AddToScheme(scheme.Scheme)
```

## Running Tests in IDE

Most IDEs (VS Code, GoLand, etc.) can run the tests directly:

1. Open `internal/webhook/bmc/v1beta1/virtualmachinebmc_webhook_test.go`
2. Click the "Run Test" button next to the test function
3. Or use the test runner panel

**Note**: Make sure `KUBEBUILDER_ASSETS` environment variable is set if running from IDE:
```bash
export KUBEBUILDER_ASSETS=$(make envtest --dry-run | grep KUBEBUILDER_ASSETS | cut -d'=' -f2)
```

## Continuous Integration

The tests are automatically run in CI/CD pipelines when you:
- Push to a branch
- Create a pull request
- Merge to main

The CI uses the same `make test` command.

