# First Docker app

In this app, we learned how to build a docker image from an application, run the image in a docker container and then stop(kill) container from running

## Basic commands

- `docker --version` to confirm that we have docker installed in our machine
- `docker build -t <LOCAL_REPO_NAME> .` creates a docker image using the Dockerfile
- `docker images` list docker images available on your local machine (see metadata about an image like created, size, id, tag)
- `docker run <IMAGE_ID>` to run a docker image on a docker container. Running this command will block the terminal thread meaning that you cannot continue running command s on the same terminal window and you will be seeing any log output your application writes
- `docker ps` to see all containers that are running. Remember that docker images are run in a container
- `docker kill <CONTAINER_ID>` to stop a container from running

## Publishing Images to DockerHub

- Make sure you have a docker account and create a remote repo. For this project a created a basic public repo with name `<REMOTE_REPO_NAME>`, description and default options
- Build the image if na
- `docker tag <LOCAL_REPO_NAME> <DOCKER_HUB_USERNAME/REMOTE_REPO_NAME>` Tag local repo to the corresponding remote repo(usually the naming convention in DockerHub) so that we can have it push into the container registry
- `docker login --username=<USERNAME>` Login into your docker account so that you can connect to DockerHub. If its your first time you will prompt to enter password
- `docker push <DOCKER_HUB_USERNAME/REMOTE_REPO_NAME>:<OPTIONAL_TAGNAME>` Push Docker image to DockerHub so that other people can use our Docker images. if you don't specify the tagname like `docker push <DOCKER_HUB_USERNAME/REMOTE_REPO_NAME>:` the latest will be used by default

**It is important to have unique REPO_NAME so that you can easily reference a build; although when you want to revert to an older code you will have to reference to the unique ID like 09d7e1dbc2c4**

## Aha moment

To be able to send http requests from your local machine to the container app, you will have to bind port (8080) of the container to TCP port (8080) on localhost(127.0.0.1) of the host machine like `docker run -p 8080:8080 <IMAGE_ID>`. See [doc](https://docs.docker.com/engine/reference/commandline/run/#publish-or-expose-port--p---expose)
