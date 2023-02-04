FROM golang AS gol

WORKDIR /build

ADD go.mod .

COPY . .

RUN ["go", "build", "./app"]

ENTRYPOINT ["./client"]