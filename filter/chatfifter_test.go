package filter

import (
	"testing"
)

func Test_fi( t *testing.T){
	f:=GetInstance()
	f.Insert("我草")
	f.Insert("我去")
	f.Insert("脏词")
	for i:=0;i<1000000;i++{
		f.Deal("脏词我一不小心就说了一句脏词，我是不小心说出这句脏词的",42)
	}
}

