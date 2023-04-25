package main

import (
	"flag"
	"fmt"
)

var base float64
var rate float64
var years int

func calProfitTotal(base, rate float64, years int) float64 {
	var profit float64
	for i := 0; i < years; i++ {
		// profit = profit+math.Pow(base,float64(i))
		profit = (profit + base) * (1 + rate)
	}

	return profit
}

func main() {
	flag.Float64Var(&base, "base", 12000, "复利起算资金")
	flag.Float64Var(&rate, "rate", 0.03, "复利利率")
	flag.IntVar(&years, "years", 30, "复利年数")
	flag.Parse()
	tol := calProfitTotal(base, rate, years)
	fmt.Println("全部钱:", tol)
	baseTol := float64(years) * base
	fmt.Println("全部投入的钱:", baseTol)
	profit := tol - baseTol
	fmt.Println("利润:", profit)
	tax := tol * 0.03
	fmt.Println("取出扣税:", tax)
	fmt.Println("真实利润:", profit-tax)
	fmt.Println("扣税后全部钱:", tol-tax)
}
