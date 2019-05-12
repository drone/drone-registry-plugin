// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"encoding/base64"
	"strings"

	"github.com/drone/drone-go/drone"
)

// helper function converts from the internal registry representation
// to the format registry required by the drone server.
func convertRegistry(from *globalRegistry) *drone.Registry {
	from.Lock()
	defer from.Unlock()
	return &drone.Registry{
		Address:  from.Address,
		Username: from.Username,
		Password: from.Password,
		Email:    from.Email,
		Token:    from.Token,
	}
}

// helper function parses the aws registry and returns the
// account number and region.
func parseRegistry(s string) (account, region string) {
	matches := reRegistry.FindStringSubmatch(s)
	if len(matches) == 3 {
		account = matches[1]
		region = matches[2]
	}
	return
}

// helper function parses the aws registry authentication token
// and returns the decoded docker username and password.
func parseToken(s string) (username string, password string, err error) {
	var data []byte
	data, err = base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	token := strings.SplitN(string(data), ":", 2)
	username = token[0]
	password = token[1]
	return
}
