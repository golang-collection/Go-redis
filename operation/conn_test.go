package operation

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-12 21:16
* @Description:
**/

func TestSetString(t *testing.T) {
	s, err := SetString("s1", "sssss")
	if err != nil{
		t.FailNow()
	}
	fmt.Println(s)
}

func TestGetString(t *testing.T) {
	s, err := GetString("s1")
	if err != nil{
		t.FailNow()
	}
	fmt.Println(s)
}

func TestAddElementToSet(t *testing.T) {
	s, err := AddElementToSet("myset", "a")
	if err != nil{
		t.FailNow()
	}
	fmt.Println(s)
}

func TestElementIsInSet(t *testing.T) {
	s, err := ElementIsInSet("myset", "b")
	if err != nil{
		t.FailNow()
	}
	fmt.Println(s)
}