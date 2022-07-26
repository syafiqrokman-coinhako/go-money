package money

import (
	"math/big"
)

type calculator struct{}

func (c *calculator) add(a, b *Amount) *Amount {
	//return &Amount{a.val + b.val}
	tmp := big.Int{}
	tmp.Add(&a.val, &b.val)
	return &Amount{tmp}
}

func (c *calculator) subtract(a, b *Amount) *Amount {
	//return &Amount{a.val - b.val}
	tmp := big.Int{}
	tmp.Sub(&a.val, &b.val)
	return &Amount{tmp}
}

func (c *calculator) multiply(a *Amount, m int64) *Amount {
	//return &Amount{a.val * m}
	tmp := big.Int{}
	tmp.Mul(&a.val, big.NewInt(m))
	return &Amount{tmp}
}

func (c *calculator) divide(a *Amount, d int64) *Amount {
	//return &Amount{a.val / d}
	tmp := big.Int{}
	tmp.Div(&a.val, big.NewInt(d))
	return &Amount{tmp}
}

func (c *calculator) modulus(a *Amount, d int64) *Amount {
	//return &Amount{a.val % d}
	tmp := big.Int{}
	tmp.Mod(&a.val, big.NewInt(d))
	return &Amount{tmp}

}

//func (c *calculator) allocate(a *Amount, r, s int) *Amount {
//	return &Amount{a.val * int64(r) / int64(s)}
//}

func (c *calculator) absolute(a *Amount) *Amount {
	//if a.val < 0 {
	//	return &Amount{-a.val}
	//}
	//
	//return &Amount{a.val}
	a.val.Abs(&a.val)
	return a
}

func (c *calculator) negative(a *Amount) *Amount {
	//if a.val > 0 {
	//	return &Amount{-a.val}
	//}
	//
	//return &Amount{a.val}
	a.val.Neg(&a.val)
	return a
}

//
//func (c *calculator) round(a *Amount, e int) *Amount {
//	if a.val == 0 {
//		return &Amount{0}
//	}
//
//	absam := c.absolute(a)
//	exp := int64(math.Pow(10, float64(e)))
//	m := absam.val % exp
//
//	if m > (exp / 2) {
//		absam.val += exp
//	}
//
//	absam.val = (absam.val / exp) * exp
//
//	if a.val < 0 {
//		a.val = -absam.val
//	} else {
//		a.val = absam.val
//	}
//
//	return &Amount{a.val}
//}
