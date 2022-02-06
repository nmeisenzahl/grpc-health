FROM golang:1.16-buster AS build

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o /app/grpc-health

FROM scratch

LABEL org.opencontainers.image.source github.com/nmeisenzahl/grpc-health

COPY --from=build /app/grpc-health /grpc-health

EXPOSE 3000

ENTRYPOINT ["/grpc-health"]
