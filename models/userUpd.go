package models

type UserUpd struct {
	NameOrg    string `json:"name_org" form:"name_org"`
	DOBOrg     string `json:"dob_org" form:"dob_org"`
	ContactOrg string `json:"contact_org" form:"contact_org"`
	Name       string `json:"name" form:"name"`
	DOB        string `json:"dob" form:"dob"`
	Contact    string `json:"contact" form:"contact"`
}
