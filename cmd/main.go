package main

import (
	"fmt"
	satSrv "meliQuasar/internal/services"
)

func main(){
	fmt.Println("Test")

	a := satSrv.GetSatellites()

	fmt.Println(a)
}