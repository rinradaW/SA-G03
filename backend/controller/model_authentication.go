package controller

// LoginPayload login body
type LoginPayload struct {
	StudentId    string `json:"student_id"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
	Stdid string `json:"stdid"`
}