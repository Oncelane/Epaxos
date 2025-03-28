package epaxos

type Command struct {
	key   string
	value []byte
	op    int
}

func (c *Command) Conflic(other *Command) bool {
	return c.key == other.key
}
