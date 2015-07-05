package main

import "io"
import "os"
import "bufio"
import "github.com/samalba/dockerclient"
import "crypto/tls"
import "crypto/x509"
import "path/filepath"
import "errors"
import "github.com/codegangsta/cli"
import "io/ioutil"
import "unicode/utf8"
import "strings"
import "unicode"
import "gopkg.in/redis.v3"

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
		Debug.Println(scanner.Text())
		text += StringTransform(scanner.Text() + "\n")
	}
	err := scanner.Err()
	if err != nil {
		Error.Println("There was an error with the scanner", err)
	}
	return text, err
}

func getRedisClient(c *cli.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.GlobalString("redis"),
		Password: c.GlobalString("redis-password"),
		DB:       int64(c.GlobalInt("redis-db")),
	})

	_, err := client.Ping().Result()
	if err != nil {
		Error.Println("Unable to connect to redis host")
		return nil, err
	}
	Trace.Printf("Connected to Redis Host %s/%d", c.GlobalString("redis"), c.GlobalInt("redis-db"))

	return client, nil
}

func getDockerCient(c *cli.Context) (*dockerclient.DockerClient, string, error) {
	host := c.GlobalString("host")
	if host == "" {
		Error.Println("Incorrect usage, please set the docker host")
		return nil, "", errors.New("Unable to connect to docker host")
	}

	tlsConfig := &tls.Config{}

	certPath := c.GlobalString("cert")
	if certPath != "" {
		caFile := filepath.Join(certPath, "ca.pem")
		if _, err := os.Stat(caFile); os.IsNotExist(err) {
			Error.Println("Cannot open file : " + caFile)
			Error.Println("Incorrect usage, please set correct cert files")
			return nil, host, errors.New("Unable to connect to docker host")
		}

		certFile := filepath.Join(certPath, "cert.pem")
		if _, err := os.Stat(certFile); os.IsNotExist(err) {
			Error.Println("Cannot open file : " + certFile)
			Error.Println("Incorrect usage, please set correct cert files")
			return nil, host, errors.New("Unable to connect to docker host")
		}

		keyFile := filepath.Join(certPath, "key.pem")
		if _, err := os.Stat(keyFile); os.IsNotExist(err) {
			Error.Println("Cannot open file : " + keyFile)
			Error.Println("Incorrect usage, please set correct cert files")
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
		Error.Println("Unable to connect to docker host")
		return nil, host, err
	}
	version, err := docker.Version()
	if err != nil {
		Error.Println("Unable to ping docker host")
		Error.Println(err)
		return nil, host, err
	}
	Trace.Println("Connected to Docker Host " + host)
	Debug.Println("Docker Version: " + version.Version)
	Debug.Println("Git Commit:" + version.GitCommit)
	Debug.Println("Go Version:" + version.GoVersion)

	return docker, host, err
}
