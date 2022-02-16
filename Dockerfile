FROM golang:1.16-alpine

RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh build-base make git gcc libtool

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main .

EXPOSE 8080

CMD ["./main"]
