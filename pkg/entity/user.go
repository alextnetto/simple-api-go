package entity

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender gender `json:"gender,omitempty"`
}

type gender string

const (
	male   gender = "male"
	female        = "female"
)
