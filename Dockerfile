FROM golang:1.14 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go get -v -t -d ./...

RUN go build -v -o kubechat .

######################
# Build the app image
######################
FROM alpine:3.11.6 AS runtime

RUN apk --no-cache add ca-certificates curl
RUN addgroup --gid 1000 -S app && adduser -S --uid 1000 -g app app
ENV GIN_MODE=release
COPY --from=build /go/src/kubechat /bin/kubechat

ENV KUBECTL_VERSION 1.15.11
RUN curl -Ls https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl -o /bin/kubectl && \
    chmod +x /bin/kubectl

EXPOSE 8080/tcp

USER app
CMD ["kubechat"]
