# ƛ📦 aws-lambda-container

Example Go AWS lambda packaged in a Docker container


## Description

This repo contains an example Go application that can run as a Dockerkised lambda both on AWS and locally.

### Why not using the aws base image?

Because it's huge! 

The official image (public.ecr.aws./lambda/go) provided by AWS ships a whole Linux distro and weights about 670MB, which is definitely too much.

### Distroless image

This container uses the images provided by the Google Container Tools here: https://github.com/GoogleContainerTools/distroless

In particular it uses the gcr.io/distroless/static image which is ideal for statically compiled languages like Go.

The resulting image weights about 10MB.


## Local execution

So the whole reason for this being interesting is that it allows developers to improved their feedback loop when working with lambdas, without having to use external tools like SAM.

**But**, in order to use a custom image, you need to either bake into the image the `aws-lambda-runtime-interface-emulator` or install it on the host machine and point the Docker entrypoint to that executable.

My approach here is to bake it into the local test image, following these steps:
 - build the app and copy it into a distroless image (multi-stage docker build)
   - this is the lambda image that can be pushed to ECR
 - build a test image that uses the lambda image as a base layer and uses the `aws-lambda-rie` as entrypoint. 


## How can I test this lambda locally?
1. Package the application:

```
make lambda
```

This will build the docker image and tag it as `fastfib:latest`

2. Package the test image that will be executed locally

```
make testlambda
```

This will build the docker image and tag it as `testfastfib:latest`

3. Run the lambda locally mapping port 8080 of the container to 9000 on the host machine

```
docker run -p 9000:8080 testfastfib:latest
```

4. Send a request to the lambda

The code itself implements a fast fibonacci sequence algorithm based on https://www.nayuki.io/page/fast-fibonacci-algorithms

So the lambda will reply with the n-th element of the Fibonacci sequence in JSON content encoding.

Request:
```
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"n":7}'
```
Response:
```
{"result":13}
```

## TODO
- [ ] implement an API Gateway Proxy request/response handler
