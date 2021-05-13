package cmd

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v3"
	"fmt"
	"strings"
)

type Config struct {
	Alias map[string]string
}

func SaveConfig(config Config) (string, error) {
	file := os.Getenv("HOME") + "/.godeprc"
	bytes, err := yaml.Marshal(&config)

	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(file, bytes, 0644)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func LoadConfig() Config {
	config := Config{}
	config.Alias = make(map[string]string)
	bytes, err := OpenConfig()
	if err != nil {
		fmt.Println("read file fail", err)
		return config
	}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		fmt.Println("parse file fail", err)
		return config
	}
	return config
}

func OpenConfig() ([]byte, error) {
	file := os.Getenv("HOME") + "/.godeprc"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("read file fail", err)
		return []byte{}, err
	}

	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return []byte{}, err
	}
	return fd, nil
}

func Aliased(name string) string {
	version := ""
	rt := strings.Split(name, "@")
	if len(rt) == 2 {
		name, version = rt[0], rt[1]
	}

	config := LoadConfig()
	val, b := config.Alias[name]

	if b {
		name = val
 	}

	if version != "" && !strings.Contains(name, "@") {
		return name + "@" + version
	} else {
		return name
	}
}
