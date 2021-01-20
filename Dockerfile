FROM golang:1.15.7-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build  -tags netgo -ldflags '-s -w -extldflags "-static"' \
    -o basicapi ./cmd/api

FROM scratch
WORKDIR /app
COPY --from=build /src/basicapi /app/
EXPOSE 8080
CMD ["./basicapi"]