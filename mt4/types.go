package mt4

// BlockAccountOptions represents parameters for the WBLOCKLOGINSUSER command.
type BlockAccountOptions struct {
	Login  int
	Master string
	Name   string
}

// CreateAccountOptions represents parameters for the WOPENTRADE command.
type CreateAccountOptions struct {
	Address     string
	Amount      float64
	City        string
	Comment     string
	Country     string
	Demo        bool
	Email       string
	Group       string
	Leverage    int
	LoginEnd    int
	LoginStart  int
	Master      string
	Name        string
	PassInvest  string
	PassPhone   string
	PassTrade   string
	Phone       string
	ReadOnly    bool
	SendReports bool
	State       string
}

// TradeBalanceOptions represents parameters for trade balance commands.
type TradeBalanceOptions struct {
	Amount   float64
	Comment  string
	Credit   string
	Currency string
	Group    string
	IP       string
	Login    int
	Type     string
}
