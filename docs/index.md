---
layout: default
title: Docker and Kubernetes Samples
---

This repository contains samples used to demonstrate concepts of docker standalone, docker swarm mode, and kubernetes.

For all samples, we will try to provide versions for Linux and Windows, with the Linux samples being created first. For all orchestration samples, we will try to provide versions on docker swarm mode and kubernetes, with the swarm mode samples being created first.

We try to follow git flow. The master branch contains "final" code; active development happens in the "development" branch.

Contributions are welcome. Feel free to fork and send pull requests.

## How to use the samples

If a sample contains images, the sources will be in a directory called `images` under the sample directory. Each subdirectory under images represents an individual image, which should be tagged as _sample-name_-_image-name_. Under each image directory, there will be linux and windows subdirectories, which will contain Dockerfiles.

To build an image, the image directory must be used as the build context directory. This means that you build Linux and Windows images separately, as follows:

**Linux** : **_image-directory_$** docker build -t _sample-name_-_image-name_:linux -f linux/Dockerfile .

**Windows** : **C:\_image-directory_>** docker build -t _sample-name_-_image-name_:windows -f windows/Dockerfile .

Pre-built images can be found on the docker hub, under the account **rajchaudhuri**.

Any docker compose and kubernetes manifest files reference the pre-built images. If you modify and re-build the images, remember to publish to your own container registry, and change the references.

## List of samples

|Sample Name|Description|
|---|---|---|
|[inside-container](inside-container.md)|This sample is a simple web server application, which can show run time information such as host name, ip addresses, environment variables and file system of the container it is run in.|
|[service-connectivity](service-connectivity.md)|This sample demostrates how an application (the client) can discover another application (the service) in the same docker host, docker swarm, or kubenetes cluster using service discovery.|

More coming.