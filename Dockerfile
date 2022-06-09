FROM golang:1.18-alpine AS build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /documentservice

# Add source files
COPY . .

# install
RUN make build

# Final image
FROM alpine:edge

RUN apk update && apk add bash

# Copy over binaries from the build-env
COPY --from=build-env /documentservice/build/documentserviced /usr/bin/documentserviced

RUN adduser container --disabled-password && mkdir -p /home/container/.documentservice

EXPOSE 26656 26657 1317 9090

COPY ./startup.sh /usr/local/bin/

COPY ./build/documentserviced /home/container/

RUN chmod +x /usr/local/bin/startup.sh

USER container:container

WORKDIR /home/container

ENTRYPOINT [ "/usr/local/bin/startup.sh" ]