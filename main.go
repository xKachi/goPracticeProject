package main

import (
	//"fmt"

	"fmt"

	"example.com/calculator-project/cmdmanager"
	//"example.com/calculator-project/filemanager"
	"example.com/calculator-project/prices"
)



func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm ,taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

	
}