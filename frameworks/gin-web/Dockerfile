FROM golang:1.20
#FROM frolvlad/alpine-glibc

#RUN apk update && apk upgrade && apk add --no-cache bash git go
#RUN go install github.com/gin-gonic/gin@v1.9.0

ENV SOURCES /go/src/gin-web/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build
RUN go env -w GOBIN=/usr/local/go/bin
RUN cd ${SOURCES} && go install
WORKDIR ${SOURCES}

CMD [ "gin-web" ]

EXPOSE 8080
