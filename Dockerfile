# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /carcatalogue

EXPOSE 8080

# Run
CMD [ "/carcatalogue" ]
