package epaxos

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Ballot struct {
	epoch  uint64
	uid    uint64
	peerid uint64
}

func (b *Ballot) String() string {
	return fmt.Sprintf("%d.%d.%d", b.epoch, b.uid, b.peerid)
}

func NewBallot(s string) (*Ballot, bool) {
	numstrs := strings.Split(s, ".")
	if len(numstrs) != 3 {
		return nil, false
	}

	epoch, err := strconv.Atoi(numstrs[0])
	if err != nil {
		return nil, false
	}
	id, err := strconv.Atoi(numstrs[1])
	if err != nil {
		return nil, false
	}
	peer, err := strconv.Atoi(numstrs[2])
	if err != nil {
		return nil, false
	}

	ballot := &Ballot{
		epoch:  uint64(epoch),
		uid:    uint64(id),
		peerid: uint64(peer),
	}
	return ballot, true
}

func (b *Ballot) Increase() {
	b.uid++
}

func piGroup(ballots ...Ballot) []Ballot {
	sort.Sort(sort.Reverse(SortByBallot(ballots)))
	maxBallot := ballots[0]
	return []Ballot{maxBallot}
}

func (b *Ballot) same(other *Ballot) bool{
	return b.epoch==other.epoch && b.peerid == other.peerid && b.uid == other.uid
}

type SortByBallot []Ballot

func (s SortByBallot) Len() int      { return len(s) }
func (s SortByBallot) Swap(x, y int) { s[x], s[y] = s[y], s[x] }
func (s SortByBallot) Less(i, j int) bool {
	if s[i].epoch != s[j].epoch {
		return s[i].epoch < s[j].epoch
	}
	if s[i].uid != s[j].uid {
		return s[i].uid < s[j].uid
	}
	return s[i].epoch < s[j].epoch
}
