FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o lifecrm ./cmd/server

FROM alpine:3.21
RUN apk add --no-cache ca-certificates && \
    addgroup -S lifecrm && adduser -S lifecrm -G lifecrm
WORKDIR /app
COPY --from=build /app/lifecrm .
COPY --from=build /app/migrations ./migrations
USER lifecrm
EXPOSE 8080
CMD ["./lifecrm"]
