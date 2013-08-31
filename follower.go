package seaturtles

func NewFollower(id, term int) *Follower {
	return &Follower{Id: id, Term: term, Log: make(map[int]int)}
}

type Follower struct {
	Id   int
	Term int
	Log  map[int]int
}

func (f *Follower) AppendEntry(call AppendEntryCall) AppendEntryResponse {
	if call.Term < f.Term {
		return f.appendEntryResponse(false)
	}

	storedTerm, present := f.Log[call.PreviousLogIndex]
	if !present {
		return f.appendEntryResponse(false)
	} else if storedTerm != call.PreviousLogTerm {
		return f.appendEntryResponse(false)
	}

	f.Term = call.Term

	return f.appendEntryResponse(false)
}

func (f *Follower) appendEntryResponse(success bool) AppendEntryResponse {
	return AppendEntryResponse{Term: f.Term, Success: success}
}
