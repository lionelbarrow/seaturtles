package seaturtles

import "testing"

func TestAppendEntriesRejectsEntryLowTermCalls(t *testing.T) {
	follower := &Follower{Id: 1, Term: 5}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 3}

	response := follower.AppendEntry(appendEntryCall)

	if response.Success {
		t.Error("Follower accepted AppendEntry call with lower term than its own.")
	}
}

func TestAppendEntriesSendsCurrentTermWhenRejecting(t *testing.T) {
	follower := &Follower{Id: 1, Term: 5}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 3}

	response := follower.AppendEntry(appendEntryCall)

	if response.Term != 5 {
		t.Error("Follower did not respond with its own term when rejecting a call")
	}
}

func TestAppendEntriesSavesHigherTermCalls(t *testing.T) {
	follower := &Follower{Id: 1, Term: 1}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2}

	follower.AppendEntry(appendEntryCall)

	if follower.Term != 2 {
		t.Error("Follower did not update its own term after receiving high term call.")
	}
}

func TestAppendEntriesReturnsGreatestKnownTerm(t *testing.T) {
	follower := &Follower{Id: 1, Term: 1}
	appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 2}

	response := follower.AppendEntry(appendEntryCall)

	if response.Term != 2 {
		t.Error("Follower response did not update to include new term")
	}
}
