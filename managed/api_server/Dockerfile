FROM golang:alpine

ENV API_SERVER_PORT 5000
ENV GCP_PROJECT_ID ""
ENV PUBSUB_TOPIC_NAME ""
ENV GCP_AUTH_FILE_LOCATION ""

ADD api_server.go /

RUN apk add --no-cache git && \
    go get -u cloud.google.com/go/pubsub && \
    go get -u github.com/google/uuid

WORKDIR /

ENTRYPOINT ["go", "run", "api_server.go"]