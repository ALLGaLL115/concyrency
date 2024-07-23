package main

import (
	// "fmt"
	"fmt"
	"math"
	"strconv"

	// "strings"
	"unicode"
)

type testStruct struct {
	On          bool
	Ammo, Power int
}

func (st *testStruct) Shoot() bool {
	if !st.On || st.Ammo == 0 {
		return false
	}
	st.Ammo = st.Ammo - 1
	return true
}

func (st *testStruct) RideBike() bool {
	if !st.On || st.Power == 0 {
		return false
	}

	st.Power = st.Power - 1
	return true
}

func main() {

	fn := func(x uint) uint {
		m := int(x)
		n := 10
		res := 0
		//u, _ := strconv.ParseUint("42", 10, 64)
		for m/n >= 1 {
			n *= 10
		}

		for n >= 1 {
			if m/n%10%2 == 0 && m/n%10 != 0 {
				res = res*10 + m/n%10
			}

			m %= n
			n /= 10
		}
		if res == 0 {
			return uint(100)
		}

		return uint(res)

	}

	fmt.Println(fn(727178))
}
func adding(a string, b string) {
	ra := []rune(a)
	rb := []rune(b)

	sa := ""
	sb := ""

	for _, val := range ra {
		if unicode.IsDigit(val) {
			sa = sa + string(val)
		}
	}

	for _, val := range rb {
		if unicode.IsDigit(val) {
			sb = sb + string(val)
		}
	}

	x, errx := strconv.Atoi(sa)
	y, erry := strconv.Atoi(sb)

	if errx != nil {
		panic(errx)
	}
	if erry != nil {
		panic(erry)
	}
	fmt.Println(math.Round(1.2323452*100) / 100)
	fmt.Println(x + y)
	// return uint64(x + y)

}
