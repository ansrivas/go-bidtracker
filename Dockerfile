FROM golang:1.14-buster as build-env

ENV GOPROXY=https://proxy.golang.org

RUN apt-get -y update \
    && apt-get install -y make git upx wget curl tar musl* \
    && go get -u github.com/go-bindata/go-bindata \
    && go get -u github.com/go-bindata/go-bindata/... \
    && go get -u github.com/swaggo/swag/cmd/swag

ARG BUILD_DATE
ARG VCS_REF

WORKDIR /opt/src

ADD . /opt/src

RUN export PATH=/go/bin:$PATH && \
    swag init pkg/api/ && \
    CC=/usr/local/bin/musl-gcc CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bid-tracker  -a -ldflags '-extldflags "-static" -s -w'  .

RUN upx /opt/src/build/bid-tracker

FROM scratch
# FROM gcr.io/distroless/base
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /opt/src/build/bid-tracker /
COPY --from=build-env /opt/src/docs /

ENTRYPOINT ["/bid-tracker"]