package generator

type CommandType int

type Dockerfile struct {
	CommandType CommandType
	Argument    interface{}
}

type DockerFunc uint8

const (
	From DockerFunc = iota
	Comment
	Run
	Maintainer
	Entrypoint
	Expose
	Cmd
	WorkDir
)

var CommandsMap = map[CommandType]string{
	CommandType(From):       "FROM",
	CommandType(Comment):    "#",
	CommandType(Run):        "RUN",
	CommandType(Cmd):        "CMD",
	CommandType(Maintainer): "MAINTAINER",
	CommandType(Expose):     "EXPOSE",
	CommandType(Entrypoint): "ENTRYPOINT",
	CommandType(WorkDir):    "WORKDIR",
}
