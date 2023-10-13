# Node Taint Remover

Removes the kubernetes.io/arch from newly created nodes.

## Description

GKE clusters with only ARM nodes do not have a way to set tolerations on pods related to Dataplane V2. This is a quick and dirty means to resolve [this](https://serverfault.com/questions/1141581/gke-arm-based-cluster-starts-in-invalid-state). Some of charts that I install don't have any settings for tolerations.

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster

1. Build and push your image to the location specified by `IMG`:

    ```sh
    make docker-build docker-push IMG=<some-registry>/node-taint-remover:tag
    ```

2. Deploy the controller to the cluster with the image specified by `IMG`:

    ```sh
    make deploy IMG=<some-registry>/node-taint-remover:tag
    ```

### Undeploy controller

Undeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing

TODO - this will be a short lived tool as I don't think many will find it useful

### How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out

1. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

    ```sh
    make run
    ```

**NOTE:** You can also run this in one step by running: `make install run`

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023 Brian Kanya.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
