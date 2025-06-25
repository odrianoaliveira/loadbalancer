FROM gcr.io/distroless/static
COPY bin/loadbalancer /
COPY cmd/loadbalancer/config.yaml /cmd/loadbalancer/config.yaml
ENTRYPOINT ["/loadbalancer"]