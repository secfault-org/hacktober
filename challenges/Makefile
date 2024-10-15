DOCKER_NAME=hacktober_base


all: dockerfile

dockerfile: Dockerfile entrypoint.sh
	podman build -t $(DOCKER_NAME) .
