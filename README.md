# node+docker

Node+Docker includes Node and the docker-engine to ease CI/CD by enabling the 
working with the an external/mapped docker host for building Docker images.

Please note: This container does not set out to be secure, it sets out to be
usable as a tool in a self-contained private CI/CD setting.

## Docker tags
The Docker tag is in the format: `node${NODE_VERSION}-docker${DOCKER_VERSION}` 
where node version and the docker version is replaced with one of the following
versions:  

| app       | supported versions         |
|-----------|----------------------------|
| node      | `6`, `8`, `9`, `10`        |
| docker    | `1.5.0`, `1.6.0`, `1.6.1`, `1.6.2`, `1.7.0`, `1.7.1`, `1.8.0`, `1.8.1`, `1.8.2`, `1.8.3`, `1.9.0`, `1.9.1`, `1.10.0`, `1.10.1`, `1.10.2`, `1.10.3`, `1.11.0`, `1.11.1`, `1.11.2`, `1.12.0`, `1.12.1`, `1.12.2`, `1.12.3`, `1.12.4`, `1.12.5`, `1.12.6`, `1.13.0`, `1.13.1` |
| docker-ce | `17.03.0ce`, `17.03.1ce`, `17.04.0ce`, `17.05.0ce` |

If you are after a specific Dockerfile, you can view the [Github Repo](https://github.com/matthewhartstonge/node-docker), 
where each version tag has been pushed to a separate branch.

### Example

Say I want, Node 8 with Docker CE 17.03, you would run:
```sh
docker pull matthewhartstonge/node-docker:node8-docker17.03.1ce
```

## Quickstart
Since this image is based on dockerhub's [node](https://hub.docker.com/r/_/node/),
all actions follow through on this container, with the exception of the 
addition of the docker-engine.

The underlying OS is debian:jessie. 

To get access to your host, simply map the docker socket into the container on
startup.

```sh
docker run -it \
    -v "/var/run/docker.sock:/var/run/docker.sock"
    matthewhartstonge/node-docker
    docker ps
```

## Development
Each node/docker version is split out into a separate branch. Please feel free 
to add Pull Requests to add the different versions you use as a token of thanks
and also to give back to the community. 

### Node build tools
For those that are using anything that requires node-gyp (SASS, argon2 e.t.c.)
the build tools (g++, make and python) are now included.

## Deployment
Simply pull your required version from dockerhub

```sh
docker pull matthewhartstonge/node-docker
```

## Testing
To use CI testing with this image, depending on how your CI environment works, 
the simplest way to do this is to add an NPM script that can run. 

For example, using Mocha, in package.json:

```
...
  },
  "scripts": {
    "start": "node app.js",
    "test": "node ./node_modules/mocha/bin/mocha test/**/*.test.js"
  },
...
```

Then in your .drone, .jenkinsfile, .whatever add:

```yml
script:
  - npm run test 
```
