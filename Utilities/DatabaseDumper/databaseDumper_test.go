package main

import (
	"testing"

	"github.com/DCNT-developer/dcnt/testHelper"
)

func TestTest(t *testing.T) {
	dbo := testHelper.CreateAndPopulateTestDatabaseOverlay()

	err := ExportDatabaseJSON(dbo, true)
	if err != nil {
		t.Error(err)
	}
}
