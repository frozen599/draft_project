package models

type Customer struct {
	Title     string `csv:"TITLE,omitempty" json:"title"`
	FirstName string `csv:"FIRST_NAME,omitempty" json:"first_name"`
	LastName  string `csv:"LAST_NAME,omitempty" json:"last_name"`
	Email     string `csv:"EMAIL,omitempty" json:"email"`
}

type EmailTemplate struct {
	From     string `json:"from,omitempty"`
	Subject  string `json:"subject,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
	Body     string `json:"body,omitempty"`
}
