FROM golang:1.25.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download
COPY . .
## Tell the compiler not to use any C code or system libraries. It forces Go to create a standalone, "pure Go" binary that does not depend on the underlying operating system's C libraries (like libc)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app .

FROM alpine:latest

ENV APP_USER=nobody \
    APP_GROUP=nobody

RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=builder /app/app .

RUN chown -R $APP_USER:$APP_GROUP /app

USER $APP_USER

EXPOSE 8080

CMD ["./app"]
