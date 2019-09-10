// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package electionMsgs_test

import (
	"testing"

	"github.com/DCNT-developer/dcnt/common/messages"
	. "github.com/DCNT-developer/dcnt/common/messages/electionMsgs"
	"github.com/DCNT-developer/dcnt/common/messages/msgsupport"
	"github.com/DCNT-developer/dcnt/common/primitives"
)

func TestUnmarshalfolunteerAudit_test(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic caught during the test - %v", r)
		}
	}()

	messages.General = new(msgsupport.GeneralFactory)
	primitives.General = messages.General

	a := new(FedVoteMsg)
	err := a.UnmarshalBinary(nil)
	if err == nil {
		t.Error("Error is nil when it shouldn't be")
	}

	err = a.UnmarshalBinary([]byte{})
	if err == nil {
		t.Error("Error is nil when it shouldn't be")
	}
}
