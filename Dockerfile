FROM golang:1.20.3-bullseye

WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-w -s"

FROM alpine:3.17.2

WORKDIR /app
COPY --from=0 /build/kverify /app/kverify
COPY --from=0 /build/index.html /app/index.html
CMD [ "/app/kverify" ]
