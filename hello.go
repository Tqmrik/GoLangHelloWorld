package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	switch {
	case (mToSend.sender.name != "" && mToSend.sender.number != 0) && (mToSend.recipient.name != "" && mToSend.recipient.number != 0):
		return true
	default:
		return false
	}
}
