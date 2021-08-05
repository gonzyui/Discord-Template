package commands

type CmdInput struct {
	command string
	args []string
}

func (I *CmdInput) GetArgs() []string {
	return I.args
}