package command


var cmds map[string]Command


func Init() {

}

func GetCommand(cmd *string, args []string) Command {

	switch *cmd {
	case "run":
		command := new(RunCommand)
		command.Args = args
		return command
	case "chile":
		command := new(ChildCommand)
		command.InitCommand = args[0]
		command.Args = args[1:]
		return command
	case "help":
		// todo
	default:
		panic("not support command")
	}
}
