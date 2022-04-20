package model

type IDCard struct {
	BaseModel
	IssueDate     string        `json:"issueDate"`
	ExiryDate     string        `json:"exiryDate"`
	AddressFull   string        `json:"addressFull"`
	IDCardAddress IDCardAddress `json:"address"`
	Photo         string        `json:"photo"`
	CustomerID    uint
}
