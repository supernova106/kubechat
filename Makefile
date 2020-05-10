APP_VERSION = 0.1.0

generate:
	openapi-generator generate -g go-gin-server --additional-properties=packageName=kubechat,withGoCodegenComment=true,packageVersion=0.1.0 -i openapi.yaml 
build:
	docker build --network=host -t supernova106/kubechat:$(APP_VERSION) .
push:
	docker tag supernova106/kubechat:$(APP_VERSION) supernova106/kubechat:latest
	docker push supernova106/kubechat:$(APP_VERSION)
	docker push supernova106/kubechat:latest