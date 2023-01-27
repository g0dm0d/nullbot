# compiler plugin for discord bot

This is a little plugin for running/compiling code. The code is run through the docker's temporary containers.

```sh
docker build -t gcc-compile ./gcc
docker build -t rust-compile ./rust
docker pull python
```
