package messages

import "github.com/DCNT-developer/dcnt/electionsCore/primitives"

func GetSigner(msg interface{}) primitives.Identity {
	switch msg.(type) {
	case VolunteerMessage:
		v := msg.(*VolunteerMessage)
		return v.Signer
	case VoteMessage:
		v := msg.(*VoteMessage)
		return v.Signer
	case LeaderLevelMessage:
		v := msg.(*LeaderLevelMessage)
		return v.Signer
	}
	return primitives.NewIdentityFromInt(-1)
}
