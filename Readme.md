# node+docker

Node+Docker includes Node and the docker-engine to ease CI/CD by enabling the 
working with the an external/mapped docker host for building Docker images.

Please note: This container does not set out to be secure, it sets out to be
usable as a tool in a self-contained private CI/CD setting.

## Supported tags
- `latest`, `node6-docker1.12.2` [(node6-docker1.12.2/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.12.2/Dockerfile)
- `node6-docker1.12.1` [(node6-docker1.12.1/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.12.1/Dockerfile)
- `node6-docker1.12.0` [(node6-docker1.12.0/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.12.0/Dockerfile)
- `node6-docker1.11.2` [(node6-docker1.11.2/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.11.2/Dockerfile)
- `node6-docker1.11.1` [(node6-docker1.11.1/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.11.1/Dockerfile)
- `node6-docker1.11.0` [(node6-docker1.11.0/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.11.0/Dockerfile)
- `node6-docker1.10.3` [(node6-docker1.10.3/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.10.3/Dockerfile)
- `node6-docker1.10.2` [(node6-docker1.10.2/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.10.2/Dockerfile)
- `node6-docker1.10.1` [(node6-docker1.10.1/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.10.1/Dockerfile)
- `node6-docker1.10.0` [(node6-docker1.10.0/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.10.0/Dockerfile)
- `node6-docker1.9.1` [(node6-docker1.9.1/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.9.1/Dockerfile)
- `node6-docker1.9.0`[(node6-docker1.9.0/Dockerfile)](https://github.com/MatthewHartstonge/node-docker/blob/node6-docker1.9.0/Dockerfile)

## Quickstart
Since this image is based on dockerhub's [node](https://hub.docker.com/_/node/),
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
