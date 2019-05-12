// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"testing"
	"time"
)

func Test_expires(t *testing.T) {
	tests := []struct {
		address string
		expires bool
	}{
		{"123456789.dkr.ecr.us-east-1.amazonaws.com", true},
		{"123456789.foo.bar.us-east-1.amazonaws.com", false},
	}
	for _, test := range tests {
		a := globalRegistry{Address: test.address}
		if got, want := a.expires(), test.expires; got != want {
			t.Errorf("Want expires %v, got %v", want, got)
		}
	}
}

func Test_expired(t *testing.T) {
	tests := []struct {
		expiry  time.Time
		expires bool
	}{
		{time.Now().Add(-time.Hour - time.Minute), true},
		{time.Now().Add(time.Hour), false},
	}
	for _, test := range tests {
		a := globalRegistry{expiry: test.expiry}
		if got, want := a.expired(), test.expires; got != want {
			t.Errorf("Want expired %v, got %v", want, got)
		}
	}
}
