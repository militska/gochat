FROM golang:latest

# Install packages here
RUN mkdir /go/src/myapp && cd /go/src/myapp &&  \
    go mod init github.com/my/repo && go get github.com/go-redis/redis/v8 && \
    go get github.com/gorilla/websocket

EXPOSE 8074
#EXPOSE 8082