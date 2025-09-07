package main

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Membership
	Name string
}

func newUser(name string, membershipType string) User {
	var messageCharLimit int
	switch membershipType {
	case "premium":
		messageCharLimit = 1000
	default:
		messageCharLimit = 100
	}

	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: messageCharLimit,
		},
	}
}
