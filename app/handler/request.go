package handler

import (
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/labstack/echo/v4"
)

type createTodoRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *createTodoRequest) bind(c echo.Context, todo *model.Todo) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	todo.Name = r.Name

	return nil
}

type addressRequest struct {
	AddressLine1          string `json:"address_line_1"`
	AddressLine2          string `json:"address_line_2"`
	Submunicipality       string `json:"submunicipality"`
	Municipality          string `json:"municipality"`
	SubAdministrativeArea string `json:"sub_administrative_area"`
	AdministrativeArea    string `json:"administrative_area"`
	Country               string `json:"country"`
	PostalCode            string `json:"postal_code"`
}

type onboardCallbackRequest struct {
	IDNumber     string `json:"id_number"`
	TitleEN      string `json:"title_en"`
	TitleTH      string `json:"title_th"`
	FirstNameEN  string `json:"first_name_en"`
	FirstNameTH  string `json:"first_name_th"`
	MiddleNameEN string `json:"middle_name_en"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameEN   string `json:"last_name_en"`
	LastNameTH   string `json:"last_name_th"`
	Birthdate    string `json:"birth_date"`
	Gender       string `json:"gender"`

	IDCardIssueDate   string `json:"id_card_issue_date"`
	IDCardExpiryDate  string `json:"id_card_expiry_date"`
	IDCardAddressFull string `json:"id_card_address_full"`

	IDCardAddress struct {
		AddressNo   string `json:"address_no"`
		Moo         string `json:"moo"`
		Trok        string `json:"trok"`
		Soi         string `json:"soi"`
		Road        string `json:"road"`
		Subdistrict string `json:"subdistrict"`
		District    string `json:"district"`
		Province    string `json:"province"`
	} `json:"id_card_address"`
	IDCardPhoto string `json:"id_card_photo"`

	FaceDetectionImage string `json:"face_detection_image"`
	MobileNumber       string `json:"mobile_number"`
	Email              string `json:"email"`

	IDAddress      addressRequest `json:"id_address"`
	ContactAddress addressRequest `json:"contact_address"`
	Occupation     string         `json:"occupation"`
}

func (r *onboardCallbackRequest) bind(c echo.Context, customer *model.Customer) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	// if err := c.Validate(r); err != nil {
	// 	return err
	// }

	customer.IDNumber = r.IDNumber
	customer.TitleEN = r.TitleEN
	customer.TitleTh = r.TitleTH
	customer.FirstNameEN = r.FirstNameEN
	customer.FirstNameTH = r.FirstNameTH
	customer.MiddleNameEN = r.MiddleNameEN
	customer.MiddleNameTH = r.MiddleNameTH
	customer.LastNameEN = r.LastNameEN
	customer.LastNameTH = r.LastNameTH
	customer.Birthdate = r.Birthdate
	customer.Gender = r.Gender

	customer.IDCard = model.IDCard{
		IssueDate:   r.IDCardIssueDate,
		ExiryDate:   r.IDCardExpiryDate,
		AddressFull: r.IDCardAddressFull,
		IDCardAddress: model.IDCardAddress{
			AddressNo:   r.IDCardAddress.AddressNo,
			Moo:         r.IDCardAddress.Moo,
			Trok:        r.IDCardAddress.Trok,
			Soi:         r.IDCardAddress.Soi,
			Road:        r.IDCardAddress.Road,
			Subdistrict: r.IDCardAddress.Subdistrict,
			District:    r.IDCardAddress.District,
			Province:    r.IDCardAddress.Province,
		},
		Photo: r.IDCardPhoto,
	}

	customer.FaceDetectionImage = r.FaceDetectionImage
	customer.MobileNumber = r.MobileNumber
	customer.Email = r.Email

	customer.IDAddress = model.IDAddress{
		BaseAddress: model.BaseAddress{
			AddressLine1:          r.IDAddress.AddressLine1,
			AddressLine2:          r.IDAddress.AddressLine2,
			Submunicipality:       r.IDAddress.Submunicipality,
			Municipality:          r.IDAddress.Municipality,
			SubAdministrativeArea: r.IDAddress.Submunicipality,
			AdministrativeArea:    r.IDAddress.AdministrativeArea,
			Country:               r.IDAddress.Country,
			PostalCode:            r.IDAddress.PostalCode,
		},
	}

	customer.ContactAddress = model.ContactAddress{
		BaseAddress: model.BaseAddress{
			AddressLine1:          r.ContactAddress.AddressLine1,
			AddressLine2:          r.ContactAddress.AddressLine2,
			Submunicipality:       r.ContactAddress.Submunicipality,
			Municipality:          r.ContactAddress.Municipality,
			SubAdministrativeArea: r.ContactAddress.Submunicipality,
			AdministrativeArea:    r.ContactAddress.AdministrativeArea,
			Country:               r.ContactAddress.Country,
			PostalCode:            r.ContactAddress.PostalCode,
		},
	}

	customer.Occupation = r.Occupation

	return nil
}
