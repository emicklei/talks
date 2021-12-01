FROM golang

RUN go get golang.org/x/tools/cmd/present
COPY present-run /run/

EXPOSE 3999

WORKDIR /app

CMD ["/run/present-run"]