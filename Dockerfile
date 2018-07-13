FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/facundomedica/go_grpc_flutter
COPY . /usr/local/go/src/github.com/facundomedica/go_grpc_flutter
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/go_grpc_flutter ./go_grpc_flutter


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/facundomedica/go_grpc_flutter/build/go_grpc_flutter /bin/go_grpc_flutter
CMD ["go_grpc_flutter", "up"]
