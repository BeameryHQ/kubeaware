FROM golang:1.10.3 AS GoBuilder


WORKDIR /go/src/github.com/BeameryHQ/kubeaware
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep &&
    dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /kubeaware

FROM alpine:3.7

LABEL Author="Sean ZO Marciniak <sean@beamery.com>" \
      Owner="Beamery HQ" \
      Version="Pre-Alpha" \
      Project="KubeAware" \
      License="MIT"

COPY --from=GoBuilder /kubeaware /kubeaware

RUN apk --no-cache add ca-certificates

ENTRYPOINT ["/kubeaware"]
