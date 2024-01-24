FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.21 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
ADD . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o /palog .

FROM --platform=${TARGETPLATFORM:-linux/amd64} debian:bookworm-slim
COPY --from=builder /palog /palog
RUN apt-get update && \
    apt-get install -y icu-devtools && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/palog"]
