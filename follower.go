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
		return AppendEntryResponse{Term: f.Term, Success: false}
	}

	storedTerm, present := f.Log[call.PreviousLogIndex]
	if !present {
		return AppendEntryResponse{Term: f.Term, Success: false}
	} else if storedTerm != call.PreviousLogTerm {
		return AppendEntryResponse{Term: f.Term, Success: false}
	}

	f.Term = call.Term

	return AppendEntryResponse{Term: f.Term, Success: true}
}
