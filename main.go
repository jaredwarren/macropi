package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	hidfp, err := os.OpenFile("/dev/hidg0", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer hidfp.Close()

	fmt.Println("starting...")
	time.Sleep(2 * time.Second)
	fmt.Println("go!")

	// gk := uint8('g')
	// // NULL_CHAR*2+chr(4)+NULL_CHAR*5)
	// n, err := hidfp.Write([]byte{gk})
	// if err != nil {
	// 	fmt.Println("err:", err.Error())
	// }
	// fmt.Println("Write:", n)

	// NULL_CHAR*2+chr(4)+NULL_CHAR*5)
	{
		n, err := hidfp.Write([]byte("\x00\x00\x04\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	{
		n, err := hidfp.Write([]byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}

	{
		n, err := hidfp.Write([]byte(" \x00\x04\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	{
		n, err := hidfp.Write([]byte("\x00\x00,\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	{
		n, err := hidfp.Write([]byte("\x00\x00\x06\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	{
		n, err := hidfp.Write([]byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("Write:", n)
	}
	// {space}\x00\x00,\x00\x00\x00\x00\x00
	// \x00\x00\x06\x00\x00\x00\x00\x00
	//
	if false {
		// x := nilChar*2 + rune(4) + nilChar*5
		for i := 0; i < 200; i++ {
			n, err := hidfp.Write([]byte{uint8(i)})
			if err != nil {
				fmt.Println("err:", err.Error())
			}
			fmt.Println("Write:", n)
			// hidfp.Write([]byte{uint8(i)})
		}
	}
}
