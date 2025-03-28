package epaxos

type core struct {
	commitch chan *Command
}

func NewCore() *core {
	return &core{}
}
func (c *core) Propose(cmd *Command) {

}

func (c *core) handlerPrepare() {

}

func (c *core) handler() {

}

func (c *core) CommitChannel() chan *Command {
	return c.commitch
}
