# intools-engine

[![wercker status](https://app.wercker.com/status/f3795ac31ee708a4ca07500d98870470/m "wercker status")](https://app.wercker.com/project/bykey/f3795ac31ee708a4ca07500d98870470)

## Development Dependencies
````
go get -u "github.com/codegangsta/cli"
go get -u "github.com/gin-gonic/gin"
go get -u "gopkg.in/robfig/cron.v2"
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

## How to use
### Command line
 - Run the server
````
 ./intools-engine daemon
````

 - Run a connector (it will register a group, a connector, its configuration, and run it once)
     - Argument to pass are : docker image name, timeout for container execution, and the container commandline

````
./intools-engine --redis-db 1 -H unix:///var/run/docker.sock run CDK helloworld debian:jessie 5 echo '{"value":"test value"}'
[INTOOLS] [INFO]  2015/11/24 15:32:08 Connected to Docker Host unix:///var/run/docker.sock
[INTOOLS] [INFO]  2015/11/24 15:32:08 Connected to Redis Host localhost:6379/1
[INTOOLS] [WARN]  2015/11/24 15:32:08 In command line, connector schedule is not available
[INTOOLS] [INFO]  2015/11/24 15:32:09 71ec23a7acb [/CDK-helloworld] successfully created
[INTOOLS] [INFO]  2015/11/24 15:32:09 71ec23a7acb [/CDK-helloworld] successfully started
{
   "ContainerId":"71ec23a7acb",
   "Host":"unix:///var/run/docker.sock",
   "Running":false,
   "Terminated":true,
   "ExitCode":0,
   "Stdout":"{\"value\":\"test value\"}",
   "JsonStdout":{
      "value":"test value"
   },
   "Stderr":"",
   "StartedAt":"2015-11-24T14:32:09.337306123Z",
   "FinishedAt":"2015-11-24T14:32:09.383803882Z",
   "Valid":true
}

````

### REST Api
#### Groups
 - Return all groups as a list
````
 GET <host:port>/groups
````
 - Create a group
````
 POST <host:port>/groups/:group
````
 - Get a group
````
 GET <host:port>/groups/:group
````
  - Returns
````
     {
         "name": "CDK"
     }
````

 - Delete the specific group
````
 DELETE <host:port>/groups/:group
````

#### Connectors
 - Connector JSON Structure
````
    {
        "group": "CDK",
        "name": "helloworld",
        "config": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": null,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": null,
            "Cmd": [
                "echo",
                "{\"value\":\"test value\"}"
            ],
            "Image": "debian:jessie",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "NetworkDisabled": false,
            "MacAddress": "",
            "OnBuild": null,
            "Labels": null,
            "StopSignal": "",
            "VolumeDriver": "",
            "Memory": 0,
            "MemorySwap": 0,
            "CpuShares": 0,
            "Cpuset": "",
            "PortSpecs": null,
            "HostConfig": {
                "Binds": null,
                "ContainerIDFile": "",
                "LxcConf": null,
                "Memory": 0,
                "MemoryReservation": 0,
                "MemorySwap": 0,
                "KernelMemory": 0,
                "CpuShares": 0,
                "CpuPeriod": 0,
                "CpusetCpus": "",
                "CpusetMems": "",
                "CpuQuota": 0,
                "BlkioWeight": 0,
                "OomKillDisable": false,
                "MemorySwappiness": 0,
                "Privileged": false,
                "PortBindings": null,
                "Links": null,
                "PublishAllPorts": false,
                "Dns": null,
                "DNSOptions": null,
                "DnsSearch": null,
                "ExtraHosts": null,
                "VolumesFrom": null,
                "Devices": null,
                "NetworkMode": "",
                "IpcMode": "",
                "PidMode": "",
                "UTSMode": "",
                "CapAdd": null,
                "CapDrop": null,
                "GroupAdd": null,
                "RestartPolicy": {
                    "Name": "",
                    "MaximumRetryCount": 0
                },
                "SecurityOpt": null,
                "ReadonlyRootfs": false,
                "Ulimits": null,
                "LogConfig": {
                    "type": "",
                    "config": null
                },
                "CgroupParent": "",
                "ConsoleSize": [
                    0,
                    0
                ],
                "VolumeDriver": ""
            }
        },
        "timeout": 5,
        "refresh": 60
    }

````

 - Get all connectors
````
 GET <host:port>/groups/:group/connectors
````
Return a JSON list of connectors as above

 - Create a connector
````
 POST <host:port>/groups/:group/connectors/:connector
````
Post the JSON object as above to create a connector

 - Get a connector
````
 GET <host:port>/groups/:group/connectors/:connector
````
Returns the JSON object

- Delete a connector
````
DELETE <host:port>/groups/:group/connectors/:connector
````

 - Force a connector refresh
````
 GET <host:port>/groups/:group/connectors/:connector/refresh
````
Force connector execution and return the detail of the container execution as `GET :host:port/groups/:group/connectors/:connector/exec`

 - Get the last result of a connector
````
 GET <host:port>/groups/:group/connectors/:connector/result
````
Return the last JSONStdout of a connector
````
    {
        "value": "test value"
    }
````

 - Get the last executor of a connector
````
 GET <host:port>/groups/:group/connectors/:connector/exec
````
Return the detail of a container execution
````
{
    "ContainerId": "71ec23a7acb",
    "Host": "unix:///var/run/docker.sock",
    "Running": false,
    "Terminated": true,
    "ExitCode": 0,
    "Stdout": "{\"value\":\"test value\"}",
    "JsonStdout": {
        "value": "test value"
    },
    "Stderr": "",
    "StartedAt": "2015-11-24T14:32:09.337306123Z",
    "FinishedAt": "2015-11-24T14:32:09.383803882Z",
    "Valid": true
}
````

## Tests
### Install Ginkgo
````
 go get github.com/onsi/ginkgo/ginkgo
 go get github.com/onsi/gomega
````

This fetches ginkgo and installs the ginkgo executable under `$GOPATH/bin` – you’ll want that on your `$PATH`.

### Run Tests suites
````
ginkgo groups
````
