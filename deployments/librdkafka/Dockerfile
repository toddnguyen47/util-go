# ------------------------------------------------------------
# global args
# ------------------------------------------------------------
ARG service_name=your_service_name
ARG librdkafka_folder=librdkafka_binary
ARG librdkafka_version=1.9.2

# ------------------------------------------------------------
# first stage: building librdkafka
# ------------------------------------------------------------

FROM alpine as librdkafka_builder
ARG librdkafka_folder
# install compiler
RUN apk add build-base curl tar bash
ARG librdkafka_version
WORKDIR /librdkafka
RUN curl -L https://github.com/edenhill/librdkafka/archive/refs/tags/v${librdkafka_version}.tar.gz \
  > librdkafka.tar.gz
RUN tar -xv -f librdkafka.tar.gz
RUN cd librdkafka-${librdkafka_version} && \
  ./configure --prefix=/${librdkafka_folder} && make && make install
WORKDIR /

# ------------------------------------------------------------
# downloading dependencies
# ------------------------------------------------------------

FROM librdkafka_builder as dep_builder
ARG service_name
RUN apk add go git pkgconfig
RUN go env -w GOPROXY=direct
WORKDIR /${service_name}
COPY go.mod go.sum /${service_name}
ENV GOPROXY=
ENV GONOSUMDB=
RUN go mod download -x

# ------------------------------------------------------------
# build Go binary
# ------------------------------------------------------------

FROM dep_builder as binary_builder
ARG service_name
ARG librdkafka_folder
WORKDIR /${service_name}
ADD internal /${service_name}/internal/
COPY main.go /${service_name}/
ENV PKG_CONFIG_PATH=/${librdkafka_folder}/lib/pkgconfig:$PKG_CONFIG_PATH
RUN go build -v -tags dynamic -o app_binary

# ------------------------------------------------------------
# run binary on clean image
# ------------------------------------------------------------

FROM alpine as final_stage
ARG service_name
ARG librdkafka_folder
ARG librdkafka_version
WORKDIR /app
COPY --from=binary_builder /${librdkafka_folder} /${librdkafka_folder}
COPY --from=binary_builder /${service_name}/app_binary /app/app_binary
ENV LD_LIBRARY_PATH=/${librdkafka_folder}/lib:$LD_LIBRARY_PATH

ENTRYPOINT [ /app/app_binary ]
