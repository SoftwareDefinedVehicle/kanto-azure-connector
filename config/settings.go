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

package config

import (
	"github.com/pkg/errors"

	"github.com/eclipse-kanto/suite-connector/config"
	"github.com/eclipse-kanto/suite-connector/logger"
	"github.com/eclipse-kanto/suite-connector/util"
)

// AzureSettings represents all configurable data that is used to setup the Cloud Agent.
type AzureSettings struct {
	TenantID                     string `json:"tenantId"`
	ConnectionString             string `json:"connectionString"`
	SASTokenValidity             string `json:"sasTokenValidity"`
	MessageMapperConfig          string `json:"messageMapperConfig"`
	AllowedLocalTopicsList       string `json:"allowedLocalTopicsList"`
	AllowedCloudMessageTypesList string `json:"allowedCloudMessageTypesList"`
	IDScope                      string `json:"idScope"`

	config.LocalConnectionSettings
	logger.LogSettings
	config.TLSSettings
}

// DefaultSettings returns the Azure connector default settings.
func DefaultSettings() *AzureSettings {
	def := config.DefaultSettings()
	defAzureSettings := &AzureSettings{
		TenantID:                "defaultTenant",
		SASTokenValidity:        "1h",
		MessageMapperConfig:     "message-mapper-config.json",
		LocalConnectionSettings: def.LocalConnectionSettings,
		TLSSettings: config.TLSSettings{
			CACert: def.CACert,
		},
		LogSettings: def.LogSettings,
	}
	defAzureSettings.LogFile = "logs/azure-connector.log"
	defAzureSettings.LogFileMaxAge = 28
	return defAzureSettings
}

// Validate validates the settings.
func (settings *AzureSettings) Validate() error {
	if err := settings.LogSettings.Validate(); err != nil {
		return err
	}

	if err := settings.LocalConnectionSettings.Validate(); err != nil {
		return err
	}

	if len(settings.CACert) > 0 && !util.FileExists(settings.CACert) {
		return errors.New("failed to read CA certificates file")
	}
	return nil
}