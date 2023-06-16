# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files to the working directory
COPY go.mod go.sum ./

RUN go mod tidy -e
# Download the Go modules
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o app

# Set the environment variable for the MongoDB connection
ENV MONGO_URL mongodb://mongo:27017/mecnave

# Expose the port on which the Go application will listen
EXPOSE 8080

# Run the Go application
CMD ["./app"]
