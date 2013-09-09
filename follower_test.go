package seaturtles

import (
	e "github.com/lionelbarrow/examples"
	"testing"
)

func createFollower(term int) *Follower {
	follower := NewFollower(term)
	follower.Log = map[int]int{1: 1}
	follower.Entries = []LogEntry{LogEntry{Term: 1, Index: 1, Item: "a"}}
	return follower
}

func TestAppendEntriesHappyPath(t *testing.T) {
	e.When("nothing is wrong with the request", t,
		e.It("appends the new entry to the log", func(expect e.Expectation) {
			follower := createFollower(1)
			call := AppendEntryCall{Term: 2,
				PreviousEntry: LogEntry{Term: 1, Index: 1},
				Entries: []LogEntry{
					LogEntry{Index: 2, Term: 2, Item: "b"},
				},
			}

			response := follower.AppendEntry(call)

			expect(response.Success).ToBeTrue()
			expect(response.Term).ToEqual(2)

			newTermEntry := follower.Entries[1]

			expect(newTermEntry.Term).ToEqual(2)
			expect(newTermEntry.Item).ToEqual("b")
		}),
	)
}

func TestAppendEntriesWithBadClient(t *testing.T) {
	e.Describe("rejected requests", t,
		e.It("rejects requests with an old term", func(expect e.Expectation) {
			follower := createFollower(5)
			call := AppendEntryCall{Term: 3, PreviousEntry: LogEntry{Term: 1, Index: 1}}

			response := follower.AppendEntry(call)

			expect(response.Success).ToBeFalse()
			expect(response.Term).ToEqual(5)
		}),

		e.It("rejects requests with a low previous log index", func(expect e.Expectation) {
			follower := createFollower(2)
			follower.Log = map[int]int{1: 1, 2: 2}
			call := AppendEntryCall{Term: 2, PreviousEntry: LogEntry{Term: 1, Index: 2}}

			response := follower.AppendEntry(call)

			expect(response.Success).ToBeFalse()
			expect(response.Term).ToEqual(2)
		}),

		e.It("rejects requests with a previous log index and non-matching previous log term", func(expect e.Expectation) {
			follower := createFollower(1)
			follower.Log = map[int]int{1: 1, 2: 1}
			call := AppendEntryCall{Term: 2, PreviousEntry: LogEntry{Term: 1, Index: 3}}

			response := follower.AppendEntry(call)

			expect(response.Success).ToBeFalse()
			expect(response.Term).ToEqual(1)
		}),
	)
}

func TestAppendEntriesReturnsGreatestKnownTerm(t *testing.T) {
	follower := createFollower(1)
	call := AppendEntryCall{Term: 2, PreviousEntry: LogEntry{Term: 1, Index: 1}}

	response := follower.AppendEntry(call)

	if response.Term != 2 {
		t.Error("Follower response did not update to include new term")
	}
}
