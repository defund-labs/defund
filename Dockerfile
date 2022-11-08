FROM golang:1.19

ARG VERSION_TAG=v0.1.0
ENV VERSION_TAG=$VERSION_TAG

ARG MONIKER=moniker
ENV MONIKER=$MONIKER

RUN git clone https://github.com/defund-labs/defund

WORKDIR /defund

COPY ./scripts/entrypoint.sh /

RUN chmod +x /entrypoint.sh

RUN git checkout $VERSION_TAG

RUN make install

RUN defundd init $MONIKER

CMD ["/bin/bash", "/entrypoint.sh"]