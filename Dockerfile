FROM golang:1.16.5-alpine as builder

ARG VERSION
ARG REVISION

WORKDIR /app

COPY go.mod go.sum main.go ./

COPY pkg ./pkg

COPY cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/shivjm/dockerfile-image-tags/cmd.Version=$VERSION' -X 'github.com/shivjm/dockerfile-image-tags/cmd.Revision=$REVISION' -w" -tags netgo .

FROM scratch

COPY --from=builder /app/dockerfile-image-tags /dockerfile-image-tags

ENTRYPOINT ["/dockerfile-image-tags"]
