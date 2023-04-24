package main

import (
	"fmt"
	satSrv "meliQuasar/internal/api/quasar/services"
)

func main(){
	fmt.Println("Test")

	a := satSrv.GetSatellites()

	fmt.Println(a)
}