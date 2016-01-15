package cli

import "github.com/codegangsta/cli"

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "intools"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port, P",
			Value:  8082,
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
		cli.StringFlag{
			Name:   "registry-username",
			Usage:  "Docker Registry Username",
			Value:  "",
			EnvVar: "DOCKER_REGISTRY_USER",
		},
		cli.StringFlag{
			Name:   "registry-password",
			Usage:  "Docker Registry Password",
			Value:  "",
			EnvVar: "DOCKER_REGISTRY_PWD",
		},
		cli.StringFlag{
			Name:   "registry-mail",
			Usage:  "Docker Registry Mail",
			Value:  "",
			EnvVar: "DOCKER_REGISTRY_MAIL",
		},
		cli.StringFlag{
			Name:   "registry-token",
			Usage:  "Docker Registry Token",
			Value:  "",
			EnvVar: "DOCKER_REGISTRY_TOKEN",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "daemon",
			Usage:       "Run intools engine as a Daemon",
			Description: "Daemon",
			Action:      daemonAction,
		},
		cli.Command{
			Name:        "test",
			Usage:       "Test your connector over intools engine",
			Description: "Test",
			Action:      testAction,
		},
		cli.Command{
			Name:        "run",
			Usage:       "Run your connector over intools engine",
			Description: "Run",
			Action:      runAction,
		},
		cli.Command{
			Name:        "publish",
			Usage:       "Publish your connector on intools engine",
			Description: "Daemon",
			Action:      publishAction,
		},
	}

	return app
}
