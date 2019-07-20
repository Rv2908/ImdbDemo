# Specify the base image 
FROM golang

# Install Some dependencies
RUN mkdir /app

COPY . /app

COPY go.mod /app
COPY go.sum /app

WORKDIR /app

RUN go build -o main .

# Default Command 
EXPOSE 3000
CMD ["/app/main"]