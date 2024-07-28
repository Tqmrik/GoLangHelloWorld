package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (DM directMessage) importance() int {
	switch {
	case DM.isUrgent:
		return 50
	default:
		return DM.priorityLevel
	}
}

func (GM groupMessage) importance() int {
	return GM.priorityLevel
}

func (SA systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch n.(type) {
	case directMessage:
		var dm = n.(directMessage)
		return dm.senderUsername, dm.importance()
	case groupMessage:
		var gm = n.(groupMessage)
		return gm.groupName, gm.importance()
	case systemAlert:
		var sa = n.(systemAlert)
		return sa.alertCode, sa.importance()
	default:
		return "", 0
	}
}
