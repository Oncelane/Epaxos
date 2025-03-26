package epaxos

type command struct {
	key   string
	value []byte
	op    int
}

func (c *command) Conflic(other *command) bool {
	return c.key == other.key
}
