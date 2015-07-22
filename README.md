# intools-engine

[![wercker status](https://app.wercker.com/status/f3795ac31ee708a4ca07500d98870470/m "wercker status")](https://app.wercker.com/project/bykey/f3795ac31ee708a4ca07500d98870470)

## Development Dependencies
````
go get -u "github.com/codegangsta/cli"
go get -u "github.com/gin-gonic/gin"
go get -u "github.com/robfig/cron"
go get -u "github.com/samalba/dockerclient"
go get -u "gopkg.in/redis.v3"
go get -u github.com/gorilla/websocket
````

## Environment Dependencies
 - Docker Host (version 1.5 min)
 - Redis Host (version 2.8 min)

## Global Options
````
 --host, -H                   Docker host [$DOCKER_HOST]
 --cert, -C 			      Docker cert path [$DOCKER_CERT_PATH]
 --redis "localhost:6379"     Redis Host [$REDIS_HOST]
 --redis-password             Redis Password [$REDIS_PWD]
 --redis-db "0"               Redis Database [$REDIS_DB]
 --debug 			          Debug mode [$INTOOLS_DEBUG]
````
