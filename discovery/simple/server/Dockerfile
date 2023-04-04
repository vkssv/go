FROM golang:1.20

ENV SOURCES /go/src/discovery/simple/
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build
RUN go env -w GOBIN=/usr/local/go/bin
RUN cd ${SOURCES}server/ && go install

ENV CONSUL_HTTP_ADDR localhost:8500
WORKDIR ${SOURCES}server/

CMD [ "server" ]

# security hardening
#EXPOSE 8080