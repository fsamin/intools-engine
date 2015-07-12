package main

import "github.com/codegangsta/cli"
import "io/ioutil"
import "os"
import "io"
import "strings"
import "strconv"
import "log"
import "fmt"

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
	InitLog(debugLogger, os.Stdout, os.Stdout, os.Stderr, flag)
}

func daemonAction(c *cli.Context) {
	initLoggers(c)
	port := c.GlobalInt("port")
	debug := c.GlobalBool("debug")
	Trace.Println("Starting Intools-Engine as daemon")

	dockerClient, dockerHost, err := getDockerCient(c)
	if err != nil {
		os.Exit(1)
	}

	redisClient, err := getRedisClient(c)
	if err != nil {
		os.Exit(1)
	}

	d := newDaemon(port, debug, dockerClient, dockerHost, redisClient)
	d.setRoutes()
	d.run()
}

func runAction(c *cli.Context) {
	initLoggers(c)

	dockerClient, host, err := getDockerCient(c)
	if err != nil {
		os.Exit(1)
	}

	redisClient, err := getRedisClient(c)
	if err != nil {
		os.Exit(1)
	}

	cmd := []string{c.Args().First()}
	cmd = append(cmd, c.Args().Tail()...)
	if len(cmd) < 4 {
		Error.Println("Incorrect usage, please check --help")
		return
	}
	group := cmd[0]
	conn := cmd[1]
	image := cmd[2]
	t := cmd[3]
	timeout, err := strconv.Atoi(t)
	if err != nil {
		// handle error
		Error.Println(err)
		os.Exit(2)
	}
	cmd = cmd[4:]
	Debug.Println("Launching " + image + " " + strings.Join(cmd, " "))
	Warning.Printf("In command line, connector schedule is not available")
	Intools = &IntoolsEngine{*dockerClient, host, *redisClient, nil}
	connector := Intools.NewConnector(group, conn)
	connector.Init(image, timeout, 0, cmd)
	Intools.InitSchedule(connector)
	executor, err := Intools.Exec(connector)
	if err != nil {
		os.Exit(3)
	}
	fmt.Println(executor.GetJSON())

}

func testAction(c *cli.Context) {

}

func publishAction(c *cli.Context) {

}
