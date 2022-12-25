# Containers using Docker

In this project, I learned why we use containers and how to use Docker to build our application to be deployed as containers

- We learned about how containers are an abstraction of our application code and all the dependencies that they need to run.
- We learned about how to build a container and how it involves using a Dockerfile to create a Docker image
- We debugged a container as it is a bit different from what we are used to and went over some helpful command to facilitate this process
- We learned that containers should be considered immutable and modifying one usually entails creating a new container with a new image

Containers are self-contained applications with all the dependencies needed to run. So rather than only deploying your code that you write, you have to deploy all of its dependencies as well. A container can be treated as a unit of deployment and makes it easier and more reliable to deploy software. Lets say that we deploy some code that has a bug in it and we need to revert our changes, we don't have to figure out how to manually revert the code and all the dependencies and their versions, we can simply change the container to use an older snapshot(image) of the code that we know that is working. Since container are stateless, its important that the code you write for your container be stateless as well because it is always expected that container can be destroyed; therefore any data that you want to be persisted for your app should be stored in databases

## Docker

Docker is a platform that helps us manage the process of creating and managing our containers.

## Docker Image

When we have an application that we want to deploy, we can package it into a snapshot of Docker Image. The image contains all of your code and dependencies.

## Docker Container

A Docker Container is an ephemeral running instance of a Docker Image. Ephemeral is a software property where an application is expected to be short-lived; therefore a running container is expected to be killed at some point.

## Dockerfile

A Dockerfile defines the steps to create a Docker Image that can be run in containers

## Container Registry

A centralized place to store and version our container images. Eg [DockerHub](https://hub.docker.com/). DockerHub works thesame way as github where you push/pull image. In order to make sure that our application is used outside of just local development, we have to push image(s) to DockerHub

## Base Images

Base images reduce time that it takes to run redundant operations. Steps that are always repeated during any of your docker build projects can be stored as a base image
