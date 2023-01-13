# nullbot

This is a little discord bot for running/compiling code. The code is run through the docker's temporary containers.

```sh
docker build -t gcc-compile ./tools/gcc
docker build -t rust-compile ./tools/rust
docker pull python


go run cmd/app/main.go -token <bot-token> -guild <guild id> -rmcmd <bool>
```
