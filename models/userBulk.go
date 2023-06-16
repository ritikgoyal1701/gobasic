package models

type UserBulk struct {
	Name    []string `json:"name" form:"name"`
	DOB     []string `json:"dob" form:"dob"`
	Contact []string `json:"contact" form:"contact"`
}
