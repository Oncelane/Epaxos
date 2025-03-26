package epaxos

type core struct {
	commitch chan command
}

func NewCore() *core {
	return &core{}
}
func (c *core) Propose(cmd *command) {

}

func (c *core) handlerPrepare() {

}

func (c *core) handler() {

}

func (c *core) CommitChannel() chan command {
	return c.commitch
}
