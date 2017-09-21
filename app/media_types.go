// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/intervention-engine/ie/design
// --out=$(GOPATH)/src/github.com/intervention-engine/ie
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// A component score of an overall risk asessment (default view)
//
// Identifier: applicaiton/vnd.riskassessment+json; view=default
type RiskCategory struct {
	// Maximum possible value
	MaxValue *int `form:"max_value,omitempty" json:"max_value,omitempty" xml:"max_value,omitempty"`
	// Risk Category Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Risk Category Value
	Value *int `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Weight On Overall Risk Value
	Weight *int `form:"weight,omitempty" json:"weight,omitempty" xml:"weight,omitempty"`
}

// RiskCategoryCollection is the media type for an array of RiskCategory (default view)
//
// Identifier: applicaiton/vnd.riskassessment+json; type=collection; view=default
type RiskCategoryCollection []*RiskCategory

// A care team (default view)
//
// Identifier: application/vnd.careteam+json; view=default
type CareTeam struct {
	// Timestamp for care team creation
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unique care team ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Care team leader
	Leader *string `form:"leader,omitempty" json:"leader,omitempty" xml:"leader,omitempty"`
	// Care team name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// A care team (link view)
//
// Identifier: application/vnd.careteam+json; view=link
type CareTeamLink struct {
	// Unique care team ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Care team name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CareTeamCollection is the media type for an array of CareTeam (default view)
//
// Identifier: application/vnd.careteam+json; type=collection; view=default
type CareTeamCollection []*CareTeam

// Event media type (default view)
//
// Identifier: application/vnd.event+json; view=default
type Event struct {
	CreatedAt   *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	Datetime    *time.Time `form:"datetime,omitempty" json:"datetime,omitempty" xml:"datetime,omitempty"`
	DisplayName *string    `form:"display_name,omitempty" json:"display_name,omitempty" xml:"display_name,omitempty"`
	EndDate     *time.Time `form:"end_date,omitempty" json:"end_date,omitempty" xml:"end_date,omitempty"`
	ID          *string    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	NewValue    *float64   `form:"new_value,omitempty" json:"new_value,omitempty" xml:"new_value,omitempty"`
	OldValue    *float64   `form:"old_value,omitempty" json:"old_value,omitempty" xml:"old_value,omitempty"`
	StartDate   *time.Time `form:"start_date,omitempty" json:"start_date,omitempty" xml:"start_date,omitempty"`
	Type        *string    `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// EventCollection is the media type for an array of Event (default view)
//
// Identifier: application/vnd.event+json; type=collection; view=default
type EventCollection []*Event

// A single gathering of a care team at a specific time (default view)
//
// Identifier: application/vnd.huddle+json; view=default
type Huddle struct {
	// ID of the care team associated with this Huddle
	CareTeamID *string `form:"care_team_id,omitempty" json:"care_team_id,omitempty" xml:"care_team_id,omitempty"`
	// Creation timestamp
	Date *time.Time `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Unique Huddle ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// List of patients scheduled for this huddle
	Patients []*PatientHuddle `form:"patients,omitempty" json:"patients,omitempty" xml:"patients,omitempty"`
}

// HuddleCollection is the media type for an array of Huddle (default view)
//
// Identifier: application/vnd.huddle+json; type=collection; view=default
type HuddleCollection []*Huddle

// A patient (default view)
//
// Identifier: application/vnd.patient+json; view=default
type Patient struct {
	Address *Address `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Age of Patient
	Age *int `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// Birth Date of Patient
	BirthDate          *time.Time       `form:"birth_date,omitempty" json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	CurrentAllergies   []*ActiveElement `form:"current_allergies,omitempty" json:"current_allergies,omitempty" xml:"current_allergies,omitempty"`
	CurrentConditions  []*ActiveElement `form:"current_conditions,omitempty" json:"current_conditions,omitempty" xml:"current_conditions,omitempty"`
	CurrentMedications []*ActiveElement `form:"current_medications,omitempty" json:"current_medications,omitempty" xml:"current_medications,omitempty"`
	// Gender of Patient
	Gender *string `form:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	// Unique patient ID
	ID                   string          `form:"id" json:"id" xml:"id"`
	Name                 *Name           `form:"name" json:"name" xml:"name"`
	RecentRiskAssessment *RiskAssessment `form:"recent_risk_assessment,omitempty" json:"recent_risk_assessment,omitempty" xml:"recent_risk_assessment,omitempty"`
}

// Validate validates the Patient media type instance.
func (mt *Patient) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// A patient (link view)
//
// Identifier: application/vnd.patient+json; view=link
type PatientLink struct {
	// Unique patient ID
	ID   string `form:"id" json:"id" xml:"id"`
	Name *Name  `form:"name" json:"name" xml:"name"`
}

// Validate validates the PatientLink media type instance.
func (mt *PatientLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// A patient (list view)
//
// Identifier: application/vnd.patient+json; view=list
type PatientList struct {
	Address *Address `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Age of Patient
	Age *int `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// Birth Date of Patient
	BirthDate *time.Time `form:"birth_date,omitempty" json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// Gender of Patient
	Gender *string `form:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	// Unique patient ID
	ID                   string          `form:"id" json:"id" xml:"id"`
	Name                 *Name           `form:"name" json:"name" xml:"name"`
	RecentRiskAssessment *RiskAssessment `form:"recent_risk_assessment,omitempty" json:"recent_risk_assessment,omitempty" xml:"recent_risk_assessment,omitempty"`
}

// Validate validates the PatientList media type instance.
func (mt *PatientList) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// PatientCollection is the media type for an array of Patient (default view)
//
// Identifier: application/vnd.patient+json; type=collection; view=default
type PatientCollection []*Patient

// Validate validates the PatientCollection media type instance.
func (mt PatientCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// PatientCollection is the media type for an array of Patient (list view)
//
// Identifier: application/vnd.patient+json; type=collection; view=list
type PatientListCollection []*PatientList

// Validate validates the PatientListCollection media type instance.
func (mt PatientListCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// An single overall assessment score of a patient's risk for a risk service (default view)
//
// Identifier: application/vnd.riskassessment+json; view=default
type RiskAssessment struct {
	// Date assessment was created
	Date *time.Time `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Unique assessment identifier
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ID for the associated Risk breakdown the score is based on
	PieID *string `form:"pie_id,omitempty" json:"pie_id,omitempty" xml:"pie_id,omitempty"`
	// Identifier for risk service that produced the assessment
	RiskServiceID *string `form:"risk_service_id,omitempty" json:"risk_service_id,omitempty" xml:"risk_service_id,omitempty"`
	// Risk Score
	Value *float64 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// An single overall assessment score of a patient's risk for a risk service (list view)
//
// Identifier: application/vnd.riskassessment+json; view=list
type RiskAssessmentList struct {
	// Date assessment was created
	Date *time.Time `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Unique assessment identifier
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Identifier for risk service that produced the assessment
	RiskServiceID *string `form:"risk_service_id,omitempty" json:"risk_service_id,omitempty" xml:"risk_service_id,omitempty"`
	// Risk Score
	Value *float64 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// RiskAssessmentCollection is the media type for an array of RiskAssessment (default view)
//
// Identifier: application/vnd.riskassessment+json; type=collection; view=default
type RiskAssessmentCollection []*RiskAssessment

// RiskAssessmentCollection is the media type for an array of RiskAssessment (list view)
//
// Identifier: application/vnd.riskassessment+json; type=collection; view=list
type RiskAssessmentListCollection []*RiskAssessmentList

// Service providing risk scores for patients (default view)
//
// Identifier: applicatoin/vnd.riskservice+json; view=default
type RiskService struct {
	ID   *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	URL  *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// Service providing risk scores for patients (list view)
//
// Identifier: applicatoin/vnd.riskservice+json; view=list
type RiskServiceList struct {
	ID   *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// RiskServiceCollection is the media type for an array of RiskService (default view)
//
// Identifier: applicatoin/vnd.riskservice+json; type=collection; view=default
type RiskServiceCollection []*RiskService

// RiskServiceCollection is the media type for an array of RiskService (list view)
//
// Identifier: applicatoin/vnd.riskservice+json; type=collection; view=list
type RiskServiceListCollection []*RiskServiceList
