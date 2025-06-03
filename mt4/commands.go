package mt4

import (
	"context"
	"strconv"
)

// BlockAccount blocks an MT4 account using the WBLOCKLOGINSUSER command.
func (c *Client) BlockAccount(ctx context.Context, opts BlockAccountOptions) (string, error) {
	params := map[string]string{
		"MASTER": opts.Master,
		"NAME":   opts.Name,
		"LOGIN":  strconv.Itoa(opts.Login),
	}
	return c.Execute(ctx, "WBLOCKLOGINSUSER", params)
}

// CreateAccount creates a new account using the WOPENTRADE command.
func (c *Client) CreateAccount(ctx context.Context, opts CreateAccountOptions) (string, error) {
	params := map[string]string{
		"MASTER":         opts.Master,
		"GROUP":          opts.Group,
		"NAME":           opts.Name,
		"PASSWORDTRADE":  opts.PassTrade,
		"PASSWORDINVEST": opts.PassInvest,
		"EMAIL":          opts.Email,
		"COUNTRY":        opts.Country,
		"STATE":          opts.State,
		"CITY":           opts.City,
		"ADDRESS":        opts.Address,
		"PHONE":          opts.Phone,
		"PHONE_PASSWORD": opts.PassPhone,
		"LEVERAGE":       strconv.Itoa(opts.Leverage),
		"AMOUNT":         strconv.FormatFloat(opts.Amount, 'f', 2, 64),
		"STARTLOGIN":     strconv.Itoa(opts.LoginStart),
		"COMMENT":        opts.Comment,
		"ENDLOGIN":       strconv.Itoa(opts.LoginEnd),
	}
	if opts.Demo {
		params["DEMO"] = "1"
	}
	if opts.SendReports {
		params["SEND_REPORTS"] = "1"
	} else {
		params["SEND_REPORTS"] = "0"
	}
	if opts.ReadOnly {
		params["READ_ONLY"] = "1"
	} else {
		params["READ_ONLY"] = "0"
	}
	return c.Execute(ctx, "WOPENTRADE", params)
}

// TradeBalance performs a trade balance operation.
func (c *Client) TradeBalance(ctx context.Context, opts TradeBalanceOptions) (string, error) {
	params := map[string]string{
		"MASTER":   "password1",
		"IP":       opts.IP,
		"GROUP":    opts.Group,
		"LOGIN":    strconv.Itoa(opts.Login),
		"CURRENCY": opts.Currency,
		"AMOUNT":   strconv.FormatFloat(opts.Amount, 'f', 2, 64),
		"COMMENT":  opts.Comment,
	}
	if opts.Credit != "" {
		params["CREDIT"] = opts.Credit
	}
	return c.Execute(ctx, opts.Type, params)
}
