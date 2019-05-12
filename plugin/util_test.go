// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"encoding/base64"
	"testing"
)

func Test_convertRegistry(t *testing.T) {
	from := &globalRegistry{
		Address:  "docker.io",
		Username: "spaceghost",
		Password: "dianarossfan",
	}
	to := convertRegistry(from)
	if got, want := to.Address, from.Address; got != want {
		t.Errorf("Got address %s, want %s", got, want)
	}
	if got, want := to.Username, from.Username; got != want {
		t.Errorf("Got username %s, want %s", got, want)
	}
	if got, want := to.Password, from.Password; got != want {
		t.Errorf("Got password %s, want %s", got, want)
	}
}

func Test_parseRegistry(t *testing.T) {
	tests := []struct {
		registry string
		account  string
		region   string
	}{
		{"123456789.dkr.ecr.us-east-1.amazonaws.com", "123456789", "us-east-1"},
		{"123456789.foo.bar.us-east-1.amazonaws.com", "", ""},
	}
	for _, test := range tests {
		account, region := parseRegistry(test.registry)
		if got, want := account, test.account; got != want {
			t.Errorf("Want account %q, got %q", want, got)
		}
		if got, want := region, test.region; got != want {
			t.Errorf("Want region %q, got %q", want, got)
		}
	}
}

func Test_parseToken(t *testing.T) {
	tests := []struct {
		token    string
		username string
		password string
		error    error
	}{
		{"YXdzOnBhc3N3b3Jk", "aws", "password", nil},
		{"12345", "", "", base64.CorruptInputError(4)},
	}
	for _, test := range tests {
		username, password, err := parseToken(test.token)
		if got, want := err, test.error; got != want {
			t.Errorf("Want error %v, got %v", want, got)
		}
		if got, want := username, test.username; got != want {
			t.Errorf("Want account %q, got %q", want, got)
		}
		if got, want := password, test.password; got != want {
			t.Errorf("Want password %q, got %q", want, got)
		}
	}
}
