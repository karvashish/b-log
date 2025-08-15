FROM golang:1.25-alpine AS build
WORKDIR /src
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /out/app ./cmd/server

FROM alpine:3.20
WORKDIR /app
RUN adduser -D -H -u 10001 appuser
COPY --from=build /out/app /app/app
COPY templates/ /app/templates/
COPY static/ /app/static/
EXPOSE 8080
USER appuser
ENTRYPOINT ["/app/app"]
