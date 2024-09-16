# base image
FROM golang:1.22.1

# set the working directory inside the container
WORKDIR /app

# copy the entire app directory to the container
COPY ./app /app

# download and install dependencies
RUN go mod download

# run the application using go run
CMD ["go", "run", "main.go"]
