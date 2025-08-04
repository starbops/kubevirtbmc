### 1. Start the Cluster and Setup Cert-Manager

```bash
./kubevirtci.sh up

export KUBECONFIG=$(./kubevirtci.sh kubeconfig)

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.2/cert-manager.yaml
```


### 2. Build Images

Build the controller and agent images locally using:

```bash
make docker-build REPO=registry:5000 TAG=dev PUSH=false
```


### 3. Push to Local Registry

Push tagged images to the local registry:

```bash
make push-local
```

### 4. Deploy the Controller

```bash
make deploy IMG=registry:5000/virtbmc-controller:dev
```

### 5. Apply Test Bundle

```bash
kubectl apply -f test-bundle.yaml
```
