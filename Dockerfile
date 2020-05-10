FROM golang:1.14 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o kubechat .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/kubechat ./

ENV KUBECTL_VERSION 1.15.11
RUN cd /usr/local/bin && \
    curl -Ls https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl -o kubectl && \
    chmod +x kubectl

EXPOSE 8080/tcp
ENTRYPOINT ["./kubechat"]
