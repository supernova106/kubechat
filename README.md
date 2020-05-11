![Publish Docker](https://github.com/supernova106/kubechat/workflows/Publish%20Docker/badge.svg?branch=v0.2.0)
![goreleaser](https://github.com/supernova106/kubechat/workflows/goreleaser/badge.svg?branch=v0.2.0)

# kubechat
Kubernetes Chat Ops

## Getting started

- required Go `>= 1.14`
- the project is generated with [OpenAPI Generator](https://openapi-generator.tech/) :)

Deploy with helm

```sh
helm repo add supernova106 https://supernova106.github.io/charts
kubectl create ns kubechat
helm install kubechat supernova106/kubechat --namespace kubechat
```

Build from source code

```sh
git clone https://github.com/supernova106/kubechat.git
cd kubechat/

go build ./...
go run main.go
```

## Usage

Using swagger

- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Using curl

```sh
curl -G localhost:8080/v1/kubectl --data-urlencode "cmd=get nodes"
```

## Contact

Binh Nguyen
