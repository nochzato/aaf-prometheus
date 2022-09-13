package main

import (
	"fmt"
	"math/big"
)

func karatsuba(x *big.Int, y *big.Int) *big.Int {
	length := big.NewInt(int64(len(x.Text(10))))
	piece_length := new(big.Int)
	piece_length.Div(length, big.NewInt(2))
	exp := new(big.Int)
	exp.Exp(big.NewInt(10), piece_length, big.NewInt(0))
	a, b, c, d := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	a.Div(x, exp)
	b.Mod(x, exp)
	c.Div(y, exp)
	d.Mod(y, exp)
	ac, bd, abcd, adbc := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	ac.Mul(a, c)
	bd.Mul(b, d)
	abcd.Mul(big.NewInt(0).Add(a, b), big.NewInt(0).Add(c, d))
	adbc.Sub(big.NewInt(0).Sub(abcd, ac), bd)
	ans := new(big.Int)
	ans.Add(big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(0).Exp(big.NewInt(10), length, big.NewInt(0)), ac), big.NewInt(0).Mul(exp, adbc)), bd)
	return ans
}

func main() {
	var x, _ = new(big.Int).SetString("1685287499328328297814655639278583667919355849391453456921116729", 0)
	var y, _ = new(big.Int).SetString("7114192848577754587969744626558571536728983167954552999895348492", 0)
	fmt.Println(karatsuba(x, y))
}
