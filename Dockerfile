FROM golang

RUN go install golang.org/x/tools/cmd/present@latest

WORKDIR /app

CMD ["/go/bin/present", "-http", "0.0.0.0:3999", "-play", "-notes"]
