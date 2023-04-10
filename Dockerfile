FROM golang:1.20-buster as build-env

ENV GOPROXY=https://proxy.golang.org

ADD . /opt/src

WORKDIR /opt/src

RUN apt-get -y update \
    && apt-get install -y make git upx wget curl tar musl* \
    && go install github.com/swaggo/swag/cmd/swag@latest

ARG BUILD_DATE
ARG VCS_REF

RUN export PATH=/go/bin:$PATH && \
    swag init && \
    CC=/usr/local/bin/musl-gcc CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bid-tracker  -a -ldflags '-extldflags "-static" -s -w'  .

RUN upx /opt/src/build/bid-tracker

FROM scratch
# FROM gcr.io/distroless/base
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /opt/src/build/bid-tracker /
COPY --from=build-env /opt/src/docs /

ENTRYPOINT ["/bid-tracker"]
