package mysql

import (
	"fmt"
	model "meliQuasar/internal/api/quasar/model"
)

func setConnection() (ConnectionInfoMS)  {
	return ConnectionInfoMS{ User: "root", Pass: "123456789", Host: "localhost", DbName: "satellitesDB" }
}

func GetSatellites()([]model.Satellite){
	db := MySqlDb{}
	config := setConnection()
	error := db.Connect(&config)

	if(error != nil){
		fmt.Println(error)
	}
	data, error := db.Query("SELECT * FROM satellitesDB.satellite")

	if(error != nil){
		fmt.Println(error)
	}

	var lstSat []model.Satellite

	for data.Next(){
		var idsatellite int
		var name string
		var positionx, positiony float32

		error=data.Scan(&idsatellite,&name,&positionx,&positiony)

		if(error!=nil){
			panic(error.Error())
		}

		sat := model.Satellite {}
		sat.Satelite(idsatellite,name)
		sat.SetPosition(positionx, positiony)

		lstSat = append(lstSat, sat)
	}

	return lstSat
}