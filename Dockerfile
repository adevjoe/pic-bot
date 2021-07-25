FROM golang:1.16 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -a -v -o /app/pic-bot

FROM alpine:3.14
RUN apk --no-cache add \
  ca-certificates python3 py3-pip ffmpeg youtube-dl

# install gallery-dl
RUN python3 -m pip install -U gallery-dl

COPY --from=builder /app/pic-bot /usr/local/bin

RUN chmod +x /usr/local/bin/pic-bot

ENTRYPOINT ["/usr/local/bin/pic-bot"]
