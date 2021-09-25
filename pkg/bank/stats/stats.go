package stats

import "github.com/NekruziKabirzoda/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)

	for i:=0; i<len(payments); i++{
    sum+=payments[i].Amount
	//sum= sum/types.Money(i)
	}
	sum = sum/types.Money(len(payments))
return sum
}