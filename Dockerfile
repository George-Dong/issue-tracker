FROM golang:alpine
WORKDIR /tracker
COPY go.mod go.sum ./
RUN go mod download
CMD ["./generate.sh"]

