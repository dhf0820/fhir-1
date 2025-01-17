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

type DocumentReference struct {
	Id               string                                `json:"id" bson:"_id"`
	MasterIdentifier *Identifier                           `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier       []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject          *Reference                            `bson:"subject,omitempty" json:"subject,omitempty"`
	Type             *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	Class            *CodeableConcept                      `bson:"class,omitempty" json:"class,omitempty"`
	Author           []Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Custodian        *Reference                            `bson:"custodian,omitempty" json:"custodian,omitempty"`
	Authenticator    *Reference                            `bson:"authenticator,omitempty" json:"authenticator,omitempty"`
	Created          *FHIRDateTime                         `bson:"created,omitempty" json:"created,omitempty"`
	Indexed          *FHIRDateTime                         `bson:"indexed,omitempty" json:"indexed,omitempty"`
	Status           string                                `bson:"status,omitempty" json:"status,omitempty"`
	DocStatus        *CodeableConcept                      `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	RelatesTo        []DocumentReferenceRelatesToComponent `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description      string                                `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel    []CodeableConcept                     `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content          []DocumentReferenceContentComponent   `bson:"content,omitempty" json:"content,omitempty"`
	Context          *DocumentReferenceContextComponent    `bson:"context,omitempty" json:"context,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DocumentReference) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DocumentReference
	}{
		ResourceType:      "DocumentReference",
		DocumentReference: *resource,
	}
	return json.Marshal(x)
}

type DocumentReferenceRelatesToComponent struct {
	Code   string     `bson:"code,omitempty" json:"code,omitempty"`
	Target *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type DocumentReferenceContentComponent struct {
	Attachment *Attachment `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Format     []Coding    `bson:"format,omitempty" json:"format,omitempty"`
}

type DocumentReferenceContextComponent struct {
	Encounter         *Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Event             []CodeableConcept                          `bson:"event,omitempty" json:"event,omitempty"`
	Period            *Period                                    `bson:"period,omitempty" json:"period,omitempty"`
	FacilityType      *CodeableConcept                           `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept                           `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference                                 `bson:"sourcePatientInfo,omitempty" json:"sourcePatientInfo,omitempty"`
	Related           []DocumentReferenceContextRelatedComponent `bson:"related,omitempty" json:"related,omitempty"`
}

type DocumentReferenceContextRelatedComponent struct {
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref        *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}
