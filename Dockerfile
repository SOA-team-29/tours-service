FROM golang:alpine AS tours-builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o tours-webapp

FROM alpine
COPY --from=tours-builder /app/tours-webapp /usr/bin/tours-webapp
EXPOSE 8081
ENTRYPOINT ["/usr/bin/tours-webapp"]