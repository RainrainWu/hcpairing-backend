# Build Base Container
FROM golang:1.14-stretch AS base

ENV CGO_ENABLED=0
WORKDIR /hcpairing

COPY . .

RUN go mod download && \
    go build -o ./build/hcpairing ./cmd && \
    chmod +x ./build/hcpairing

# Application Container
FROM alpine:3.13

COPY --from=base /hcpairing/build/hcpairing /app/hcpairing

RUN adduser -S rain && chown -R rain /app
USER rain
EXPOSE 8080
EXPOSE 443
CMD ["./app/hcpairing"]