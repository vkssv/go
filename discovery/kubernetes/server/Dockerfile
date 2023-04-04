FROM golang:1.20

ENV SOURCES /go/src/kubernetes/
COPY . ${SOURCES}

RUN cd ${SOURCES}server && CGO_ENABLED=0 go build
RUN go env -w GOBIN=/usr/local/go/bin
RUN cd ${SOURCES}server && go install
#WORKDIR ${SOURCES}server

CMD [ "server" ]
