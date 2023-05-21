FROM golang:1.19-alpine AS build-env

ARG MONIKER
ENV MONIKER=$MONIKER

ENV PACKAGES ca-certificates build-base curl make git bash linux-headers python3
RUN apk add --no-cache $PACKAGES

COPY . /go/delivery/defund

COPY ./devtools/entrypoint.sh /
RUN chmod +x /entrypoint.sh

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 9ecb037336bd56076573dc18c26631a9d2099a7f2b40dc04b6cae31ffb4c8f9a \
    && sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6e4de7ba9bad4ae9679c7f9ecf7e283dd0160e71567c6a7be6ae47c81ebe7f32

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp "/lib/libwasmvm_muslc.$(uname -m).a" /lib/libwasmvm_muslc.a

WORKDIR /go/delivery/defund

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make install

RUN defundd init $MONIKER

EXPOSE 26656 26657 1317 9090

CMD ["/bin/bash", "/entrypoint.sh"]
