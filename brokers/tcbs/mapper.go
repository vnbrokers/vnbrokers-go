package tcbs

import "github.com/vnbrokers/vnbrokers-go/domain"

func MapAccountInformation(custodyCode string, response AccountInformationResponse) domain.Account {
	accountID := custodyCode
	if response.BasicInfo != nil && response.BasicInfo.Code105C != "" {
		accountID = response.BasicInfo.Code105C
	}
	displayName := ""
	if response.PersonalInfo != nil {
		displayName = response.PersonalInfo.FullName
	}
	return domain.Account{
		Broker:      "tcbs",
		AccountID:   accountID,
		DisplayName: displayName,
		Raw:         rawPayload(response, nil),
	}
}
