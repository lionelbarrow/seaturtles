package seaturtles

type Follower struct {
	Id   int
	Term int
}

func (f *Follower) AppendEntry(call AppendEntryCall) AppendEntryResponse {
	if call.Term < f.Term {
		return AppendEntryResponse{Term: f.Term, Success: false}
	} else if call.Term > f.Term {
		f.Term = call.Term
	}

	return AppendEntryResponse{Term: f.Term, Success: true}
}
