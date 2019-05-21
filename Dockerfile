# Dockerfile for chat

# Start from golang 1.12
FROM golang:1.12

LABEL maintainer="Sebastian Edholm <sebbe@sebb.io>"

WORKDIR $GOPATH/src/mycode/chat

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...  

EXPOSE 8080
CMD ["chat"]
