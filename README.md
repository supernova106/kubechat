# kubechat
Kubernetes Chat Ops

## Getting started

- required Go `>= 1.14`
- the project is generated with [OpenAPI Generator](https://openapi-generator.tech/) :)

Deploy with helm

```sh
helm repo add supernova106 https://supernova106.github.io/charts
helm install supernova106/kubechat
```

Build from source code

```sh
git clone https://github.com/supernova106/kubechat.git
cd kubechat/

go build ./...
go run main.go
```

## Usage

Using curl

```sh
curl -G localhost:8080/v1/kubectl --data-urlencode "cmd=get nodes"
```

## Contact

Binh Nguyen
