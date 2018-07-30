FROM golang:1.10.3 AS GoBuilder
# Copy in all the supporting code into the container
# ready to copy over the final binary into the tagged container
# This also ensures that we build a static binary that could be used in any docker image
WORKDIR /go/src/github.com/BeameryHQ/kubeaware
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    GODEBUG=netdns=go CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /kubeaware

# Final static container that only contains the precompiled binary.
FROM alpine:3.7
LABEL Author="Sean ZO Marciniak <sean@beamery.com>" \
      Owner="Beamery HQ" \
      Version="Pre-Alpha" \
      Project="KubeAware" \
      License="MIT"
COPY --from=GoBuilder /kubeaware /kubeaware
RUN apk --no-cache add ca-certificates

ENTRYPOINT ["/kubeaware"]
