FROM golang:latest
WORKDIR /build
COPY main.go /build
COPY start.sh /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=0 /build/start.sh .
COPY --from=0 /build/app .
RUN chmod 755 /app/start.sh && mkdir /app/static && chmod 755 /app/static
CMD ["/app/start.sh"]