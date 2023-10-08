package tst

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
	"testing"
)

func TestDoECDSA(t *testing.T) {
	// MzU5OTczMDA3NzMxNzE0OTA0MzI1ODcwMTE2NzA0MDg0ODI3ODQ2NTA1NTcyNTQ2MzE2NjYxMjk0NTA0NDA1MDUzNjY3ODU4MzUzODM=

	x, ok := new(big.Int).SetString("66271406351504038125715128559405452881170750034949756136589154064159035447322", 10)
	if !ok {
		t.Error("Set X Failure")
	}

	y, ok := new(big.Int).SetString("114288593703500244685566124739300491454999152685226803507402288876530978169306", 10)
	if !ok {
		t.Error("Set Y Failure")
	}

	fmt.Printf("(X,Y) :: (%s,%s)\n", x.String(), y.String())

	pub := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	r, ok := new(big.Int).SetString("29657361150645529277577426606756283524550134900958298248265400694010018017020", 10)
	if !ok {
		t.Error("Set R Failure")
	}

	s, ok := new(big.Int).SetString("95412299490148537794322908080753161676206659394390803133140016554321124816370", 10)
	if !ok {
		t.Error("Set S Failure")
	}

	fmt.Printf("(R,S) :: (%s,%s)\n", r.String(), s.String())

	str := string([]byte{})
	h := sha256.New()
	io.WriteString(h, str)
	res := h.Sum(nil)

	fmt.Printf("%#v\n", ecdsa.Verify(pub, res, r, s))

	hmac.New(sha256.New,[]byte(""))
}
