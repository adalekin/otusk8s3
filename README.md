# otusk8s3

## TL;DR

```
helm repo add bitnami https://charts.bitnami.com/bitnami

brew install sops
helm plugin install https://github.com/zendesk/helm-secrets

helm install otusk8s3 deployments/helm
```

## Prerequisites

* Docker 18.06 or later
* Kubernetes 1.19.7 or later
    * NGINX Ingress Controller 0.45.0 or later
* Helm 3.4.1 or later
    * helm-secrets 2.0.3 or later
        * Mozilla SOPS 3.7.1 or later


## Installing the Helm chart

To install the chart with the release name `otusk8s3`:


```
helm install otusk8s3 deployments/helm
```

**NOTE:** you don't need to build a Docker image cause it already published at Docker Hub (https://hub.docker.com/repository/docker/adalekin/otusk8s3/).

## Uninstalling the Helm chart

To uninstall/delete the `otusk8s3` deployment:

```
helm delete otusk8s3
```

## Parameters

TODO
