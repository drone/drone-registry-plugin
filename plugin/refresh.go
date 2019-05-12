// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin


import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// refreshFunc refreshes an ecr registry username and password.
type refreshFunc func(registry *globalRegistry) error

func defaultRefreshFunc(r *globalRegistry) error {
	account, region := parseRegistry(r.Address)

	var creds *credentials.Credentials
	if r.Access != "" {
		creds = credentials.NewStaticCredentials(r.Access, r.Secret, "")
	}
	sess := session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})

	service := ecr.New(sess, aws.NewConfig().WithRegion(region))

	params := &ecr.GetAuthorizationTokenInput{
		RegistryIds: []*string{&account},
	}

	resp, err := service.GetAuthorizationToken(params)
	if err != nil {
		return err
	}

	if len(resp.AuthorizationData) == 0 {
		return nil
	}

	user, pass, err := parseToken(*resp.AuthorizationData[0].AuthorizationToken)
	if err != nil {
		return err
	}

	r.Lock()
	r.Username = user
	r.Password = pass
	r.expiry = time.Now().Add(time.Hour)
	r.Unlock()

	return nil
}