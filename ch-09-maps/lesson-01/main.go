package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	users := make(map[string]user)
	for i := range names {
		new := user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
		users[names[i]] = new
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}
