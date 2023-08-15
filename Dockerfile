# syntax=docker/dockerfile:1

FROM golang:latest as build

WORKDIR /app

COPY . .
# RUN apt-get update && apt-get install -y default-mysql-server
# RUN go mod tidy
RUN go mod vendor
RUN go mod tidy

# RUN go build -o mvc ./cmd/main.go

# FROM mysql:5.7
# WORKDIR /app/mvc
# 
FROM mysql:latest
WORKDIR /app/mvc
COPY --from=build /app .
EXPOSE 8000
CMD /app/mvc/script.sh

# FROM mysql:latest

# # Copy the SQL scripts into the container
# COPY sql-scripts/ /docker-entrypoint-initdb.d/

