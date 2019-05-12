// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"io/ioutil"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/registry"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// New returns a new registry auth plugin that sources registry
// credentials from a static list, typically loaded from file.
func New() registry.Plugin {
	return &plugin{}
}

type plugin struct {
	registries []*globalRegistry
	refresher  refreshFunc
}

// Load loads a registry plugin from file.
func Load(path string) (registry.Plugin, error) {
	data, ferr := ioutil.ReadFile(path)
	if ferr != nil {
		return nil, ferr
	}
	return parseBytes(data)
}

func parseBytes(data []byte) (*plugin, error) {
	g := new(plugin)
	g.refresher = defaultRefreshFunc
	err := yaml.Unmarshal(data, &g.registries)
	if err != nil {
		return nil, err
	}
	for _, r := range g.registries {
		if r.Access != "" || r.Secret != "" {
			logrus.Infof("registry credentials: loaded: aws@%s %s", r.Address, r.Access)
		} else {
			logrus.Infof("registry credentials: loaded: %s@%s", r.Username, r.Address)
		}
	}
	return g, nil
}

func (p *plugin) List(context.Context, *registry.Request) ([]*drone.Registry, error) {
	var list []*drone.Registry
	for _, registry := range p.registries {
		if registry.expires() && registry.expired() {
			logrus.Debugf("registry credentials: refresh: %s", registry.Access)
			if err := p.refresher(registry); err != nil {
				return nil, err
			}
		}
		list = append(list, convertRegistry(registry))
	}
	return list, nil
}
