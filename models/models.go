package models

type Customer struct {
	Title     string `csv:"TITLE,omitempty"`
	FirstName string `csv:"FIRST_NAME,omitempty"`
	LastName  string `csv:"LAST_NAME,omitempty"`
	Email     string `csv:"EMAIL,omitempty"`
}

type EmailTemplate struct {
	From     string `json:"from,omitempty"`
	Subject  string `json:"subject,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
	Body     string `json:"body,omitempty"`
}
