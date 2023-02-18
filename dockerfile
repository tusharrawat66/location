# Start with the official Go image
FROM golang:latest
RUN mkdir /app

COPY . /app
WORKDIR /app

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main .


# Run the application
ENTRYPOINT [ "./main" ]
