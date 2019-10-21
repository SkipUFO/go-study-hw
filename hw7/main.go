package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type env struct {
	name  string
	value string
}

var path string
var executable string

func getEnvs(path string) ([]env, error) {
	var envs []env

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return envs, err
	}

	for _, f := range files {
		if f.Mode().IsRegular() {
			file, err := os.Open(path + "/" + f.Name())
			if err != nil {
				return envs, err
			}
			defer file.Close()

			value, err := ioutil.ReadAll(file)
			if err != nil {
				return envs, err
			}
			envs = append(envs, env{name: f.Name(), value: string(value)})
		}
	}

	return envs, nil
}

func main() {

	args := os.Args

	if len(args) < 3 {
		fmt.Println("Not enough parameters. Usage <this app> /path/to/env/dir some_prog")
		os.Exit(1)
	}

	path = args[1]
	executable = args[2]

	// проверим, что path - директория, executable - файл
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !fi.Mode().IsDir() {
		fmt.Println("Invalid /path/to/env/dir parameter is not Dir. Usage <this app> /path/to/env/dir some_prog")
		os.Exit(1)
	}

	fi, err = os.Stat(executable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !fi.Mode().IsRegular() {
		fmt.Println("Invalid some_prog parameter is not File. Usage <this app> /path/to/env/dir some_prog")
		os.Exit(1)
	}

	newEnvs, err := getEnvs(path)
	if err != nil {
		fmt.Println(err)
	}

	var currentEnvs []env

	// Сохраняем старые значения, если значения не было, то ставим пустую строку, чтобы потом по этим переменным
	// делать unset
	for _, val := range newEnvs {
		if value, ok := os.LookupEnv(val.name); ok {
			currentEnvs = append(currentEnvs, env{name: val.name, value: value})
		} else {
			currentEnvs = append(currentEnvs, env{name: val.name, value: ""})
		}

		os.Setenv(val.name, val.value)
	}

	cmd := exec.Command(executable)
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	for _, val := range currentEnvs {
		if val.value == "" {
			os.Unsetenv(val.name)
		} else {
			os.Setenv(val.name, val.value)
		}
	}

	os.Exit(0)
}
