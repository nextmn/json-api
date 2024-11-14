// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package healthcheck

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Healthcheck allows to check status of the node
type Healthcheck struct {
	uri       string
	userAgent string
}

// Status of the node
type Status struct {
	Ready bool `json:"ready"`
}

// Create a new Healthcheck
func NewHealthcheck(uri string, userAgent string) *Healthcheck {
	return &Healthcheck{
		uri:       uri,
		userAgent: userAgent,
	}
}

// Run returns an error if the node status is not `ready`
func (h *Healthcheck) Run(ctx context.Context) error {
	client := http.Client{
		Timeout: 100 * time.Millisecond,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.uri+"/status", nil)
	if err != nil {
		logrus.WithError(err).Error("Error while creating http get request")
		return err
	}
	req.Header.Add("User-Agent", h.userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"remote-server": h.uri}).WithError(err).Info("No http response")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logrus.WithFields(logrus.Fields{"remote-server": h.uri}).WithError(err).Info("Http response is not 200 OK")
		return err
	}
	decoder := json.NewDecoder(resp.Body)
	var status Status
	if err := decoder.Decode(&status); err != nil {
		logrus.WithFields(logrus.Fields{"remote-server": h.uri}).WithError(err).Info("Could not decode json response")
		return err
	}
	if !status.Ready {
		err := fmt.Errorf("Server is not ready")
		logrus.WithFields(logrus.Fields{"remote-server": h.uri}).WithError(err).Info("Server is not ready")
		return err
	}
	return nil
}
