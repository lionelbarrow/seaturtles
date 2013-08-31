package seaturtles

type AppendEntryCall struct {
	LeaderId         int
	Term             int
	PreviousLogIndex int
	PreviousLogTerm  int
}

type AppendEntryResponse struct {
	Term    int
	Success bool
}
