FROM golang:latest

# Install packages here
RUN mkdir /go/src/myapp && cd /go/src/myapp &&  \
    go mod init github.com/my/repo && go get github.com/go-redis/redis/v8

EXPOSE 8074