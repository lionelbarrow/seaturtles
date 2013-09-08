package seaturtles

type AppendEntryCall struct {
	Term          int
	PreviousEntry LogEntry
	Entries       []LogEntry
}

type AppendEntryResponse struct {
	Term    int
	Success bool
}
