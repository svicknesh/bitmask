package bitmask

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

const (
	SrvCapRSA Bit = 1 << iota // fastest way to initialize capabilities
	SrvCapEC
	SrvCapED
	SrvCapReserved1
	SrvCapReserved2
	SrvCapCrvPub
)

func TestBitmaskNew(T *testing.T) {

	b := New(3)
	fmt.Println("new\t", b)

	b.Set(SrvCapRSA, SrvCapEC, SrvCapED, SrvCapCrvPub)
	//b.Set(SrvCapRSA, SrvCapED)
	fmt.Println("set\t", b)

	fmt.Println("has Curve\t", b.Has(SrvCapCrvPub))

	b.Remove(SrvCapCrvPub)
	//b.Set(SrvCapRSA, SrvCapED)
	fmt.Println("remove\t", b)

	fmt.Println("has Curve\t", b.Has(SrvCapCrvPub))

	b.SetAll()
	fmt.Println("setall\t", b)

	b.Clear()
	fmt.Println("clear\t", b)

}

func TestBitmaskString(T *testing.T) {
	b, err := NewFromStr("0100111")
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("new\t", b)

	fmt.Println("has RSA\t", b.Has(SrvCapRSA))

}

func TestBitmaskJSON(T *testing.T) {

	type Example struct {
		//B *Bitmask
		//B string `json:"B"`
		B *Bitmask `json:"b"`
	}

	e := new(Example)
	e.B = new(Bitmask)

	e.B.SetLength(12)
	e.B.Set(SrvCapRSA, SrvCapEC, SrvCapED, SrvCapCrvPub)
	//fmt.Println(e.B)
	//fmt.Println(e.b.Has(SrvCapCrvPub))
	//e.B.Remove(SrvCapCrvPub)
	//fmt.Println(e.b.Has(SrvCapCrvPub))

	eB, err := json.Marshal(e)
	if nil != err {
		log.Println("marshal: ", err)
		os.Exit(1)
	}
	fmt.Println(string(eB))

	ex := new(Example)
	ex.B = new(Bitmask)
	err = json.Unmarshal(eB, ex)
	if nil != err {
		log.Println("unmarshal: ", err)
		os.Exit(1)
	}

	fmt.Println("unmarshalled", ex.B)
	fmt.Println(ex.B.Has(SrvCapCrvPub))

}
