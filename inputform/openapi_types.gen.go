// Package payment provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package payment

// AddressInformation defines model for AddressInformation.
type AddressInformation struct {
	City           string  `json:"City"`
	State          string  `json:"State"`
	StreetAddress  string  `json:"StreetAddress"`
	StreetAddress2 *string `json:"StreetAddress2,omitempty"`
	ZipCode        int64   `json:"ZipCode"`
}

// ContactInformation defines model for ContactInformation.
type ContactInformation struct {
	Cell        *int64 `json:"Cell,omitempty"`
	Email       string `json:"Email"`
	HomePhone   *int64 `json:"HomePhone,omitempty"`
	PhoneNumber int64  `json:"PhoneNumber"`
	WorkPhone   *int64 `json:"WorkPhone,omitempty"`
}

// HealthcheckObject defines model for HealthcheckObject.
type HealthcheckObject struct {
	ErrorMessage string `json:"ErrorMessage"`
	Status       string `json:"Status"`
}

// MessageFromBene defines model for MessageFromBene.
type MessageFromBene struct {
	Message     string `json:"Message"`
	PhoneNumber *int64 `json:"PhoneNumber,omitempty"`
}

// NewFormEntry defines model for NewFormEntry.
type NewFormEntry struct {
	AddressInformation *AddressInformation `json:"AddressInformation,omitempty"`
	ContactInformation *ContactInformation `json:"ContactInformation,omitempty"`
	MessageFromBene    *MessageFromBene    `json:"MessageFromBene,omitempty"`
	UserInformation    *UserInformation    `json:"UserInformation,omitempty"`
}

// UserInformation defines model for UserInformation.
type UserInformation struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

// NewFormEntryJSONBody defines parameters for NewFormEntry.
type NewFormEntryJSONBody interface{}

// NewFormEntryJSONRequestBody defines body for NewFormEntry for application/json ContentType.
type NewFormEntryJSONRequestBody NewFormEntryJSONBody

