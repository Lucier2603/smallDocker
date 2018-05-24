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


// exec 以交互方式进入某个容器
type ExecCommand struct {
	id int
}


func (this *ExecCommand) Process() {

}


// ps 查看当时正在运行的容器
type PsCommand struct {
}

func (this *PsCommand) Process() {

}
