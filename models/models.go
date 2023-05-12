package models

type LawBySubject struct {
	Subject        string `json:"Subject,omitempty"`
	SubjectID      string `json:"SubjectID,omitempty"`
	LawDescription string `json:"LawDescription,omitempty"`
}

type Typification struct {
	Legislation      []Legislation `json:"Legislation"`
	TypeOfOccurrence string        `json:"TypeOfOccurrence,omitempty"`
}

type Legislation struct {
	Law string
}
