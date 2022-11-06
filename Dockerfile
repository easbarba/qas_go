FROM golang:1.19-alpine AS build
MAINTAINER EAS Barbosa <easbarba@outlook.com>

WORKDIR /app

COPY examples /root/.config/qas

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o qas .

COPY . ./

COPY --from=build /app/qas

CMD [ "/app/qas" ]
