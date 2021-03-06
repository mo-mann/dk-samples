---
layout: default
title: Service-Connectivity Samples
---

This sample demostrates how an application (the _client_) can discover another application (the _server_) in the same docker host, docker swarm, or kubenetes cluster using _service discovery_.

In this sample, the **server** is a web server, built from the [inside-container](inside-container.md) sample. The **client** is a shell script, built from a custom image whose Dockerfile is provided as part of the sample.

## The client image

This image runs a shell script that continuously performs an http get on the url contained in an environment variable called SERVICE_URL, and logs the results.

## How to Use the image by itself

Create a container using:

```bash
docker container run -d --name t1 -e SERVICE_URL=http://www.google.com rajchaudhuri/service-connectivity-tester
```

Then, check the logs with:

```bash
docker container logs t1
```

A pre-built version of the image, tagged as `rajchaudhuri/service-connectivity-tester` is available on the docker hub.

The docker compose file and the kubernetes manifest file both use this image, and configure the SERVICE_URL environment variable using the name of the docker or kubernetes service.

## Standalone Docker

In standalone docker, the client should discover the server using the name of the container that runs the server application, or an alias. This can be enabled by using the `--link` option of the `docker create` command while creating the client container, or by putting both containers on a private network.

The supplied docker compose file (`service-connectivity-docker.yml`) does exactly this when deployed using the `docker-compose` tool. It uses the name `sample-server` to define the server container. In the compose file, the `SERVER_URL` environment variable for the client container is set to point to that name.

Deploy it with the command:

```docker-compose -f service-connectivity-docker.yml up -d```

You can then check that the client can connect to the server by using the command:

```docker-compose -f service-connectivity-docker.yml logs servicetester```

## Docker Swarm Mode

In docker swarm mode, the client and the server are each deployed as _swarm services_. Each can be scaled to run multiple containers.

All containers that are part of the client swarm service should discover containers in the server swarm service using the _name_ of the server swarm service. Docker ensures that the name resolves to any container that belongs to the server swarm service, and even performs load balancing.

The supplied docker compose file (`service-connectivity-docker.yml`) causes exactly this to happen, when deployed as a docker stack. Deploy it with the command:

```docker stack deploy -f service-connectivity-docker.yml teststack```

The server can be scaled using the command:

```docker service scale teststack_sample-server=3```

and the client can be scaled with:

```docker service scale teststack_sample-client=2```

You can then check that _any_ **client** can **connect** to _any_ server by using the command:

```docker service logs teststack_sample-client```

## Kubernetes

In kubernetes, the client and the server are each deployed via kubernetes _deployments_. Each deployment can be scaled to run multiple kubernetes _pods_, which encapsulate containers. For client pods to discover server pods, kubernetes mandates that we define a _kubernetes service_.

A service in kubenetes has a name, and is assigned an IP address. It also has a _selector_ that lets it identify pods which contain the actual service application. In our case, the service points to the server pods.

Pods in the client deployment should discover pods in the server deployment by using the _name of the service_. Kubernetes ensures that the name resolves to any pod that the service points to, and routes the request accordingly.

The supplied kubernetes manifest file (`service-connectivity-k8s.yml`) causes exactly this to happen when deployed. In the manifest file, the service is defined using the name `service-connectivity-service`. The `SERVER_URL` environment variable for the client container is set to point to that name.

Deploy it with the command:

```kubectl apply -f service-connectivity-k8s.yml```

The server can be scaled using the command:

```kubectl scale deployment service-connectivity-server --replicas=3```

and the client can be scaled with:

```kubectl scale deployment service-connectivity-client --replicas=2```

You can then check that _any_ client can connect to _any_ server by using the command:

```kubectl logs -lapp=service-connectivity-client```
