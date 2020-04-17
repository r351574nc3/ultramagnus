FROM r351574nc3/bazel-onbuild:latest as builder

FROM debian:stretch

ENV GIN_MODE=release\
    DB_PATH=/opt/lib/ultramagnus/bolt.db\
    SLACK_TOKEN=token\
    HTTP_CONFIG=":8080"

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && apt-get clean

WORKDIR /opt

RUN mkdir -p /opt/bin \
    && mkdir -p /opt/lib/ultramagnus \
    && chmod -R 755 /opt/bin \
    && chmod -R 755 /opt/lib/ultramagnus 

COPY --from=builder /bazel/.output/execroot/__main__/bazel-out/k8-fastbuild/bin/application /opt/bin/ultramagnus

CMD ["/opt/bin/ultramagnus"]