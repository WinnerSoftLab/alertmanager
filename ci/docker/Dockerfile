# Build image
FROM nexus.adsrv.wtf/click/golang:1.21.7-202402091158 as build


WORKDIR /app
USER root
RUN  apt-get update && \
     apt-get install -y --no-install-recommends git-restore-mtime

COPY --chown=10000:10000 ./ ./

USER 10000

RUN git restore-mtime && \
    make assets common-build

# Main image
FROM quay.io/prometheus/busybox-linux-amd64:latest
LABEL maintainer="The Prometheus Authors <prometheus-developers@googlegroups.com>"


COPY --from=build /app/amtool                            /bin/amtool
COPY --from=build /app/alertmanager                      /bin/alertmanager
COPY --from=build /app/examples/ha/alertmanager.yml      /etc/alertmanager/alertmanager.yml

RUN mkdir -p /alertmanager && \
    chown -R nobody:nobody etc/alertmanager /alertmanager

USER       nobody
EXPOSE     9093
VOLUME     [ "/alertmanager" ]
WORKDIR    /alertmanager
ENTRYPOINT [ "/bin/alertmanager" ]
CMD        [ "--config.file=/etc/alertmanager/alertmanager.yml", \
             "--storage.path=/alertmanager" ]
