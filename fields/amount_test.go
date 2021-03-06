package fields

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"
)

func Test_bbb1(t *testing.T) {

	amt, _ := NewAmountFromFinString("2:248")

	fmt.Println(amt.GetValue().String())

}

func Test_amt_cut(t *testing.T) {

	bigbts := bytes.Repeat([]byte{255}, 9)

	bigamt := new(big.Int).SetBytes(bigbts)

	fmt.Println(bigamt.String())
	fmt.Println("3927027935805858720555")

}

func Test_t1(t *testing.T) {

	amt, _ := NewAmountFromFinString("HAC123456789:240")

	a1, _, _ := amt.CompressForMainNumLen(4, true)
	a2, _, _ := amt.CompressForMainNumLen(6, false)

	fmt.Println(a1.ToFinString(), a2.ToFinString())

}

func Test_t2(t *testing.T) {

	bignum, _ := new(big.Int).SetString("387465837000000", 10)

	amount, _ := NewAmountByBigInt(bignum)

	fmt.Println(len(amount.Numeral))
	fmt.Println(amount.Serialize())
	fmt.Println(amount.ToFinString())

	new_amount, ischg, err := amount.EllipsisDecimalFor11SizeStore()
	fmt.Println(new_amount, ischg, err)

	if new_amount != nil {
		fmt.Println(len(new_amount.Numeral))
		fmt.Println(new_amount.Serialize())
		fmt.Println(new_amount.ToFinString())
	}

}
