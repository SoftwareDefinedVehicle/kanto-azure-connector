// Copyright (c) 2022 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// http://www.eclipse.org/legal/epl-2.0
//
// SPDX-License-Identifier: EPL-2.0

package flags_test

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/eclipse-kanto/azure-connector/config"
	"github.com/eclipse-kanto/azure-connector/flags"

	"github.com/stretchr/testify/require"
)

func TestVersionParse(t *testing.T) {
	exitCall := false
	exit := func(_ int) {
		exitCall = true
	}

	f := flag.NewFlagSet("testing", flag.ContinueOnError)
	cmd := new(config.AzureSettings)
	flags.Add(f, cmd)

	args := []string{
		"-version",
	}

	require.NoError(t, flags.Parse(f, args, "0.0.0", exit))
	require.True(t, exitCall)
}

func TestInvalidFlag(t *testing.T) {
	f := flag.NewFlagSet("testing", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	cmd := new(config.AzureSettings)
	flags.Add(f, cmd)

	args := []string{
		"-invalid",
	}

	require.Error(t, flags.Parse(f, args, "0.0.0", os.Exit))
}