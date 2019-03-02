FROM golang:alpine

ENV API_SERVER_PORT 5000
ENV GCP_PROJECT_ID ""
ENV PUBSUB_TOPIC_NAME ""

ADD api_server.go /
ADD auth.json /

RUN apk add --no-cache git && \
    go get -u cloud.google.com/go/pubsub && \
    go get -u github.com/google/uuid

WORKDIR /

ENTRYPOINT ["go", "run", "api_server.go"]