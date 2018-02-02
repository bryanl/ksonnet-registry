# running minio in minikube

## 1. Prerequisites

Ensure `minikube` and `kubectl` are installed.

## 2. Start minio

* Run minio

```sh
./minio_distributed.sh
```

To expose the minio service, execute:

```sh
kubectl port-forward minio-0 9001:9000
```
