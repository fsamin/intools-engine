package utils

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/codegangsta/cli"
	"github.com/samalba/dockerclient"
	"github.com/fsamin/intools-engine/common/logs"
	"gopkg.in/redis.v3"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	redisOptions *redis.Options
)

func StringTransform(s string) string {
	v := make([]rune, 0, len(s))
	for i, r := range s {
		if r == utf8.RuneError {
			_, size := utf8.DecodeRuneInString(s[i:])
			if size == 1 {
				continue
			}
		}
		//check unicode chars
		if !unicode.IsControl(r) {
			v = append(v, r)
		}
	}
	s = string(v)

	return strings.TrimSpace(s)
}

func ReadLogs(reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
	var text string
	for scanner.Scan() {
		logs.Debug.Println(scanner.Text())
		text += StringTransform(scanner.Text() + "\n")
	}
	err := scanner.Err()
	if err != nil {
		logs.Error.Println("There was an error with the scanner", err)
	}
	return text, err
}

func GetRedis(c *cli.Context) (*redis.Client, error) {
	redisOptions = &redis.Options{
		Addr:     c.GlobalString("redis"),
		Password: c.GlobalString("redis-password"),
		DB:       int64(c.GlobalInt("redis-db")),
	}

	client, err  := GetRedisClient()

	logs.Trace.Printf("Connected to Redis Host %s/%d", c.GlobalString("redis"), c.GlobalInt("redis-db"))
	return client, err
}

func GetRedisClient() (*redis.Client, error) {
	client := redis.NewClient(redisOptions)

	_, err := client.Ping().Result()
	if err != nil {
		logs.Error.Println("Unable to connect to redis host")
		return nil, err
	}

	return client, nil
}

func GetDockerCient(c *cli.Context) (*dockerclient.DockerClient, string, error) {
	host := c.GlobalString("host")
	if host == "" {
		logs.Error.Println("Incorrect usage, please set the docker host")
		return nil, "", errors.New("Unable to connect to docker host")
	}

	tlsConfig := &tls.Config{}

	certPath := c.GlobalString("cert")
	if certPath != "" {
		caFile := filepath.Join(certPath, "ca.pem")
		if _, err := os.Stat(caFile); os.IsNotExist(err) {
			logs.Error.Println("Cannot open file : " + caFile)
			logs.Error.Println("Incorrect usage, please set correct cert files")
			return nil, host, errors.New("Unable to connect to docker host")
		}

		certFile := filepath.Join(certPath, "cert.pem")
		if _, err := os.Stat(certFile); os.IsNotExist(err) {
			logs.Error.Println("Cannot open file : " + certFile)
			logs.Error.Println("Incorrect usage, please set correct cert files")
			return nil, host, errors.New("Unable to connect to docker host")
		}

		keyFile := filepath.Join(certPath, "key.pem")
		if _, err := os.Stat(keyFile); os.IsNotExist(err) {
			logs.Error.Println("Cannot open file : " + keyFile)
			logs.Error.Println("Incorrect usage, please set correct cert files")
			return nil, host, errors.New("Unable to connect to docker host")
		}

		cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
		pemCerts, _ := ioutil.ReadFile(caFile)

		tlsConfig.RootCAs = x509.NewCertPool()
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		tlsConfig.Certificates = []tls.Certificate{cert}
		tlsConfig.RootCAs.AppendCertsFromPEM(pemCerts)
	}
	docker, err := dockerclient.NewDockerClient(host, tlsConfig)
	if err != nil {
		logs.Error.Println("Unable to connect to docker host")
		return nil, host, err
	}
	version, err := docker.Version()
	if err != nil {
		logs.Error.Println("Unable to ping docker host")
		logs.Error.Println(err)
		return nil, host, err
	}
	logs.Trace.Println("Connected to Docker Host " + host)
	logs.Debug.Println("Docker Version: " + version.Version)
	logs.Debug.Println("Git Commit:" + version.GitCommit)
	logs.Debug.Println("Go Version:" + version.GoVersion)

	return docker, host, err
}
