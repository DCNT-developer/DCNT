package main

import (
	//"fmt"
	"os"

	"github.com/DCNT-developer/dcnt/common/interfaces"
	"github.com/DCNT-developer/dcnt/database/databaseOverlay"
	"github.com/DCNT-developer/dcnt/database/hybridDB"
	"github.com/DCNT-developer/dcnt/database/mapdb"
	"github.com/DCNT-developer/dcnt/util"
)

//DBInit

func InitBolt(cfg *util.dcntConfig) interfaces.DBOverlay {
	//fmt.Println("InitBolt")
	path := cfg.App.BoltDBPath + "/"

	os.MkdirAll(path, 0777)
	dbase := hybridDB.NewBoltMapHybridDB(nil, path+"FactomBolt-Import.db")
	return databaseOverlay.NewOverlay(dbase)
}

func InitLevelDB(cfg *util.dcntConfig) interfaces.DBOverlay {
	//fmt.Println("InitLevelDB")
	path := cfg.App.LdbPath + "/" + "FactoidLevel-Import.db"

	dbase, err := hybridDB.NewLevelMapHybridDB(path, false)

	if err != nil || dbase == nil {
		dbase, err = hybridDB.NewLevelMapHybridDB(path, true)
		if err != nil {
			panic(err)
		}
	}

	return databaseOverlay.NewOverlay(dbase)
}

func InitMapDB(cfg *util.dcntConfig) interfaces.DBOverlay {
	//fmt.Println("InitMapDB")
	dbase := new(mapdb.MapDB)
	dbase.Init(nil)
	return databaseOverlay.NewOverlay(dbase)
}
