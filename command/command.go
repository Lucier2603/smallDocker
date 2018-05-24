package command


const (
	EXEC CommandType = 1
	PS CommandType = 2
	CHILD CommandType = 3
)

type CommandType int32


type Command interface {
	// type
	Type() CommandType
	// process
	Process()
}


