---
layout: default
title: Docker and Kubernetes Samples
---

This repository contains samples used to demonstrate concepts of docker standalone, docker swarm mode, and kubernetes.

For all samples, we will try to provide versions for Linux and Windows, with the Linux samples being created first. For all orchestration samples, we will try to provide versions on docker swarm mode and kubernetes, with the swarm mode samples being created first.

We try to follow git flow. The master branch contains "final" code; active development happens in the "development" branch.

Contributions are welcome. Feel free to fork and send pull requests.

## How to use the samples

Every sample will have accompanying documentation which explains how to use that sample. However, all samples will follow some common rules.

Every sample will contain at least one of the following:

* Images (Build context directories with Dockerfiles)
* Docker Compose files (for standalone and/or stack deployment)
* Kubernetes manifest files (yaml files)

### Images

If a sample contains images, the sources will be in a directory called `images` under the sample directory. Each subdirectory under images represents an individual image, which should be tagged as _sample-name_-_image-name_. Under each image directory, there will be linux and windows subdirectories, which will contain Dockerfiles. For windows, there may be multiple subdirectories for different Windows versions.

To build an image, the image directory must be used as the build context directory. This means that you build Linux and Windows images separately, as follows:

#### Linux

```bash
docker build -t <sample-name>-<image-name>:linux -f linux/Dockerfile .
```

#### Windows

```cmd
docker build -t <sample-name>-<image-name>:windows -f windows/Dockerfile .
```

Pre-built images for all samples can be found on the docker hub, under the account **rajchaudhuri**.

### Docker Compose files

If a sample includes docker compose files, these will be in a directory called `docker` under the sample directory. There may be separate compose files for linux and windows, in which case they will be in `linux` and `windows` subdirectories. If a compose file is found directly under the `docker` directory, it is safe to use unmodified in linux and windows.

The compose files will always be named for the sample, and never be called ```docker-compose.yml```. This means that the `docker-compose` command will always need to be used with the `-f` option.

All compose files will be suitable for both standalone docker deployment with `docker-compose`, or for swarm deployment with `docker stack create -c`. In case a sample supports only standalone or only swarm, this will be called out in the sample's documentation.

### Kubernetes manifest files

If a sample includes kubernetes manifest files, these will be in a directory called `kubernetes` under the sample directory. There may be separate manifest files for linux and windows, in which case they will be in `linux` and `windows` subdirectories. If manifest files are found directly under the `kubernetes` directory, it is safe to use them unmodified in linux and windows.

### Modifying and building sample images

All docker compose and kubernetes manifest files in all samples initially reference the pre-built images. If you modify and re-build the images, remember to publish to your own container registry, and change the references.

## List of samples

|Sample Name|Description|
|---|---|---|
|[inside-container](inside-container.md)|This sample is a simple web server application, which can show run time information such as host name, ip addresses, environment variables and file system of the container it is run in.|
|[service-connectivity](service-connectivity.md)|This sample demostrates how an application (the client) can discover another application (the service) in the same docker host, docker swarm, or kubenetes cluster using service discovery.|

More coming.