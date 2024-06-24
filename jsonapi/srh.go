// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

import "fmt"

type SRH []*Segment

func NewSRH(segments []string) (*SRH, error) {
	if len(segments) < 1 {
		return nil, fmt.Errorf("SRH should contain at least one segment")
	}
	srh := make(SRH, 0)
	for _, s := range segments {
		if seg_n, err := NewSegment(s); err == nil {
			srh = append(srh, seg_n)
		} else {
			return nil, err
		}
	}
	return &srh, nil
}
