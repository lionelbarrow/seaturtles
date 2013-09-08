package seaturtles

func NewFollower(term int) *Follower {
	return &Follower{Term: term, Log: make(map[int]int)}
}

type Follower struct {
	Term    int
	Log     map[int]int
	Entries []LogEntry
}

func (f *Follower) AppendEntry(call AppendEntryCall) AppendEntryResponse {
	if call.Term < f.Term {
		return f.appendEntryResponse(false)
	}

	storedTerm, present := f.Log[call.PreviousEntry.Index]
	if !present {
		return f.appendEntryResponse(false)
	} else if storedTerm != call.PreviousEntry.Term {
		return f.appendEntryResponse(false)
	}

	f.Term = call.Term
	f.Entries = append(f.Entries, call.Entries...)

	return f.appendEntryResponse(true)
}

func (f *Follower) appendEntryResponse(success bool) AppendEntryResponse {
	return AppendEntryResponse{Term: f.Term, Success: success}
}
