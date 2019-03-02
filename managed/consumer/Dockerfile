FROM golang:alpine

ENV GCP_PROJECT_ID ""
ENV PUBSUB_TOPIC_NAME ""
ENV PUBSUB_SUBSCRIPTION_ID ""

ADD pubsub_consumer.go /
ADD auth.json /

RUN apk add --no-cache git && \
    go get -u cloud.google.com/go/pubsub && \
    go get -u github.com/google/uuid

WORKDIR /

ENTRYPOINT ["go", "run", "pubsub_consumer.go"]