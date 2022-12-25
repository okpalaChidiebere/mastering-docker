# Debugging Containers

We learn how to debug containers in this project. We had a simple application that logs the favorite of of a user. The favorite food is a string that maybe be defined as an environment variable called FAVORITE_FOOD. if this variable is not defined an error is printed out in the logs instead of us exiting or killing the application.

A Docker Container is an ephemeral running instance of a Docker Image. Ephemeral is a software property where an application is expected to be short-lived; therefore a running container is expected to be killed at some point.

## Handy docker commands to debu g a container

- `docker --version` to confirm that we have docker installed in our machine
- `docker build -t <ANY_UNIQUE_IMAGE_NAME> .` creates a docker image using the Dockerfile
- `docker ps` list all containers running
- `docker ps -a` to see all status containers currently running including dead containers
- `docker inspect <CONTAINER_ID>` to see meta data about a container like status, uptime, etc
- `docker logs <CONTAINER_ID>` View output like logs from an application in the Docker container for troubleshooting
- `docker exec -it <CONTAINER_ID> sh` connect into the container shell interactively. You will not be able to run this all the time especially if the container is running production data. You can run linux commands like ps aux, etc. You run `exit` to . See [Attaching to Containers doc](https://docs.docker.com/engine/reference/commandline/container_attach/)
- `docker run -e FAVORITE_FOOD=<ANY_NAME> -d <IMAGE_ID>` the `-e` flag is used to feed environmental variables into the docker container running a docker image with env variable defined in the dockerfile used to build the image. The `-d` flag will start run the docker image without blocking the current window of terminal where its ran

## Modifying Containers

Docker images should be considered a single unit of deployment. You shouldn't be editing code or making changes to the system at all in a container like using connecting into the container shell and the editing the file(s). If something is broken, change the file(s) and fox the bug on your local environment, you build a new image and deploy that to a new container.
