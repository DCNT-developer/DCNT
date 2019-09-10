package testHelper

import (
	"github.com/DCNT-developer/dcnt/common/constants"
	"github.com/DCNT-developer/dcnt/common/interfaces"
	"github.com/DCNT-developer/dcnt/common/primitives"
)

func NewRepeatingHash(b byte) interfaces.IHash {
	tmp := make([]byte, constants.HASH_LENGTH)
	for i := range tmp {
		tmp[i] = b
	}
	return primitives.NewHash(tmp)
}
