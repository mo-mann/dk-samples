---
layout: default
title: Docker and Kubernetes Samples
---

This sample demostrates how an application (the _client_) can discover another application (the _server_) in the same docker host, docker swarm, or kubenetes cluster using _service discovery_.

In this sample, the **server** is a web server, built from the stock `nginx` image from the Docker Hub. The **client** is a shell script, built from a custom image whose Dockerfile is provided as part of the sample.

## Standalone Docker

In standalone docker, the **client** should discover the **server** using the name of the container that runs the **server** application, or an alias. This can be enabled by using the `--link` option of the `docker create` command while creating the **client** container, or by putting both containers on a private network.

The supplied docker compose file (`service-connectivity-docker.yml`) does exactly this when deployed using the `docker-compose` tool. The command is:

```docker-compose -f service-connectivity-docker.yml up -d```

You can then check that the **client** can **connect** to the server by using the command:

```docker-compose -f service-connectivity-docker.yml logs servicetester```

## Docker Swarm Mode

In docker swarm mode, the **client** and the **server** are each deployed as _swarm services_. Each can be scaled to run multiple containers. 

All containers that are part of the **client** swarm service should discover containers in the **server** swarm service using the name of the swarm service. Docker ensures that the name resolves to any container that belongs to the **server** swarm service, and even performs load balancing.

The supplied docker compose file (`service-connectivity-docker.yml`) causes exactly this to happen, when deployed as a docker stack. The command is:

```docker stack deploy -f service-connectivity-docker.yml teststack```

The **server** can be scaled using the command

```docker service scale teststack_sample-server```

You can then check that _any_ **client** can **connect** to _any_ server by using the command:

```docker service logs teststack_sample-client```

## Kubernetes

In kubernetes, the **client** and the **server** are each deployed via kubernetes _deployments_. Each deployment can be scaled to run multiple kubernetes _pods_, which encapsulate containers. For **client** pods to discover **server** pods, kubernetes mandates that we define a _kubernetes service_.

A service in kubenetes has a name, and is assigned an IP address. It also has a _selector_ that lets it identify pods which contain the actual service application. In our case, the service points to the **server** pods.

Pods in the **client** deployment should discover pods in the **server** deployment by using the name of the service. Kubernetes ensures that the name resolves to any pod that the service points to, and routes the request accordingly.

The supplied kubernetes manifest file (`service-connectivity-k8s.yml`) causes exactly this to happen when deployed. The command is:

```kubectl apply -f service-connectivity-k8s.yml```

You can then check that _any_ **client** can **connect** to _any_ server by using the command:

```kubectl logs -lapp=service-connectivity-client```

## The client image

The client image runs a simple shell script that performs an http get on the url contained in an environment variable called SERVICE_URL, and logs the result.

A pre-built version of the image, tagged as `rajchaudhuri/service-connectivity:alpine` is available on the docker hub.

The docker compose file and the kubernetes manifest file both use this image, and configure the SERVICE_URL using the service name.

