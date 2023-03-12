# ethproxy
Small golang service that proxifies the getLastBlock RPC from an Ethereum execution node

# Compile

Simply use `make compile`, it will launch a docker image, and compile the source code.

The binary `eth-proxy` freshly built, is in the `build` directory

# Quality

The `make test` will run obviously the unit test, and `make lint` the golang-ci linter.

This steps will produce report for `sonarqube` and `gitlab-ci`.

# Build

In order to build the docker image, use `make build`, the tag of the new image is based on the git commit hash.

This image is not pushed, if you want so, first you have to duplicate the `.docker_hub_token.tmpl` for `.docker_hub_token`
and fill it with your docker hub token, and change the docker registry in the Makefile `DOCKER_REGISTRY` for yours.

# Run
First, you will need an Ethplorer Api Key, if you don't have one please visit : [Ethplorer](https://github.com/EverexIO/Ethplorer/wiki/Ethplorer-API#api-key-limits).

Then, create a `config.json` in the `resources/config/` directory, based on the template `resources/config/template/config.json`.

After compiling the code, you can run the binary, or use the docker image `make start`.

You could pull the image `make pull` or build it locally.

With Curl : 
`curl http://localhost:9090/eth/lastblock`

The metrics are avaible at this url : `http://localhost:9090/metrics`

The liveness, and readyness are availble at this url : `http://localhost:9090/healthz`


# To do :
Generate a unique identifier for each Ethplorer request, and check if the identifier received in the response is equal, returns an error otherwise.

Fix the use of Gin Gonic framework when the user send a signal term, there's alway an error message `Error while starting the web server`.
Create a go builder image to avoid install and update during build process.

Github action scripts.

Push sonarqube report (must install server somewhere).

Improve security (store api key in k8s secret manager).

Set githook, with pre-commit and gitlint, to force the commit message to respect a regular expression