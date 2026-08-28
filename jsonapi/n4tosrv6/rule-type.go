// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import "errors"

var (
	ErrInvalidRuleType = errors.New("invalid rule type")
)

//go:generate stringer -type RuleType -linecomment -output rule-type_string.go

type RuleType int

const (
	RuleTypeUplink             RuleType = iota // uplink
	RuleTypeDownlink                           // downlink
	RuleTypeDownlinkForwarding                 // downlink-forwarding
)

func (r *RuleType) UnmarshalText(text []byte) error {
	switch string(text) {
	case RuleTypeUplink.String():
		*r = RuleTypeUplink
	case RuleTypeDownlink.String():
		*r = RuleTypeDownlink
	case RuleTypeDownlinkForwarding.String():
		*r = RuleTypeDownlinkForwarding
	default:
		return ErrInvalidRuleType
	}
	return nil
}

func (r RuleType) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil

}
