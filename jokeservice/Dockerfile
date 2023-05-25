FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o run 
EXPOSE 8080
CMD ["./run"]
