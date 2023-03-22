FROM golang:1.20-alpine AS build

ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/deepset-ai/prompthub
COPY . /go/src/github.com/deepset-ai/prompthub
RUN go build


FROM alpine:3.17

COPY --from=build /go/src/github.com/deepset-ai/prompthub/prompthub /usr/bin/prompthub
COPY prompthub.yaml.example /prompthub.yaml
COPY prompts ./prompts
EXPOSE 80
CMD ["prompthub", "-c /prompthub.yaml"]
