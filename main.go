// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/drone/drone-go/plugin/registry"
	"github.com/drone/drone-registry-plugin/plugin"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type spec struct {
	Debug   bool   `envconfig:"DRONE_DEBUG"`
	Address string `envconfig:"DRONE_ADDRESS"     default:":3000"`
	Secret  string `envconfig:"DRONE_SECRET"      required:"true"`
	Config  string `envconfig:"DRONE_CONFIG_FILE" required:"true"`
}

func main() {
	spec := new(spec)
	err := envconfig.Process("", spec)
	if err != nil {
		logrus.WithError(err).Fatalln("invalid configuration")
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	plugin, err := plugin.Load(spec.Config)
	if err != nil {
		logrus.WithError(err).Fatalln("cannot load configuration")
	}

	handler := registry.Handler(
		spec.Secret, plugin, logrus.StandardLogger())

	logrus.Infof("server listening on address %s", spec.Address)

	http.Handle("/", handler)
	logrus.Fatal(http.ListenAndServe(spec.Address, nil))
}
