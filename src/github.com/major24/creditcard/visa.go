package creditcard

type Visa struct {
	Rate float64
	Total float64
}

func (v Visa) TotalInterest() float64 {
    return v.Total * v.Rate
}
