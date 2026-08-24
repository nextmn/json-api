// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

func (msg MessageBase) WithError(err error) MessageWithError {
	return MessageWithError{
		MessageBase: msg,
		Error:       err.Error(),
	}
}

type MessageWithError struct {
	MessageBase
	Error string `json:"error"`
}
