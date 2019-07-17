FROM golang:1.12.6-alpine3.10

RUN apk update \
    && apk add --virtual build-dependencies build-base gcc

WORKDIR /go

COPY src/go-hpa /go/src/go-hpa

RUN go build -ldflags "-w" -o /go/bin/go-hpa /go/src/go-hpa/go-hpa.go

RUN apk del build-dependencies build-base gcc

#EXPOSE 8000

ENTRYPOINT ["go"]
