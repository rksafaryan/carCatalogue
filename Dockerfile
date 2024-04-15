# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o carcatalogue cmd/serve/main.go

EXPOSE 8080

# Run
RUN ls
CMD [ "./cmd/serve" ]
