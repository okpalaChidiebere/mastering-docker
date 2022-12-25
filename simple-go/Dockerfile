# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image.
# since our app is in golang, we need a golang base image to build it. Base images reduce time that it takes to run redundant operations; therefore helps us reduce build times! :)
#in order to produce a small golang output image we use the alphine version; I used one of the base dimages in the official doc https://hub.docker.com/_/golang
FROM golang:1.19-alpine3.17

# Add Maintainer Info
LABEL maintainer="Chidiebere Okpala <okpalacollins4@gmail.com>"

# define the Current Working Directory inside the image
WORKDIR /app

# Copy all necessary files from the current directory to the Working Directory inside the image
#the first dot is the current folder and he secont dot is the current directory inside the image `\app` where all the files will be copied over
COPY . .

# Build the Go app to a single executable file called `main`
RUN go build -o main main.go

# Expose port 8080 to the outside world. This informs docker to listen for this specific network port at runtime
EXPOSE 8080

# Command to run the executable file we build. We define the Docker image's behavior at runtime
CMD ["./main"]