FROM golang AS build-env

RUN GO111MODULE=off go get -u github.com/esrrhs/tcloud-cns-tool
RUN GO111MODULE=off go get -u github.com/esrrhs/tcloud-cns-tool/...
RUN GO111MODULE=off go install github.com/esrrhs/tcloud-cns-tool

FROM debian
COPY --from=build-env /go/bin/tcloud-cns-tool .
WORKDIR ./
