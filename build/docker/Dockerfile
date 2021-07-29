FROM golang:1.16-alpine
RUN apk add --no-cache make
WORKDIR /app
COPY . .
RUN make build-linux

FROM scratch
WORKDIR /root/
COPY --from=0 /app/bin .
COPY --from=0 /app/db .
CMD ["./otusk8s3"]
