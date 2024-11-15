FROM golang:1.23-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .


# Copy the source code. Note the slash at the end, as explained in

COPY *.go ./

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.

EXPOSE 9090

# Run
CMD ["main.go"]