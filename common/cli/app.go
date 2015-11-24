package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
    "github.com/codegangsta/cli"
    "github.com/fsamin/intools-engine/common/logs"
    "github.com/fsamin/intools-engine/common/server"
    "github.com/fsamin/intools-engine/common/utils"
    "github.com/fsamin/intools-engine/connectors"
    "github.com/fsamin/intools-engine/intools"
    "github.com/fsamin/intools-engine/groups")

func initLoggers(c *cli.Context) {
	var debugLogger io.Writer
	var flag int
	if c.GlobalBool("debug") {
		debugLogger = os.Stdout
		flag = log.Ldate | log.Ltime | log.Lshortfile
	} else {
		debugLogger = ioutil.Discard
		flag = log.Ldate | log.Ltime
	}
	logs.InitLog(debugLogger, os.Stdout, os.Stdout, os.Stderr, flag)
}

func daemonAction(c *cli.Context) {
	initLoggers(c)
	port := c.GlobalInt("port")
	debug := c.GlobalBool("debug")
	logs.Trace.Println("Starting Intools-Engine as daemon")

	dockerClient, dockerHost, err := utils.GetDockerCient(c)
	if err != nil {
		os.Exit(1)
	}

	redisClient, err := utils.GetRedisClient(c)
	if err != nil {
		os.Exit(1)
	}

	d := server.NewDaemon(port, debug, dockerClient, dockerHost, redisClient)
	d.SetRoutes()
	d.Run()
}

func runAction(c *cli.Context) {
	initLoggers(c)

	dockerClient, host, err := utils.GetDockerCient(c)
	if err != nil {
		os.Exit(1)
	}

	redisClient, err := utils.GetRedisClient(c)
	if err != nil {
		os.Exit(1)
	}

	cmd := []string{c.Args().First()}
	cmd = append(cmd, c.Args().Tail()...)
	if len(cmd) < 4 {
		logs.Error.Println("Incorrect usage, please check --help")
		return
	}
	group := cmd[0]
	conn := cmd[1]
	image := cmd[2]
	t := cmd[3]
	timeout, err := strconv.Atoi(t)
	if err != nil {
		// handle error
		logs.Error.Println(err)
		os.Exit(2)
	}
	cmd = cmd[4:]
	logs.Debug.Println("Launching " + image + " " + strings.Join(cmd, " "))
	logs.Warning.Printf("In command line, connector schedule is not available")
	intools.Engine = &intools.IntoolsEngineImpl{dockerClient, host, redisClient, nil}
	connector := connectors.NewConnector(group, conn)
	connector.Init(image, timeout, 0, cmd)
	groups.CreateGroup(group)
    if err != nil {
        os.Exit(3)
    }
	executor, err := connectors.Exec(connector)
	if err != nil {
		os.Exit(3)
	}
	fmt.Println(executor.GetJSON())

}

func testAction(c *cli.Context) {
	logs.Error.Println("Not yet implemented")
}

func publishAction(c *cli.Context) {
	logs.Error.Println("Not yet implemented")
}