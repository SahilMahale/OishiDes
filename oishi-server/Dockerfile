FROM golang:1.20.5-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy 

COPY . .

RUN go build /app/cmd/main.go

EXPOSE 8080

#CMD [ "/app/main","args" ]


