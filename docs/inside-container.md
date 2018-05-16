---
layout: default
title: Inside-Container Sample
---

This image is a simple web server application, which can show the host name, ip addresses, environment variables and file system of the container it is run in.

The web server listens on port 8080 by default. The number can be changed using an environment variable called PORT_NUMBER.

## How to Use

Create a container using:

```bash
docker container run -d -e PORT_NUMBER=8070 -p 8070:8070 rajchaudhuri/inside-container-sample
```

Then, access `http://localhost:8070` to see the container's hostname and IP addresses. Access `http://localhost:8070/env` to see its environment variables, and `http://localhost:8070/files` to see the file system.

If you want to explore a particular path, say **/proc**, append that to the request like this: `http://localhost:8070/files/proc`. By default, the web server traverses subdirectories one level deep. You can go deeper by providing a query parameter called **depth**, like so: `http://localhost:8070/files/proc?depth=3`.