FROM 1.20.3-bullseye

WORKDIR /build
COPY . /build
RUN go build -trimpath -ldflags "-w -s"

FROM alpine:3.17.2

WORKDIR /app
COPY --from=0 /build/kverify /app/kverify
CMD [ "/app/kverify" ]
