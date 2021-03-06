package pkg

import "math"

func fn() {
	var x float64

	_ = math.Pow(x, 0) //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(x, 1) //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(x, 2) //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(x, 3) //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(x, 6)

	const a = 2
	const b = 2.0
	const c float64 = 2

	_ = math.Pow(a, 2)     //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(b, 2)     //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(c, 2)     //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(a*1.0, 2) //@ diag(`could expand call to math.Pow`)

	_ = math.Pow(x*2, 2) //@ diag(`could expand call to math.Pow`)
	_ = math.Pow(x+2, 2) //@ diag(`could expand call to math.Pow`)

	_ = math.Pow(x, x)
	_ = math.Pow(x, -1)
}
