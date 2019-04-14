package bitmap

import (
	"errors"
)

type BitMap struct {
	bits []byte
	cap int64
	oneCount int64
}
func NewBitMap(cap int64)*BitMap{
	bm:=&BitMap{}
	bm.bits=make([]byte,cap/8+1)
	bm.cap=cap
	return bm
}
func (bm *BitMap)Set(x int64)error{
	div:=x/8
	mod:=byte(x%8)
	flag:=byte(1<<mod)
	if x>bm.cap{
		return errors.New("value is more than bitmap cap")
	}
	bm.bits[div]|=flag
	return nil
}
func (bm *BitMap)SetN(x int64)(bool,error){
	div:=x/8
	mod:=byte(x%8)
	flag:=byte(1<<mod)
	if x>bm.cap{
		return false,errors.New("value is more than bitmap cap")
	}
	v:=bm.bits[div]
	v&=flag
	if v>0{
		return false,nil
	}
	bm.bits[div]|=flag
	bm.oneCount++
	return true,nil
}
func (bm *BitMap)ResetSetN(x int64)(bool,error){
	div:=x/8
	mod:=byte(x%8)
	flag:=byte(1<<mod)
	if x>bm.cap{
		return false,errors.New("value is more than bitmap cap")
	}
	v:=bm.bits[div]
	v&=flag
	if v==0{
		return false,nil
	}
	bm.bits[div]&=^flag
	bm.oneCount--
	return true,nil
}
func (bm *BitMap)Get(x int64)(bool,error){
	div:=x/8
	mod:=byte(x%8)
	flag:=byte(1<<mod)
	if x>bm.cap{
		return false,errors.New("value is more than bitmap cap")
	}
	v:=bm.bits[div]
	v&=flag
	if v>0{
		return true,nil
	}else{
		return false,nil
	}

}
func (bm *BitMap)Len()int64{
	return bm.oneCount
}
