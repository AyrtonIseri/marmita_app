FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

ENV TWILIO_NUMBER=whatsapp:+14155238886

CMD ["go", "run", "cmd/webserver/main.go"]
