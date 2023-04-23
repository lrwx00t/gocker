package generator

import (
	"fmt"
	"strings"
)

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

func WORKDIR(arg string) Dockerfile {
	return Dockerfile{CommandType(WorkDir), arg}
}

func RUN(args []string) Dockerfile {
	var args_literal string
	for _, val := range args {
		args_literal += "'" + val + "', "
	}
	args_literal = "[" + strings.TrimSuffix(args_literal, ", ") + "]"
	return Dockerfile{CommandType(Run), args}
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
