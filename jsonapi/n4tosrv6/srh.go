// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import (
	"fmt"
	"net"
)

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

func (srh SRH) AsSlice() []net.IP {
	r := make([]net.IP, len(srh))
	for i, seg := range srh {
		r[i] = seg.AsSlice()
	}
	return r
}
