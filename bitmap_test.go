package bitmap

import (
	"testing"
)

func TestBitMap_Get(t *testing.T) {
	bm := NewBitMap(8)
	err := bm.Set(1)
	if err != nil {
		t.Fatal(err)
	}
	res, err := bm.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if !res {
		t.Fatal("except 1, actual 0")
	}
}
func TestBitMap_SetN(t *testing.T) {
	bm := NewBitMap(8)
	flag, err := bm.SetN(1)
	if err != nil {
		t.Fatal(err)
	}
	if flag == false {
		t.Fatal("except 0, actual 1")
	}
	flag, err = bm.SetN(1)
	if err != nil {
		t.Fatal(err)
	}
	if flag == true {
		t.Fatal("except 1, actual 0")
	}
	res, err := bm.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if !res {
		t.Fatal("expect 1, actual 0")
	}
}
func TestBitMap_ResetSetN(t *testing.T) {
	bm := NewBitMap(8)
	flag, err := bm.SetN(1)
	if err != nil {
		t.Fatal(err)
	}
	if flag == false {
		t.Fatal("except not setting before, actual already sed")
	}
	flag, err = bm.ResetN(1)
	if err != nil {
		t.Fatal(err)
	}
	if flag == false {
		t.Fatal("except already setting before, actual not sed")
	}

	res, err := bm.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if res {
		t.Fatal("expect 0，actual 1")
	}
}
func TestBitMap_Overflow(t *testing.T) {
	bm := NewBitMap(8)
	err := bm.Set(9)
	if err == nil {
		t.Fatal("expect report error，actual err is empty")
	}
}
