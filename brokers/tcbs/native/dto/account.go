// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type BasicInfo struct {
	Code105C   string `json:"code105C"`
	Depository bool   `json:"depository"`
	Status     string `json:"status"`
	TcbsID     string `json:"tcbsId"`
	Type       string `json:"type"`
}

type PersonalInfo struct {
	Acronym             string       `json:"acronym"`
	AvatarURL           string       `json:"avatarUrl"`
	Birthday            string       `json:"birthday"`
	BusinessType        string       `json:"businessType"`
	ContactAddress      string       `json:"contactAddress"`
	CreatedDate         string       `json:"createdDate"`
	Email               string       `json:"email"`
	FirstName           string       `json:"firstName"`
	FlowOpenAccount     string       `json:"flowOpenAccount"`
	FullName            string       `json:"fullName"`
	FullNameNoAccent    string       `json:"fullNameNoAccent"`
	Gender              string       `json:"gender"`
	IdentityCard        IdentityCard `json:"identityCard"`
	LastName            string       `json:"lastName"`
	Nationality         string       `json:"nationality"`
	NationalityName     string       `json:"nationalityName"`
	PermanentAddress    string       `json:"permanentAddress"`
	PhoneNumber         string       `json:"phoneNumber"`
	PpBusinessField     string       `json:"ppBusinessField"`
	PpBusinessFieldName string       `json:"ppBusinessFieldName"`
	PpBusinessType      string       `json:"ppBusinessType"`
	PpBusinessTypeName  string       `json:"ppBusinessTypeName"`
	TaxIdNumber         string       `json:"taxIdNumber"`
	UpdatedDate         string       `json:"updatedDate"`
}

type IdentityCard struct {
	ExpireDate string `json:"expireDate"`
	IdDate     string `json:"idDate"`
	IdNumner   string `json:"idNumner"`
	IdPlace    string `json:"idPlace"`
	IdType     string `json:"idType"`
	Object     string `json:"object"`
}

type BankAccount struct {
	AccountName         string `json:"accountName"`
	AccountNameNoAccent string `json:"accountNameNoAccent"`
	AccountNo           string `json:"accountNo"`
	Authorized          string `json:"authorized"`
	BankAccountType     string `json:"bankAccountType"`
	BankCode            string `json:"bankCode"`
	BankName            string `json:"bankName"`
	BankSys             string `json:"bankSys"`
	BankType            string `json:"bankType"`
	BranchCode          string `json:"branchCode"`
}

type BankSubAccount struct {
	AccountName     string `json:"accountName"`
	AccountNo       string `json:"accountNo"`
	AccountType     string `json:"accountType"`
	AccountTypeName string `json:"accountTypeName"`
	IsDefault       string `json:"isDefault"`
	Status          string `json:"status"`
}

type SubAccountInformation struct {
	BankAccounts    []BankAccount    `json:"bankAccounts"`
	BankSubAccounts []BankSubAccount `json:"bankSubAccounts"`
	BasicInfo       BasicInfo        `json:"basicInfo"`
	PersonalInfo    PersonalInfo     `json:"personalInfo"`
}
