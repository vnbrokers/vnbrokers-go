package tcbs

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type AccountService struct {
	broker *Broker
}

func (s *AccountService) GetSubAccountInfo(
	ctx context.Context,
	custodyCode string,
	fields string,
) (AccountInformationResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsList); err != nil {
		return AccountInformationResponse{}, err
	}
	query := url.Values{}
	query.Set("fields", fields)
	path := "/eros/v2/get-profile/by-username/" + url.PathEscape(custodyCode)
	response, err := s.broker.send(ctx, "account.get_sub_account_info", true, transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url(path) + "?" + query.Encode(),
		Headers: s.broker.headers(true, false),
	})
	if err != nil {
		return AccountInformationResponse{}, err
	}
	var accountResponse AccountInformationResponse
	if err := decode(response, &accountResponse); err != nil {
		fmt.Println(string(response.Raw))
		return AccountInformationResponse{}, sdkerrors.Decode("tcbs", "account.get_sub_account_info", "decode account response", response.Body, err)
	}
	return accountResponse, nil
}

func (s *AccountService) GetBasicInfo(ctx context.Context, custodyCode string) (domain.Account, error) {
	response, err := s.GetSubAccountInfo(ctx, custodyCode, "basicInfo,personalInfo")
	if err != nil {
		return domain.Account{}, err
	}
	return MapAccountInformation(custodyCode, response), nil
}
