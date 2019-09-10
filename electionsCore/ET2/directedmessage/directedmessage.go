package directedmessage

import "github.com/DCNT-developer/dcnt/electionsCore/imessage"

type DirectedMessage struct {
	LeaderIdx int
	Msg       imessage.IMessage
}
