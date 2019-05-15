FROM golang:1.12-alpine AS builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN apk add --nocache git
WORKDIR /go/src/github.com/clozed2u/example-cicd
COPY . .
RUN go get -u ./...
RUN go build -o api ./transport/http

FROM alpine
RUN apk add --nocache ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/clozed2u/example-cicd/api .
CMD ./api