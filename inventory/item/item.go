package item

type Item struct {
	Name     string
	Price    float64
	Quantity int
	Tax      float64
	Calc     func(float64) float64
}

func checkneg(cur float64) bool {
	if cur < 0 {
		return true
	}

	return false
}

func RawCalc(cur float64) float64 {
	if checkneg(cur) {
		return 0
	}
	total := cur + cur*0.125
	return total
}

func ImportedCalc(cur float64) float64 {
	if checkneg(cur) {
		return 0
	}
	var total float64
	total = cur + cur*0.1
	if total <= 100 {
		total = total + 5
	} else if total <= 200 {
		total = total + 10
	} else {
		total = total + total*0.05
	}
	return total
}

func ManufacturedCalc(cur float64) float64 {
	if checkneg(cur) {
		return 0
	}
	var total float64
	total = cur + cur*0.125
	total = total + total*0.02
	return total
}
