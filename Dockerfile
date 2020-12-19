FROM golang
ADD . /go/src/github.com/rasha-hantash/grpc-gateway
WORKDIR /go/src/github.com/rasha-hantash/grpc-gateway
ENV GO111MODULE=on
RUN go mod download
RUN go install /go/src/github.com/rasha-hantash/grpc-gateway/server
ENTRYPOINT ["/go/bin/server"]
EXPOSE 5566
EXPOSE 8080
