package command

import (
	"fmt"
	"os/exec"
	"os"
	"syscall"
)

type ChildCommand struct {
	InitCommand string
	Args []string
}

func (this *ChildCommand) Process() {

	fmt.Println("ChildCommand...")
	fmt.Println(this.InitCommand)

	// todo cgroup


	cmd := exec.Command(this.InitCommand, this.Args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/home/mydocker/centosfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("thing", 0))


}

func (this *ChildCommand) Type() CommandType {
	return CHILD
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}