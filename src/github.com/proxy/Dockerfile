FROM alpine:3.7

VOLUME /go/bin

RUN if [ "$SERVER_PORT_CACHE" = "" ]; \
    then export SERVER_PORT_CACHE=9200; \
    fi

EXPOSE $SERVER_PORT

COPY ./bin/sidecar-cache /go/bin/app

ENTRYPOINT ["/go/bin/app"]