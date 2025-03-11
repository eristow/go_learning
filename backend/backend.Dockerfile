FROM golang:1.23.6 AS builder

ARG DATABASE_URL
ENV DATABASE_URL=$DATABASE_URL

WORKDIR /build

COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./main


FROM scratch AS backend

ARG DATABASE_URL
ENV DATABASE_URL=$DATABASE_URL

WORKDIR /app

COPY --from=builder /build/main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]
