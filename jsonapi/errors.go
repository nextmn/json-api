// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import "errors"

var (
	ErrEmptyURI         = errors.New("control URI should not be empty")
	ErrTrailingSlashURI = errors.New("control URI should not contains trailing slash")
)
