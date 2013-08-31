package seaturtles

type Follower struct {
  Id int
  Term int
}

func (f Follower) AppendEntry(call AppendEntryCall) AppendEntryResponse {
  if call.Term < f.Term {
    return AppendEntryResponse{Term: f.Term, Success: false}
  }

  return AppendEntryResponse{Term: 1, Success: true}
}
