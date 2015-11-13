package handler

import (
	"github.com/ory-am/ladon/policy"
	"github.com/pborman/uuid"
)

func superUserPolicy(subject string) policy.Policy {
	return &policy.DefaultPolicy{
		ID:          uuid.New(),
		Description: "Super user policy generated by the CLI.",
		Effect:      policy.AllowAccess,
		Subjects:    []string{subject},
		Permissions: []string{".*"},
		Resources:   []string{".*"},
	}
}