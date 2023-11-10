# KubeBMC

KubeBMC unleashes the power management for virtual machines on Kubernetes in a traditional way, i.e., IPMI. This allows users to power on/off/reset and set the boot device (work in progress :wink:) for the VM. It was initially designed for Tinkerbell/Seeder to provision KubeVirt VMs, just like we did in the good old days.

The project was born in SUSE Hack Week 23.

## Description

KubeBMC was inspired by [VirtualBMC](https://opendev.org/openstack/virtualbmc). The difference between them could be illustrated below:

```mermaid
flowchart LR
    client1[Client]
    client2[Client]
    BMC1[BMC]
    VM[VM]
    subgraph KubeBMC
    direction LR
    client2-->|IPMI|kBMC-->|K8s API|VM
    end
    subgraph VirtualBMC
    direction LR
    client1-->|IPMI|vBMC-->|libvirt API|BMC1
    end
```

**Goals**

- Providing BMC functionalities for virtual machines powered by KubeVirt
- Providing in-cluster accessibility

**Non-goals**

- Providing BMC functionalities for bare-metal machines

KubeBMC consists of two components:

- **kubebmc-controller**: A typical Kubernetes controller built with kubebuilder that reconciles on the KubeBMC, VirtualMachine, and Service objects
- **kbmc**: A BMC simulator for serving IPMI and translating the requests to native Kubernetes API requests

Below is the workflow of KubeBMC when a VirtualMachine was created and booted up:

```mermaid
flowchart LR
    controller["kubebmc-controller"]
    cr["kubebmc CR"]
    kbmc-pod["kbmc Pod"]
    kbmc-svc["kbmc Service"]
    controller-.->|watches|cr
    cr-.->|owns|kbmc-svc
    cr-.->|owns|kbmc-pod
    client--->|IPMI|kbmc-svc
    kbmc-svc-->kbmc-pod
    kbmc-pod-->|HTTP|apiserver
    apiserver-->|modifies|vm
    vm-->|creates|vmi
```

The KubeBMC CR (CustomResource):

```go
type KubeBMCSpec struct {
	// To authenticate who the user is.
	// +optional
	Username string `json:"username,omitempty"`

	// The credential part of the IPMI service
	// +optional
	Password string `json:"password,omitempty"`

	// The namespace where the virtual machine is in
	VirtualMachineNamespace string `json:"vmNamespace"`

	// The actual virtual machine that this BMC controls
	VirtualMachineName string `json:"vmName"`
}

// KubeBMCStatus defines the observed state of KubeBMC
type KubeBMCStatus struct {
	// The listen IP address for the IPMI service.
	ServiceIP string `json:"serviceIP"`

	// The indicator that shows the readiness of the IPMI service for the virtual machine
	Ready bool `json:"ready"`
}
```

## Getting Started

### Prerequisites

- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster

**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/kubebmc:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/kubebmc:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin privileges or be logged in as admin.

**Create instances of your solution**

You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

The corresponding KubeBMC object will be created automatically when the VirtualMachine object exists. It will scaffold the kbmc Pod and Service object.

```sh
$ kubectl -n kubebmc-system get svc
NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
default-test-vm-kbmc               ClusterIP   10.53.106.65    <none>        623/UDP   3h13m
```

To access the VM's BMC, you need to be in the cluster network. Run a Pod that comes with `ipmitool` built in:

```sh
$ kubectl run -it --rm ipmitool --image=mikeynap/ipmitool --command -- /bin/sh
```

Inside the Pod, you can for example turn on the VM via `ipmitool`:

```sh
$ ipmitool -I lan -U admin -P password -H default-test-vm-kbmc.kubebmc-system.svc.cluster.local power status
Chassis Power is off
$ ipmitool -I lan -U admin -P password -H default-test-vm-kbmc.kubebmc-system.svc.cluster.local power on
Chassis Power Control: Up/On
$ ipmitool -I lan -U admin -P password -H default-test-vm-kbmc.kubebmc-system.svc.cluster.local power status
Chassis Power is on
```

> **NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall

**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Contributing

// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

