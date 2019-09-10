package main

import (
	"os"
	"testing"

	"github.com/DCNT-developer/dcnt/common/interfaces"
	"github.com/DCNT-developer/dcnt/database/databaseOverlay"
	"github.com/DCNT-developer/dcnt/database/leveldb"
	"github.com/DCNT-developer/dcnt/testHelper"
)

func TestCheckDatabaseFromDBO(t *testing.T) {
	dbo := testHelper.CreateAndPopulateTestDatabaseOverlay()
	CheckDatabase(dbo)
}

func TestCheckDatabaseFromState(t *testing.T) {
	state := testHelper.CreateAndPopulateTestStateAndStartValidator()
	CheckDatabase(state.DB.(interfaces.DBOverlay))
}

func TestCheckDatabaseFromWSAPI(t *testing.T) {
	ctx := testHelper.CreateWebContext()
	state := ctx.Server.Env["state"].(interfaces.IState)
	dbase := state.GetDB().(interfaces.DBOverlay)

	CheckDatabase(dbase)
}

var dbFilename string = "levelTest.db"

func TestCheckDatabaseForLevelDB(t *testing.T) {
	m, err := leveldb.NewLevelDB(dbFilename, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer CleanupLevelDB(t, m)

	dbo := databaseOverlay.NewOverlay(m)
	testHelper.PopulateTestDatabaseOverlay(dbo)

	CheckDatabase(dbo)

}

func CleanupLevelDB(t *testing.T, b interfaces.IDatabase) {
	err := b.Close()
	if err != nil {
		t.Errorf("%v", err)
	}
	err = os.RemoveAll(dbFilename)
	if err != nil {
		t.Errorf("%v", err)
	}
}
