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

type NutritionOrder struct {
	Id                     string                                 `json:"id" bson:"_id"`
	Patient                *Reference                             `bson:"patient,omitempty" json:"patient,omitempty"`
	Orderer                *Reference                             `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Identifier             []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Encounter              *Reference                             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateTime               *FHIRDateTime                          `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Status                 string                                 `bson:"status,omitempty" json:"status,omitempty"`
	AllergyIntolerance     []Reference                            `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept                      `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept                      `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDietComponent       `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplementComponent    `bson:"supplement,omitempty" json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormulaComponent `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *NutritionOrder) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		NutritionOrder
	}{
		ResourceType:   "NutritionOrder",
		NutritionOrder: *resource,
	}
	return json.Marshal(x)
}

type NutritionOrderOralDietComponent struct {
	Type                 []CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Schedule             []Timing                                  `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrientComponent `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTextureComponent  `bson:"texture,omitempty" json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                         `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	Instruction          string                                    `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type NutritionOrderOralDietNutrientComponent struct {
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Amount   *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}

type NutritionOrderOralDietTextureComponent struct {
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	FoodType *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}

type NutritionOrderSupplementComponent struct {
	Type        *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ProductName string           `bson:"productName,omitempty" json:"productName,omitempty"`
	Schedule    []Timing         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity    *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Instruction string           `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type NutritionOrderEnteralFormulaComponent struct {
	BaseFormulaType           *CodeableConcept                                      `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    string                                                `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                                      `bson:"additiveType,omitempty" json:"additiveType,omitempty"`
	AdditiveProductName       string                                                `bson:"additiveProductName,omitempty" json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                             `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                                      `bson:"routeofAdministration,omitempty" json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministrationComponent `bson:"administration,omitempty" json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                             `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction string                                                `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdministrationComponent struct {
	Schedule           *Timing   `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity           *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	RateSimpleQuantity *Quantity `bson:"rateSimpleQuantity,omitempty" json:"rateSimpleQuantity,omitempty"`
	RateRatio          *Ratio    `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
}
