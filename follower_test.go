package seaturtles

import "testing"

func TestAppendEntriesRejectsEntryLowTermCalls(t *testing.T) {
	follower := NewFollower(1, 5)
	follower.Log = map[int]int{1: 1}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 3, PreviousLogIndex: 1, PreviousLogTerm: 1}

	response := follower.AppendEntry(appendEntryCall)

	if response.Success {
		t.Error("Follower accepted AppendEntry call with lower term than its own.")
	}
}

func TestAppendEntriesSendsCurrentTermWhenRejecting(t *testing.T) {
	follower := NewFollower(1, 5)
	follower.Log = map[int]int{1: 1}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 3, PreviousLogIndex: 1, PreviousLogTerm: 1}

	response := follower.AppendEntry(appendEntryCall)

	if response.Term != 5 {
		t.Error("Follower did not respond with its own term when rejecting a call")
	}
}

func TestAppendEntriesSavesHigherTermCalls(t *testing.T) {
	follower := NewFollower(1, 1)
	follower.Log = map[int]int{1: 1}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 1, PreviousLogTerm: 1}

	follower.AppendEntry(appendEntryCall)

	if follower.Term != 2 {
		t.Error("Follower did not update its own term after receiving high term call.")
	}
}

func TestAppendEntriesReturnsGreatestKnownTerm(t *testing.T) {
	follower := NewFollower(1, 1)
	follower.Log = map[int]int{1: 1, 2: 2}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 1, PreviousLogTerm: 1}

	response := follower.AppendEntry(appendEntryCall)

	if response.Term != 2 {
		t.Error("Follower response did not update to include new term")
	}
}

func TestAppendEntriesRejectsIfPreviousLogIndexIsTooLow(t *testing.T) {
	follower := NewFollower(1, 2)
	follower.Log = map[int]int{1: 1, 2: 2}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 2, PreviousLogTerm: 1}

	response := follower.AppendEntry(appendEntryCall)

	if response.Success {
		t.Error("Follower accepted AppendEntry call where PreviousLogIndex's Term was below recorded Term.")
	}
}

func TestAppendEntriesRejectsIfLogDoesNotContainPreviousLogIndexAndTerm(t *testing.T) {
	follower := NewFollower(1, 1)
	follower.Log = map[int]int{1: 1, 2: 1}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2, PreviousLogIndex: 3, PreviousLogTerm: 1}

	response := follower.AppendEntry(appendEntryCall)

	if response.Success {
		t.Error("Follower accepted AppendEntry call with PreviousLogIndex below its LogIndex")
	}
}
