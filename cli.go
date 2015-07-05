package main

import "github.com/codegangsta/cli"

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "intools"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port, P",
			Value:  8080,
			Usage:  "Intools engine daemon port",
			EnvVar: "INTOOLS_PORT",
		},
		cli.StringFlag{
			Name:   "host, H",
			Usage:  "Docker host",
			EnvVar: "DOCKER_HOST",
		},
		cli.StringFlag{
			Name:   "cert, C",
			Usage:  "Docker cert path",
			EnvVar: "DOCKER_CERT_PATH",
		},
		cli.StringFlag{
			Name:   "redis",
			Usage:  "Redis Host",
			Value:  "localhost:6379",
			EnvVar: "REDIS_HOST",
		},
		cli.StringFlag{
			Name:   "redis-password",
			Usage:  "Redis Password",
			Value:  "",
			EnvVar: "REDIS_PWD",
		},
		cli.IntFlag{
			Name:   "redis-db",
			Usage:  "Redis Database",
			Value:  0,
			EnvVar: "REDIS_DB",
		},
		cli.StringFlag{
			Name:   "debug",
			Usage:  "Debug mode",
			EnvVar: "INTOOLS_DEBUG",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "daemon",
			Aliases:     []string{"d"},
			Usage:       "Run intools engine as a Daemon",
			Description: "Daemon",
			Action:      daemonAction,
		},
		cli.Command{
			Name:        "test",
			Aliases:     []string{"t"},
			Usage:       "Test your connector over intools engine",
			Description: "Test",
			Action:      testAction,
		},
		cli.Command{
			Name:        "run",
			Aliases:     []string{"r"},
			Usage:       "Run your connector over intools engine",
			Description: "Run",
			Action:      runAction,
		},
		cli.Command{
			Name:        "publish",
			Aliases:     []string{"p"},
			Usage:       "Publish your connector on intools engine",
			Description: "Daemon",
			Action:      publishAction,
		},
	}

	return app
}
