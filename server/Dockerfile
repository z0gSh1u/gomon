# Stage 1 - Build the Go app
FROM golang:alpine as AlpineBuilder

WORKDIR /go/src/github.com/z0gSh1u/gomon/server
COPY . .

RUN go env -w GO111MODULE=on \
 && go env -w GOPROXY=https://goproxy.cn,direct \
 && go env \
 && go mod tidy \
 && go build -o server .

FROM alpine:latest

WORKDIR /go/src/github.com/z0gSh1u/gomon/server

COPY --from=0 /go/src/github.com/z0gSh1u/gomon/server/server ./

EXPOSE 8080
ENTRYPOINT [ "./server" ]