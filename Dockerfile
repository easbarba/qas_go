FROM golang:1.19-alpine AS build
MAINTAINER EAS Barbosa <easbarba@outlook.com>

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux GOARCH=amd64
RUN go build -o ./qas ./cmd/qas/main.go

FROM golang:1.19-alpine
COPY --from=build /app/qas /opt/qas
COPY examples /root/.config/qas
CMD [ "/opt/qas" ]
