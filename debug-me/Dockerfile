FROM golang:1.19-alpine3.17
WORKDIR /app
COPY . .
# we define custom variable that we can pass to our set in our docker container during runtime like `docker run -e FAVORITE_FOOD=ogbono <IMAGE_ID>`
ARG FAVORITE_FOOD
#define environmetal variable for our image. 
ENV FAVORITE_FOOD=$FAVORITE_FOOD
RUN go build -o main main.go

EXPOSE 8080
CMD ["./main"]