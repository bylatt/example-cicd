FROM golang:1.12-alpine
ENV GO111MODULE=on
RUN apk add --update git
WORKDIR /go/src/github.com/clozed2u/example-cicd
COPY . .
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -u ./...
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o api ./transport/http/" -command="./api"
