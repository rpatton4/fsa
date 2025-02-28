// Robert Patton ("DEVELOPER") CONFIDENTIAL
// Unpublished Copyright (c) 2025 Robert Patton, All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of DEVELOPER. The
// intellectual and technical concepts contained herein are proprietary to DEVELOPER and may be
// covered by U.S. and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Dissemination of this information or reproduction of this material is strictly
// forbidden unless prior written permission is obtained from DEVELOPER. Access to the source code
// contained herein is hereby forbidden to anyone except current DEVELOPER employees, managers or
// contractors who have executed Confidentiality and Non-disclosure agreements explicitly covering
// such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of
// this source code, which includes information that is confidential and/or proprietary, and is a
// trade secret, of DEVELOPER. ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC  PERFORMANCE,
// OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF
// COMPANY IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND INTERNATIONAL TREATIES.
// THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY OR
// IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR
// SELL ANYTHING THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package fsamodels

import (
	"time"
)

// Person A representation of a person (versus student, parent etc.) with properties commonly
// used by FSA.  This was primarily influenced by the CommonRecord's XML schema definition of
// person.
type Person struct {
	FirstName   string
	MiddleName  string
	LastName    string
	NameSuffix  string
	DateOfBirth time.Time
	SSN         string
	ITIN        string
}
