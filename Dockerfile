FROM golang:latest
WORKDIR /app
COPY . .


RUN echo "golang" /app/a.txt

Run go build -o main main.go