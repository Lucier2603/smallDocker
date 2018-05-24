package command

import (
	"os/exec"
	"os"
	//"syscall"
	"fmt"
)

// run 运行容器
type RunCommand struct {
	Name string
	InitCmd string
	Args []string
}

func (this *RunCommand) Process() {

	fmt.Println("RunCommand...")
	// copy self with 'chile' param
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, this.Args...)...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	 //This should build with GOOS=linux
	//	 //see https://github.com/lizrice/containers-from-scratch/issues/1
	//	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
	//		syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	//}
	
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func (this *RunCommand) Type() CommandType {
	return EXEC
}
