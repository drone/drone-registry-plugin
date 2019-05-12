// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"regexp"
	"sync"
	"time"
)

// regexp pattern used to match an ecr registry uri
var reRegistry = regexp.MustCompile("(.+).dkr.ecr.(.+).amazonaws.com")

type globalRegistry struct {
	sync.Mutex

	Address  string
	Username string
	Password string
	Email    string
	Token    string

	//
	// the below fields are for aws ecr repositories, whose
	// credentials need to be periodically refreshed.
	//

	expiry time.Time
	Access string `yaml:"aws_access_key_id"`
	Secret string `yaml:"aws_secret_access_key"`
}

// expires returns true if the registry credentials can expire.
func (r *globalRegistry) expires() bool {
	return reRegistry.MatchString(r.Address)
}

// expired returns true if the registry credentials are expired.
func (r *globalRegistry) expired() bool {
	r.Lock()
	defer r.Unlock()
	return time.Now().After(r.expiry)
}
