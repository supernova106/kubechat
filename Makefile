generate:
	openapi-generator generate -g go-gin-server --additional-properties=packageName=kubechat,withGoCodegenComment=true,packageVersion=0.1.0 -i openapi.yaml 
build:
	docker build --network=host -t supernova106/kubechat .
push:
	docker push supernova106/kubechat