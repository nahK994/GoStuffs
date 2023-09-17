FROM golang:1.21.1

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .

EXPOSE 8080
# RUN go mod download && go mod verify

# FROM python:3.10.6
# CMD [ "python3 script.py" ]
