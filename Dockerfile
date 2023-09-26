FROM public.ecr.aws/docker/library/golang:1.21 as build

RUN echo "nobody:*:65534:65534:nobody:/_nonexistent:/bin/false" > /etc/passwd.minimal

WORKDIR /app

COPY go.mod go.mod ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -a -trimpath \
        -buildvcs=true -o /go/bin/gobinary \
        -tags osusergo,netgo \
        -ldflags "-s -w -extldflags '-static'" \
        -v .

#################################################
### Production Image
FROM scratch as runner
COPY --from=build /etc/passwd.minimal /etc/passwd
USER nobody

COPY --from=build --chown=nobody /etc/ssl/certs /etc/ssl/certs
COPY --from=build --chown=nobody /usr/share/zoneinfo /usr/share/zoneinfo

# Metadata params
ARG BUILD_DATE=`date`

# Metadata
LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.vendor="Frontend Masters" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.base.name="scratch"

COPY --from=build --chown=nobody /go/bin/gobinary /gobinary
ENTRYPOINT ["/gobinary"]
