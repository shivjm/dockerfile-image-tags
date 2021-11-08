FROM golang:1.16.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum main.go ./

COPY pkg ./pkg

COPY cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags '-w' .

FROM scratch

COPY --from=builder /app/dockerfile-image-tags /dockerfile-image-tags

ENTRYPOINT ["/dockerfile-image-tags"]
