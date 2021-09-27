package tax

type TaxSystem interface {
	CalcurateTax() int
}

type IndianTax struct {
	TaxPercentage int
	Income        int
}

func (i *IndianTax) CalcurateTax() int {
	tax := i.Income * i.TaxPercentage / 100
	return tax
}

type SingaporeTax struct {
	TaxPercentage int
	Income        int
}

func (i *SingaporeTax) CalcurateTax() int {
	tax := i.Income * i.TaxPercentage
	return tax
}
