FROM golang:1.25-alpine AS build
WORKDIR /src
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /out/app ./cmd

FROM alpine:3.20
WORKDIR /app
RUN adduser -D -H -u 10001 appuser && apk add --no-cache ca-certificates
COPY --from=build /out/app /app/app
COPY templates/ /app/templates/
COPY static/ /app/static/
COPY internal/repository/blog/ /app/internal/repository/blog/
RUN mkdir -p /app/tmp /tmp && chown -R appuser:appuser /app /tmp
USER appuser
EXPOSE 8080
ENTRYPOINT ["/app/app"]
