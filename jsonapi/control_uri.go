// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type ControlURI struct {
	url.URL
}

func (u *ControlURI) UnmarshalText(text []byte) error {
	if a, err := ParseControlURI(string(text[:])); err != nil {
		return err
	} else {
		*u = *a
	}
	return nil
}

func (u ControlURI) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func ParseControlURI(text string) (*ControlURI, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("Control URI should not be empty.")
	}
	if text[len(text)-1] == '/' {
		return nil, fmt.Errorf("Control URI should not contains trailing slash.")
	}
	if u, err := url.ParseRequestURI(text); err != nil {
		return nil, err
	} else {
		return &ControlURI{
			URL: *u,
		}, nil
	}
}
