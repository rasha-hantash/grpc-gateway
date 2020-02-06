
FROM golang:latest as builder

RUN apt update && apt install unzip

WORKDIR /tmp
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
RUN unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/
RUN rm -rf protoc3

WORKDIR /go
RUN go get github.com/grpc-ecosystem/grpc-gateway/...

ENV GO111MODULE=on

RUN go get github.com/githubnemo/CompileDaemon && cp /go/bin/CompileDaemon /CompileDaemon

#RUN mkdir -p /go/src/github.com/rasha-hantash/grpc-gateway/grpc-example
WORKDIR /go/src/github.com/rasha-hantash/grpc-gateway/grpc-example

COPY ./ /go/src/github.com/rasha-hantash/grpc-gateway/grpc-example
RUN make all

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/rasha-hantash/grpc-gateway/grpc-example/bin/server /app/
COPY --from=builder /go/src/github.com/rasha-hantash/grpc-gateway/grpc-example/rewardsrefund/rewardsrefund.swagger.json /app/www/swagger.json


#EXPOSE 8080
#EXPOSE 5566


WORKDIR /app
CMD ["./server"]
