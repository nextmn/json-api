// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package healthcheck

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

// Healthcheck allows to check status of the node
type Healthcheck struct {
	url       string
	userAgent string
}

// Status of the node
type Status struct {
	Ready bool `json:"ready"`
}

// Create a new Healthcheck
func NewHealthcheck(url url.URL, userAgent string) *Healthcheck {
	return &Healthcheck{
		url:       url.String(),
		userAgent: userAgent,
	}
}

// Run returns an error if the node status is not `ready`
func (h *Healthcheck) Run(ctx context.Context) error {
	client := http.Client{
		Timeout: 100 * time.Millisecond,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.url, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Error while creating http get request", "error", err)
		return err
	}
	req.Header.Add("User-Agent", h.userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")
	resp, err := client.Do(req)
	if err != nil {
		slog.InfoContext(ctx, "No HTTP response",
			"error", err,
			"remote-server", h.url,
		)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		slog.InfoContext(ctx, "HTTP response is not 200 OK",
			"error", err,
			"remote-server", h.url,
		)
		return err
	}
	var status Status
	if err := json.UnmarshalRead(resp.Body, &status); err != nil {
		slog.InfoContext(ctx, "Could not decode JSON response",
			"error", err,
			"remote-server", h.url,
		)
		return err
	}
	if !status.Ready {
		err := fmt.Errorf("server is not ready")
		slog.InfoContext(ctx, "Server is not ready",
			"error", err,
			"remote-server", h.url,
		)
		return err
	}
	return nil
}
