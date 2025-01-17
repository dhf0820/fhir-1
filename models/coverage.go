// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import "encoding/json"

type Coverage struct {
	Id           string       `json:"id" bson:"_id"`
	Issuer       *Reference   `bson:"issuer,omitempty" json:"issuer,omitempty"`
	Bin          *Identifier  `bson:"bin,omitempty" json:"bin,omitempty"`
	Period       *Period      `bson:"period,omitempty" json:"period,omitempty"`
	Type         *Coding      `bson:"type,omitempty" json:"type,omitempty"`
	SubscriberId *Identifier  `bson:"subscriberId,omitempty" json:"subscriberId,omitempty"`
	Identifier   []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Group        string       `bson:"group,omitempty" json:"group,omitempty"`
	Plan         string       `bson:"plan,omitempty" json:"plan,omitempty"`
	SubPlan      string       `bson:"subPlan,omitempty" json:"subPlan,omitempty"`
	Dependent    *uint32      `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Sequence     *uint32      `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Subscriber   *Reference   `bson:"subscriber,omitempty" json:"subscriber,omitempty"`
	Network      *Identifier  `bson:"network,omitempty" json:"network,omitempty"`
	Contract     []Reference  `bson:"contract,omitempty" json:"contract,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Coverage) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Coverage
	}{
		ResourceType: "Coverage",
		Coverage:     *resource,
	}
	return json.Marshal(x)
}
