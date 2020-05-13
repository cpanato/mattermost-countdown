# Build the matterwick
ARG DOCKER_BUILD_IMAGE=golang:1.14.2
ARG DOCKER_BASE_IMAGE=alpine:3.11

FROM ${DOCKER_BUILD_IMAGE} AS build
WORKDIR /countdown/
COPY . /countdown/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .

# Final Image
FROM ${DOCKER_BASE_IMAGE}


ENV COUNTDOWN=/app/countdown \
    USER_UID=10001 \
  USER_NAME=countdown

WORKDIR /app/

RUN  apk update && apk add ca-certificates

COPY --from=build /countdown/countdown /app/
COPY --from=build /countdown/build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
