package main

import (
	"fmt"

	"github.com/hiroki-kondo-git/calc_tax_pra/tax"
)

func main() {
	systems := []tax.TaxSystem{&tax.IndianTax{TaxPercentage: 30, Income: 1000}, &tax.SingaporeTax{TaxPercentage: 10, Income: 2000}}
	calcurateTotalTax(systems)
}

func calcurateTotalTax(taxSystems []tax.TaxSystem) {
	total := 0
	for _, system := range taxSystems {
		tax := system.CalcurateTax()
		total += tax
	}
	fmt.Println(total)
}
