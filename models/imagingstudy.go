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

type ImagingStudy struct {
	Id                string                        `json:"id" bson:"_id"`
	Started           *FHIRDateTime                 `bson:"started,omitempty" json:"started,omitempty"`
	Patient           *Reference                    `bson:"patient,omitempty" json:"patient,omitempty"`
	Uid               string                        `bson:"uid,omitempty" json:"uid,omitempty"`
	Accession         *Identifier                   `bson:"accession,omitempty" json:"accession,omitempty"`
	Identifier        []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Order             []Reference                   `bson:"order,omitempty" json:"order,omitempty"`
	ModalityList      []Coding                      `bson:"modalityList,omitempty" json:"modalityList,omitempty"`
	Referrer          *Reference                    `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Availability      string                        `bson:"availability,omitempty" json:"availability,omitempty"`
	Url               string                        `bson:"url,omitempty" json:"url,omitempty"`
	NumberOfSeries    *uint32                       `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	NumberOfInstances *uint32                       `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Procedure         []Reference                   `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Interpreter       *Reference                    `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Description       string                        `bson:"description,omitempty" json:"description,omitempty"`
	Series            []ImagingStudySeriesComponent `bson:"series,omitempty" json:"series,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImagingStudy) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ImagingStudy
	}{
		ResourceType: "ImagingStudy",
		ImagingStudy: *resource,
	}
	return json.Marshal(x)
}

type ImagingStudySeriesComponent struct {
	Number            *uint32                               `bson:"number,omitempty" json:"number,omitempty"`
	Modality          *Coding                               `bson:"modality,omitempty" json:"modality,omitempty"`
	Uid               string                                `bson:"uid,omitempty" json:"uid,omitempty"`
	Description       string                                `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfInstances *uint32                               `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Availability      string                                `bson:"availability,omitempty" json:"availability,omitempty"`
	Url               string                                `bson:"url,omitempty" json:"url,omitempty"`
	BodySite          *Coding                               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Laterality        *Coding                               `bson:"laterality,omitempty" json:"laterality,omitempty"`
	Started           *FHIRDateTime                         `bson:"started,omitempty" json:"started,omitempty"`
	Instance          []ImagingStudySeriesInstanceComponent `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingStudySeriesInstanceComponent struct {
	Number   *uint32      `bson:"number,omitempty" json:"number,omitempty"`
	Uid      string       `bson:"uid,omitempty" json:"uid,omitempty"`
	SopClass string       `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Type     string       `bson:"type,omitempty" json:"type,omitempty"`
	Title    string       `bson:"title,omitempty" json:"title,omitempty"`
	Content  []Attachment `bson:"content,omitempty" json:"content,omitempty"`
}
