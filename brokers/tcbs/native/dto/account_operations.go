package dto

type GetSubAccountInformationRequest struct {
	CustodyCode string
	Fields      string
}

type GetSubAccountInformationResponse SubAccountInformation
