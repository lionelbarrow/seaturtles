package seaturtles

import "testing"

func TestAppendRejectsEntryLowTermCalls(t *testing.T) {
  follower := Follower{Id: 1, Term: 5}
  appendEntryCall := AppendEntryCall{LeaderId: 2, Term: 3}

  response := follower.AppendEntry(appendEntryCall)

  if response.Success {
    t.Error("Follower accepted AppendEntry call with lower term than its own.")
  }
}
