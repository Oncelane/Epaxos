package epaxos

const (
	NULL = iota
	GET
	GETPREFIX
	PUT
	DELETE
	DELETEPREFIX
	CAS
	BATCH
)
