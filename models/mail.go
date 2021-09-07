package models

type EditMail struct {
	Subject  *string `json:"subject"`
	Body     *string `json:"body"`
	Type     *string `json:"type"`
	Schedule *string `json:"schedule"`
	Status   *bool   `json:"status"`
}

type Mail struct {
	ID        string `json:"id"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Type      string `json:"type"`
	Schedule  string `json:"schedule"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MailFilter struct {
	Limit  *int    `json:"limit"`
	Offset *int    `json:"offset"`
	Status *bool   `json:"status"`
	SiteID *string `json:"site_id"`
	UserID *string `json:"user_id"`
}

type Mails struct {
	Hits  []*Mail `json:"hits"`
	Count int     `json:"count"`
}

type NewMail struct {
	Subject  string  `json:"subject"`
	Body     string  `json:"body"`
	Type     string  `json:"type"`
	Schedule string  `json:"schedule"`
	SiteID   *string `json:"site_id"`
}
