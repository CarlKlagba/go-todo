FROM golang:1.19


WORKDIR /app

COPY ../go.mod ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-todo-app



CMD ["/docker-todo-app"]