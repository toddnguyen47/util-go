# ------------------------------------------------------------
# first stage: building librdkafka
# ------------------------------------------------------------

FROM gcc:12 as builder1
ARG librdkafka_version=1.9.2
WORKDIR /librdkafka
RUN curl -L https://github.com/edenhill/librdkafka/archive/refs/tags/v${librdkafka_version}.tar.gz \
  > librdkafka.tar.gz
RUN tar -xv -f librdkafka.tar.gz
RUN cd librdkafka-${librdkafka_version} && ./configure && make && make install
WORKDIR /
RUN rm -rfv /librdkafka

# ------------------------------------------------------------
# second stage: building go
# ------------------------------------------------------------

FROM builder1 as builder2
ARG go_version=1.19
WORKDIR /go${go_version}
RUN curl -L https://go.dev/dl/go${go_version}.linux-arm64.tar.gz > go.tar.gz
RUN tar -xv -f go.tar.gz
RUN rm go.tar.gz
RUN export PATH=/go${go_version}/go:$PATH

# ------------------------------------------------------------
# final stage: run binary
# ------------------------------------------------------------

WORKDIR /
CMD [ "sh" ]
