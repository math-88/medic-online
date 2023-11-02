package entity

import "github.com/google/uuid"

type Protocol struct {
	ID          uuid.UUID
	Name        string
	Description string
	Actions     []Action
}

type Action struct {
	ID          uuid.UUID
	Drug        string
	Dosage      string
	Time        string
	Duration    string
	Description string
}

// Clone returns a clone of this protocol
func (protocol *Protocol) Clone() *Protocol {
	return &Protocol{
		ID:          protocol.ID,
		Name:        protocol.Name,
		Description: protocol.Description,
	}
}

// Clone returns a clone of this action
func (action *Action) Clone() *Action {
	return &Action{
		ID:          action.ID,
		Drug:        action.Drug,
		Dosage:      action.Dosage,
		Time:        action.Time,
		Duration:    action.Duration,
		Description: action.Description,
	}
}
