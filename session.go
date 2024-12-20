/**
 * @license
 * Copyright Google LLC All Rights Reserved.
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://opensource.org/licenses/MIT
 */

// Package session contains the logic to initialize the looker sdk session.
package session

import (
	"google3/third_party/looker_sdk_codegen/(internal)"
	"google3/third_party/looker_sdk_codegen/(internal)"
)

// InitSession initializes the looker sdk session.
func InitSession(apiIDKey string, apiSecretKey string, lookerLocation string, ssl bool) *v4.LookerSDK {
	settings := rtl.ApiSettings{
		BaseUrl:      lookerLocation,
		VerifySsl:    ssl,
		Timeout:      2000,
		ClientId:     apiIDKey,
		ClientSecret: apiSecretKey,
		ApiVersion:   "4.0",
	}

	authSession := rtl.NewAuthSession(settings)
	return v4.NewLookerSDK(authSession)
}
