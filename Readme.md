# node+docker

Node+Docker includes Node and the docker-engine to ease CI/CD by enabling the 
working with the an external/mapped docker host for building Docker images.

Please note: This container does not set out to be secure, it sets out to be
usable as a tool in a self-contained private CI/CD setting.

# Quickstart
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

# Development
Each node/docker version is split out into a separate branch.

## Node build tools
For those that are using anything that requires node-gyp (SASS, argon2 e.t.c.)
you will need to install the required build tools. 

### dockerfile
```
FROM matthewhartstonge/node-docker:node6-docker1.12

RUN apt-get update \
    && apt-get install \
        g++ \
        make \
        python \
    && npm install \
    && apt-get remove --purge \
        g++ \
        make \
        python
```

### CI Script
```sh
script:
  - apt-get update && apt-get install g++ make python
  - npm install
  - npm test
```

# Deployment
Simply pull your required version from dockerhub

```sh
docker pull matthewhartstonge/node-docker
```

# Testing
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

Then in your .drone, .travis, .jenkinsfile, .whatever add: 

```yml
script:
  - npm run test 
```
