FROM golang:latest

WORKDIR /app

copy go.mod go.sum ./