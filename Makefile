generate:
	openapi-generator generate -g go-gin-server --additional-properties=packageName=kubechat,withGoCodegenComment=true,packageVersion=0.1.0 -i openapi.yaml 
