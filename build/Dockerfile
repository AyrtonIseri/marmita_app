FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

ENV TWILIO_NUMBER=testing

CMD ["go", "run", "cmd/webserver/main.go"]
