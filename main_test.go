package bitmap

import (
	"testing"
)

func TestBitMap_Get(t *testing.T) {
	bm:=NewBitMap(8)
	err:=bm.Set(1)
	if err!=nil{
		t.Fatal(err)
	}
	res,err:=bm.Get(1)
	if err!=nil{
		t.Fatal(err)
	}
	if !res{
		t.Fatal("该位没有置1")
	}
}
func TestBitMap_SetN(t *testing.T) {
	bm:=NewBitMap(8)
	flag,err:=bm.SetN(1)
	if err!=nil{
		t.Fatal(err)
	}
	if flag==false{
		t.Fatal("该位应该还没有设置，但是返回已经设置为1")
	}
	flag,err=bm.SetN(1)
	if err!=nil{
		t.Fatal(err)
	}
	if flag==true{
		t.Fatal("该位应该已经设置了，但是返回没有设置")
	}
	res,err:=bm.Get(1)
	if err!=nil{
		t.Fatal(err)
	}
	if !res{
		t.Fatal("该位应该已经设置，但是返回没有置1")
	}
}
func TestBitMap_ResetSetN(t *testing.T) {
	bm:=NewBitMap(8)
	flag,err:=bm.SetN(1)
	if err!=nil{
		t.Fatal(err)
	}
	if flag==false{
		t.Fatal("该位应该还没有设置，但是返回已经设置为1")
	}
	flag,err=bm.ResetSetN(1)
	if err!=nil{
		t.Fatal(err)
	}
	if flag==false{
		t.Fatal("该位应该已经设置为0，但是返回没有设置")
	}

	flag,err=bm.ResetSetN(1)
	if err!=nil{
		t.Fatal(err)
	}
	if flag==true{
		t.Fatal("该位应该已经设置了0，但是返回没有设置")
	}

	res,err:=bm.Get(1)
	if err!=nil{
		t.Fatal(err)
	}
	if res{
		t.Fatal("该位应该已经设置0，但是返回没有置")
	}
}
func TestBitMap_Overflow(t *testing.T) {
	bm:=NewBitMap(8)
	err:=bm.Set(9)
	if err==nil{
		t.Fatal("应该越界，但是返回正常")
	}

}

