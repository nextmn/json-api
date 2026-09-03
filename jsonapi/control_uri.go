// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import (
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

func (u ControlURI) String() string {
	return u.URL.String()
}

func ParseControlURI(text string) (*ControlURI, error) {
	if len(text) == 0 {
		return nil, ErrEmptyURI
	}
	if text[len(text)-1] == '/' {
		return nil, ErrTrailingSlashURI
	}
	if u, err := url.ParseRequestURI(text); err != nil {
		return nil, err
	} else {
		return &ControlURI{
			URL: *u,
		}, nil
	}
}
