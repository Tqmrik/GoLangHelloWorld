package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	user := User{}
	user.Name = name
	user.Membership.Type = membershipType
	switch membershipType {
	case "premium":
		user.Membership.MessageCharLimit = 1000
	default:
		user.Membership.MessageCharLimit = 100
	}
	return user
}
