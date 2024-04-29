# syntax=docker/dockerfile:experimental

############################
# STEP 2 build executable binary
############################
# FROM golang:alpine AS builder
FROM public.ecr.aws/docker/library/golang:1.19 as builder

RUN update-ca-certificates
# RUN apk update && apk add --no-cache git
# Create zincsearch user.
ENV USER=gateway
ENV GROUP=gateway
ENV UID=10001
ENV GID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN
RUN groupadd --gid "${GID}" "${GROUP}"
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    --gid "${GID}" \
    "${USER}"

WORKDIR $GOPATH/src/gateway
COPY ./gateway $GOPATH/src/gateway/
#COPY ./fs-lib-public $GOPATH/src/fs-lib-public/
RUN go mod tidy
RUN go build -o wzinc
############################
# STEP 3 build a small image
############################
FROM ubuntu:latest

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/src/gateway/wzinc /go/bin/wzinc

RUN apt-get update
RUN apt-get install -y poppler-utils wv unrtf tidy
RUN apt-get install -y inotify-tools
RUN apt-get install -y ca-certificates curl

# RUN sysctl -w fs.inotify.max_user_watches=8192

# Use an unprivileged user.
USER gateway:gateway
# Port on which the gateway service will be exposed.
EXPOSE 6317

CMD /go/bin/wzinc start
