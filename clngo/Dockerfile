#FROM golang:1.20.1-alpine3.17

#LABEL org.opencontainers.image.authors="VK"

#ENV SOURCES /go/src/clngo/

#COPY . ${SOURCES}

#RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a

#ENV PORT 8080
#EXPOSE 8080

#ENTRYPOINT [ "clngo" ]

#FROM  golang:1.20.1-alpine3.17
FROM frolvlad/alpine-glibc

LABEL org.opencontainers.image.authors="VK"

COPY ./clngo /app/clngo

RUN chmod +x /app/clngo

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "/app/clngo" ]
