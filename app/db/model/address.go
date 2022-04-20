package model

type IDCardAddress struct {
	BaseModel
	AddressNo   string `json:"addressNo"`
	Moo         string `json:"moo"`
	Trok        string `json:"trok"`
	Soi         string `json:"soi"`
	Road        string `json:"road"`
	Subdistrict string `json:"subdistrict"`
	District    string `json:"district"`
	Province    string `json:"province"`
	IDCardID    uint
}

type BaseAddress struct {
	BaseModel
	AddressLine1          string `json:"addressLine1"`
	AddressLine2          string `json:"addressLine2"`
	Submunicipality       string `json:"submunicipality"`       // sublocality/subdistrict // แขวง/คำบล
	Municipality          string `json:"municipality"`          // city/town/district // เขต/อำเภอ
	SubAdministrativeArea string `json:"subAdministrativeArea"` // county
	AdministrativeArea    string `json:"administrativeArea"`    // state/province/region/territory // จังหวัด
	Country               string `json:"country"`
	PostalCode            string `json:"postalCode"` // postal code / zip code
}

type IDAddress struct {
	BaseAddress
	CustomerID uint
}

type ContactAddress struct {
	BaseAddress
	CustomerID uint
}
