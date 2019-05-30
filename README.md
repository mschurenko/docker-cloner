# docker-cloner
[![release](http://img.shields.io/github/release/mschurenko/docker-cloner.svg?style=flat-square)](https://github.com/mschurenko/docker-cloner/releases)
[![CircleCI](https://circleci.com/gh/mschurenko/docker-cloner.svg?style=svg)](https://circleci.com/gh/mschurenko/docker-cloner)

This small Go program will "clone" a running docker container by fetching the running container's metadata and then creating a new container. The new container id will be written to stdout. If there are any errors an exit code of `1` will be returned. A few common attributes (command, entrypoint, etc can be overriden). Run `docker-cloner -h` for more info.

The new container will be in a stopped state (use `docker ps -a` to see it). Use `docker start <new_container_id>` to start it.

## Example Usage
```sh
container_id=$(docker-cloner -memory 200 -id <running_container_name> -cmd "echo arg1 arg2" -new_name <clone_of_running_container>)
docker start -ia $container_id
```
