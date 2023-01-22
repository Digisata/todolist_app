FROM golang:1.18-alpine
LABEL maintener="Hanif Naufal <hnaufal123@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./out/dist .
EXPOSE 3030
CMD ["./out/dist"]