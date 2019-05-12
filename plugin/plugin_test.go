// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"testing"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/registry"
)

var noContext = context.Background()

func TestPlugin(t *testing.T) {
	plugin, err := parseBytes(
		[]byte(`[ {address: c.docker.com}, {address: d.docker.com} ]`),
	)
	if err != nil {
		t.Error(err)
		return
	}

	req := &registry.Request{
		Repo:  drone.Repo{},
		Build: drone.Build{},
	}
	list, err := plugin.List(noContext, req)
	if err != nil {
		t.Errorf("Expected combined registry list, got error %q", err)
		return
	}

	if got, want := list[0], plugin.registries[0]; got.Address != want.Address {
		t.Errorf("Expected correct precedence. Want %s, got %s", want.Address, got.Address)
	}
	if got, want := list[1], plugin.registries[1]; got.Address != want.Address {
		t.Errorf("Expected correct precedence. Want %s, got %s", want.Address, got.Address)
	}
}
