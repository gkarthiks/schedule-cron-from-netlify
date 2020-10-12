FROM golang:alpine3.11
ARG BUILD_VERSION
RUN mkdir -p /usr/local/src
COPY . /usr/local/src
WORKDIR /usr/local/src/
RUN go build .
CMD ./schedule-cron-from-netlify
