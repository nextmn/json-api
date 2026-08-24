// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import (
	"net"
)

type SRH []Segment

func ParseSRH(segmentZero string, segments ...string) (SRH, error) {
	seg0, err := ParseSegment(segmentZero)
	if err != nil {
		return nil, err
	}
	srh := SRH{seg0}
	for _, s := range segments {
		seg_n, err := ParseSegment(s)
		if err != nil {
			return nil, err
		}
		srh = append(srh, seg_n)
	}
	return srh, nil
}

func (srh SRH) AsSlice() []net.IP {
	r := make([]net.IP, len(srh))
	for i, seg := range srh {
		r[i] = seg.AsSlice()
	}
	return r
}
