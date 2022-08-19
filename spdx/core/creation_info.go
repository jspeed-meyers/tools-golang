// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package core

import "github.com/spdx/tools-golang/spdx/common"

// CreationInfo is a Document Creation Information section of an
// SPDX Document for version 2.2 of the spec.
type CreationInfo struct {
	// CoreMetadata includes additional information to map back and forth
	// to other SPDX types
	CoreMetadata Metadata

	// 2.7: License List Version
	// Cardinality: optional, one
	LicenseListVersion string `json:"licenseListVersion"`

	// 2.8: Creators: may have multiple keys for Person, Organization
	//      and/or Tool
	// Cardinality: mandatory, one or many
	Creators []common.Creator `json:"creators"`

	// 2.9: Created: data format YYYY-MM-DDThh:mm:ssZ
	// Cardinality: mandatory, one
	Created string `json:"created"`

	// 2.10: Creator Comment
	// Cardinality: optional, one
	CreatorComment string `json:"comment"`
}
