# Build stage
FROM golang:1.21-alpine3.19 AS builder
WORKDIR /app
COPY . .

RUN apk update
# RUN apk add --virtual build-dependencies
RUN apk add libc-dev
RUN apk add make
RUN apk add gcc
RUN apk add bash
RUN apk add git

RUN make build

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/start_wars .
COPY app.env .
COPY db ./db

EXPOSE 8080
CMD [ "/app/start_wars" ]
ENTRYPOINT [ "/app/start_wars" ]