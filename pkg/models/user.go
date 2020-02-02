package models

// User user
type User struct {
	ID int `json:"ID"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"UserName"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"FullName"`
	// City holds the value of the "city" field.
	City string `json:"City"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate CustomTime `json:"BirthDate"`
	// Department holds the value of the "department" field.
	Department string `json:"Department"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"Gender"`
	// ExperienceYears holds the value of the "experience_years" field.
	ExperienceYears int `json:"ExperienceYears"`
}
