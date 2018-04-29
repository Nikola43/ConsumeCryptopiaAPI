package main

type Data struct {
	ID                   int         `json:"Id"`
	Name                 string      `json:"Name"`
	Symbol               string      `json:"Symbol"`
	Algorithm            string      `json:"Algorithm"`
	WithdrawFee          float64     `json:"WithdrawFee"`
	MinWithdraw          float64     `json:"MinWithdraw"`
	MaxWithdraw          float64     `json:"MaxWithdraw"`
	MinBaseTrade         float64     `json:"MinBaseTrade"`
	IsTipEnabled         bool        `json:"IsTipEnabled"`
	MinTip               float64     `json:"MinTip"`
	DepositConfirmations int         `json:"DepositConfirmations"`
	Status               string      `json:"Status"`
	StatusMessage        interface{} `json:"StatusMessage"`
	ListingStatus        string      `json:"ListingStatus"`
}
