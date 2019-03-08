FROM golang:1.12-alpine as builder

WORKDIR /go/src/gorm-hello-world
RUN apk add --no-cache gcc musl-dev
COPY . .

RUN go build -v -ldflags "-linkmode external -extldflags -static -s -w"

FROM alpine:3.9

WORKDIR /app
COPY --from=builder /go/src/gorm-hello-world/gorm-hello-world /app/
# Might need to copy over models, pics, or whatever data needed for ML
ENTRYPOINT /app/gorm-hello-world
EXPOSE 4531