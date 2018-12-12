package btcjson


type Omni_GetbalanceCmd struct {
	Address *string
	Propertyid *int
}

func NewOmni_GetbalanceCmd(account *string,propertyid *int) *Omni_GetbalanceCmd{
	return &Omni_GetbalanceCmd{
		Address: account,
		Propertyid: propertyid,
	}
}


type Omni_ListtransactionsCmd struct {
	Txid *string
	Count *int
	Skip *int
}

func NewOmni_ListtransactionsCmd(account *string,count *int,skip *int) *Omni_ListtransactionsCmd{
	return &Omni_ListtransactionsCmd{
		Txid: account,
		Count: count,
		Skip: skip,
	}
}

type Omni_SendCmd struct{
	Fromaddress *string
	Toaddress *string
	Propertyid *int
	Amount *string
	Redeemaddress *string
	Referenceamount *string
}

func NewOmni_SendCmd(account *string,toaccount *string,propertyid *int,amount *string) *Omni_SendCmd{
	return &Omni_SendCmd{
		Fromaddress: account,
		Toaddress: toaccount,
		Propertyid: propertyid,
		Amount: amount,
	}
}
