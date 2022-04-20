package model

type Customer struct {
	BaseModel
	IDNumber     string `json:"idNumber"`
	TitleEN      string `json:"titleEn"`
	TitleTh      string `json:"titleTh"`
	FirstNameEN  string `json:"firstNameEn"`
	FirstNameTH  string `json:"firstNameTh"`
	MiddleNameEN string `json:"middleNameEn"`
	MiddleNameTH string `json:"middleNameTh"`
	LastNameEN   string `json:"lastNameEn"`
	LastNameTH   string `json:"lastNameTh"`
	Birthdate    string `json:"birthdate"`
	Gender       string `json:"gender"`
	IDCard       IDCard `json:"idCard"`

	FaceDetectionImage string         `json:"faceDetectionImage"`
	MobileNumber       string         `json:"mobileNumber"`
	Email              string         `json:"email"`
	IDAddress          IDAddress      `json:"idAddress"`
	ContactAddress     ContactAddress `json:"contactAddress"`
	Occupation         string         `json:"occupation"`
}
