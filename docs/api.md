# Ksonnet Registry API

## Parts

Parts will require a change in their spec. If a part has a dependency on another part, it will be required to state that dependency.

## S3 Storage

In S3 store the following

* `digest/parts.yaml`
* `digest/part.tar.gz`
* `digest/README.md`

In the repository, store the namespace/package/release metadata.