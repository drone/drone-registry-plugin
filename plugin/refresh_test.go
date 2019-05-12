// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"os"
	"testing"
)

//
// These unit tests are used for integration and require
// aws credentials. They are skipped by default.
//

var (
	testAccess   = os.Getenv("AWS_ACCESS_KEY_ID")
	testSecret   = os.Getenv("AWS_SECRET_ACCESS_KEY")
	testRegistry = os.Getenv("AWS_REPOSITORY_URI")
)

func Test_refresh(t *testing.T) {
	if testAccess == "" || testSecret == "" || testRegistry == "" {
		t.SkipNow()
		return
	}

	r := &globalRegistry{
		Address: testRegistry,
		Access:  testAccess,
		Secret:  testSecret,
	}

	if err := defaultRefreshFunc(r); err != nil {
		t.Error(err)
		return
	}

	if got, want := r.Username, "AWS"; got != want {
		t.Errorf("Got registry username %q, want %q", got, want)
	}

	if r.Password == "" {
		t.Errorf("Expect registry password, got zero value")
	}

	if r.expiry.Unix() == 0 {
		t.Error("Expect expiry date incremented, got zero value")
	}

	if r.expired() {
		t.Error("Expect expired false, got true")
	}
}