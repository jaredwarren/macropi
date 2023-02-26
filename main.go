package main

import (
	"fmt"
	"os"
	"time"
)

var (
	Nil = byte(0x00)
	// ADownCmd Key      = []byte("\x00\x04\x00\x00\x00\x00\x00")
	ADownCmd Key      = []byte{Nil, byte(0x04), Nil, Nil, Nil, Nil, Nil}
	Shift    Modifier = true
)

type Key []byte
type Modifier bool

func PressKey(key Key, opts ...Modifier) {
	prefix := byte(0x00)
	for _, k := range opts {
		if k == Shift {
			prefix = byte(0x20)
		}
	}
	k := []byte{prefix}
	k = append(k, key...)
	fmt.Println(k)
}

func main() {

	PressKey(ADownCmd, Shift)
	if true {
		return
	}

	a := uint8(0x04)
	nilC := uint8(0x00)
	s := uint8(0x20)
	one := uint8(0x1e)

	// shift (right?)
	fmt.Println([]byte(" \x00\x04\x00\x00\x00\x00\x00"))
	fmt.Println([]byte("\x20\x00\x04\x00\x00\x00\x00\x00")) // 32 dec = 20 hex -> {space-char} -> 32d

	// {space}
	fmt.Println([]byte("\x00\x00,\x00\x00\x00\x00\x00"))
	fmt.Println([]byte("\x00\x00\x2c\x00\x00\x00\x00\x00")) // 2c hex = 44 dec

	fmt.Println([]byte("\x00\x00\x04\x00\x00\x00\x00\x00"))
	fmt.Println([]byte{nilC, nilC, a, nilC, nilC, nilC, nilC, nilC})

	fmt.Println("--")
	fmt.Println([]byte("\x20\x00\x04\x00\x00\x00\x00\x00")) // 32 dec = 20 hex -> {space-char} -> 32d
	fmt.Println([]byte{s, nilC, a, nilC, nilC, nilC, nilC, nilC})
	fmt.Println([]byte{nilC, nilC, one, nilC, nilC, nilC, nilC, nilC})

	if true {
		return
	}

	hidfp, err := os.OpenFile("/dev/hidg0", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer hidfp.Close()

	fmt.Println("starting...")
	time.Sleep(2 * time.Second)
	fmt.Println("go!")

	// a
	{
		n, err := hidfp.Write([]byte("\x00\x00\x04\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	// release
	{
		n, err := hidfp.Write([]byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}

	// A
	{
		n, err := hidfp.Write([]byte(" \x00\x04\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	// {space}
	{
		n, err := hidfp.Write([]byte("\x00\x00,\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	// b
	{
		n, err := hidfp.Write([]byte("\x00\x00\x06\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	// release
	{
		n, err := hidfp.Write([]byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
}
