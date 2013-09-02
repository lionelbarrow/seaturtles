package seaturtles

import (
	e "github.com/lionelbarrow/examples"
	"testing"
)

func createFollower(id, term int) *Follower {
	follower := NewFollower(id, term)
	follower.Log = map[int]int{1: 1}
	return follower
}

func TestAppendEntriesWithBadClient(t *testing.T) {
	e.Describe("rejected requests", t,
		e.It("rejects requests with an old term", func(ex *e.Example) {
			follower := createFollower(1, 5)
			appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 3, PreviousLogIndex: 1, PreviousLogTerm: 1}

			response := follower.AppendEntry(appendEntryCall)

			ex.Expect(response.Success).ToBeFalse()
			ex.Expect(response.Term).ToEqual(5)
		}),

		e.It("rejects requests with a low previous log index", func(ex *e.Example) {
			follower := createFollower(1, 2)
			follower.Log = map[int]int{1: 1, 2: 2}
			appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 2, PreviousLogTerm: 1}

			response := follower.AppendEntry(appendEntryCall)

			ex.Expect(response.Success).ToBeFalse()
			ex.Expect(response.Term).ToEqual(2)
		}),

		e.It("rejects requests with a previous log index and non-matching previous log term", func(ex *e.Example) {
			follower := createFollower(1, 1)
			follower.Log = map[int]int{1: 1, 2: 1}
			appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 3, PreviousLogTerm: 1}

			response := follower.AppendEntry(appendEntryCall)

			ex.Expect(response.Success).ToBeFalse()
			ex.Expect(response.Term).ToEqual(1)
		}),
	)
}

func TestAppendEntriesReturnsGreatestKnownTerm(t *testing.T) {
	follower := createFollower(1, 1)
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 1, PreviousLogTerm: 1}

	response := follower.AppendEntry(appendEntryCall)

	if response.Term != 2 {
		t.Error("Follower response did not update to include new term")
	}
}
