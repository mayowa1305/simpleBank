#build stage
FROM golang:1.17.5-alpine AS builder
WORKDIR /app
COPY . . 
RUN go build -o main main.go

#run stage 
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
