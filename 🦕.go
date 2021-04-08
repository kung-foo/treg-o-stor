/*
Copyright © 2000 Jonathan Camp <github@jonathan.camp>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the COPYING file for more details.
*/

package main

import (
	"math/rand"
	"runtime"

	"github.com/mackerelio/go-osstat/memory"
)

var (
	electronBuffer []byte
)

func init() {
	ram, _ := memory.Get()

	electronBuffer = make([]byte, int64(float64(ram.Available)*1.10))

	rand.Read(electronBuffer)
}

func main() {
	for i := 0; i < runtime.NumCPU()+1; i++ {
		go func() {
			for {
			}
		}()
	}

	select {}
}
