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

type Composition struct {
	Id              string                         `json:"id" bson:"_id"`
	Identifier      *Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Date            *FHIRDateTime                  `bson:"date,omitempty" json:"date,omitempty"`
	Type            *CodeableConcept               `bson:"type,omitempty" json:"type,omitempty"`
	Class           *CodeableConcept               `bson:"class,omitempty" json:"class,omitempty"`
	Title           string                         `bson:"title,omitempty" json:"title,omitempty"`
	Status          string                         `bson:"status,omitempty" json:"status,omitempty"`
	Confidentiality string                         `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Subject         *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Author          []Reference                    `bson:"author,omitempty" json:"author,omitempty"`
	Attester        []CompositionAttesterComponent `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian       *Reference                     `bson:"custodian,omitempty" json:"custodian,omitempty"`
	Event           []CompositionEventComponent    `bson:"event,omitempty" json:"event,omitempty"`
	Encounter       *Reference                     `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Section         []CompositionSectionComponent  `bson:"section,omitempty" json:"section,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Composition) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Composition
	}{
		ResourceType: "Composition",
		Composition:  *resource,
	}
	return json.Marshal(x)
}

type CompositionAttesterComponent struct {
	Mode  []string      `bson:"mode,omitempty" json:"mode,omitempty"`
	Time  *FHIRDateTime `bson:"time,omitempty" json:"time,omitempty"`
	Party *Reference    `bson:"party,omitempty" json:"party,omitempty"`
}

type CompositionEventComponent struct {
	Code   []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Detail []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CompositionSectionComponent struct {
	Title       string                        `bson:"title,omitempty" json:"title,omitempty"`
	Code        *CodeableConcept              `bson:"code,omitempty" json:"code,omitempty"`
	Text        *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Mode        string                        `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy   *CodeableConcept              `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Entry       []Reference                   `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason *CodeableConcept              `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section     []CompositionSectionComponent `bson:"section,omitempty" json:"section,omitempty"`
}
