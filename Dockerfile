
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...  # "go get -d -v ./..."
RUN go install -v ./...    # "go install -v ./..."

#final stage
FROM alpine:3.7
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=jq-in-go Version=0.0.1
