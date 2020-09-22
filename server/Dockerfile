FROM golang:1.15.1-alpine3.12 AS build

# RUN apk update && apk add build-base
COPY . /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.12.0

ARG PORT
ENV GO_PORT=$PORT

EXPOSE ${GO_PORT} 
COPY --from=build /go/bin/server ./
CMD ./server ${GO_PORT}