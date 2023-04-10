FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /26-docker-alta

EXPOSE 8000

CMD [ "/26-docker-alta" ]