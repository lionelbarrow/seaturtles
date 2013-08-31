package seaturtles

type AppendEntryCall struct {
  LeaderId int
  Term int
}

type AppendEntryResponse struct {
  Term int
  Success bool
}
