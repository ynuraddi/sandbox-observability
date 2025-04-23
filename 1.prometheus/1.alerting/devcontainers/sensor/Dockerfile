FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ../../ .

RUN go mod download

RUN go build -o /main ./cmd/sensor

FROM scratch

WORKDIR /app

COPY --from=builder main /app

EXPOSE 8081

CMD [ "./main" ]