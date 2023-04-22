package main

// TODO:
// Formatter for single and double quotes
// Multiline
// runner/tester
// remaining docker keywords

import (
	"fmt"
	"strings"
)

type DockerFunc uint8

const (
	From DockerFunc = iota
	Comment
	Run
	Maintainer
	Entrypoint
	Expose
	Cmd
)

type CommandType int

type Dockerfile struct {
	CommandType CommandType
	Argument    interface{}
}

func FROM(arg string) Dockerfile {
	return Dockerfile{CommandType(From), arg}
}

func ENTRYPOINT(args []string) Dockerfile {
	// args_literal := "[" + strings.Join(args, ", ") + "]"
	var args_literal string
	for _, val := range args {
		args_literal += "'" + val + "', "
	}
	args_literal = "[" + strings.TrimSuffix(args_literal, ", ") + "]"
	return Dockerfile{CommandType(Entrypoint), args_literal}
}

func COMMENT(arg string) Dockerfile {
	return Dockerfile{CommandType(Comment), arg}
}

func RUN(args []string) Dockerfile {
	var args_literal string
	for _, val := range args {
		args_literal += "'" + val + "', "
	}
	args_literal = "[" + strings.TrimSuffix(args_literal, ", ") + "]"
	return Dockerfile{CommandType(Run), args_literal}
}

func CMD(arg string) Dockerfile {
	return Dockerfile{CommandType(Cmd), arg}
}

func MAINTAINER(arg string) Dockerfile {
	return Dockerfile{CommandType(Maintainer), arg}
}

func EXPOSE(arg string) Dockerfile {
	return Dockerfile{CommandType(Expose), arg}
}

var CommandsMap = map[CommandType]string{
	CommandType(From):       "FROM",
	CommandType(Comment):    "#",
	CommandType(Run):        "RUN",
	CommandType(Cmd):        "CMD",
	CommandType(Maintainer): "MAINTAINER",
	CommandType(Expose):     "EXPOSE",
	CommandType(Entrypoint): "ENTRYPOINT",
}

func GenerateDockerfileFromMap(funcInputMap map[CommandType]interface{}) {
	for key, value := range funcInputMap {
		generateDockerfile(key, value)
	}
}

func generateDockerfile(commandType CommandType, argument interface{}) {
	commandKey := CommandsMap[commandType]
	if s, ok := argument.(string); ok {
		fmt.Println(commandKey, s)
	} else if slice, ok := argument.([]string); ok {
		for _, s := range slice {
			fmt.Println(commandKey, s)
		}
	} else {
		panic("Invalid argument type")
	}
}

func GenerateDockerfile(s []Dockerfile) {
	for _, value := range s {
		commandMap := map[CommandType]interface{}{value.CommandType: value.Argument}
		GenerateDockerfileFromMap(commandMap)
	}
}

func main() {
	test := `hello
	there 
	commnd`
	GenerateDockerfile([]Dockerfile{
		FROM("ubuntu"),
		CMD("This is a sample comment"),
		COMMENT("This is a sample comment"),
		RUN([]string{"echo 'hello world'"}),
		MAINTAINER("John Doe"),
		ENTRYPOINT([]string{"ansible", "python"}),
		EXPOSE("8080"),
		EXPOSE("8080"),
		RUN([]string{"echo 'hello world'2"}),
		RUN([]string{test}),
	})
}
