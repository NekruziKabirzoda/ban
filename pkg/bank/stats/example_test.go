package stats

import (
	"github.com/NekruziKabirzoda/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
		{ID: 1, Amount:10},
		{ID: 2, Amount:400},
		{ID: 3, Amount:10},
	}


    fmt.Println(Avg (payments))
	//Output: 140
}
