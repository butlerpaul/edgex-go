//
// Copyright (c) 2018
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"testing"

	"github.com/edgexfoundry/edgex-go/core/metadata"
	"github.com/edgexfoundry/edgex-go/pkg/config"
)

func TestToml(t *testing.T) {
	configuration := &metadata.ConfigurationStruct{}
	if err := config.VerifyTomlFiles(configuration); err != nil {
		t.Fatalf("%v", err)
	}
}
