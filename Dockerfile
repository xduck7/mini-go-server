FROM golang:latest

LABEL maintainer="<inventionsList@xducky.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5050

RUN go build -o ./app ./cmd/app/main.go

RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./app"]
