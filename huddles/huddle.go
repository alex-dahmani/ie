package huddles

import (
	"fmt"
	"time"

	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

// Huddle provides convenient functions on a Group to get access to extended huddle data fields
type Huddle models.Group

// NewHuddle constructs a new Huddle using the provided information
func NewHuddle(name string, leaderID string, date time.Time) *Huddle {
	tru := true
	huddle := Huddle(models.Group{
		DomainResource: models.DomainResource{
			Resource: models.Resource{
				Id:           bson.NewObjectId().Hex(),
				ResourceType: "Group",
				Meta: &models.Meta{
					Profile: []string{"http://interventionengine.org/fhir/profile/huddle"},
				},
			},
			Extension: []models.Extension{
				{
					Url:           "http://interventionengine.org/fhir/extension/group/activeDateTime",
					ValueDateTime: &models.FHIRDateTime{Time: date, Precision: models.Precision(models.Date)},
				},
				{
					Url: "http://interventionengine.org/fhir/extension/group/leader",
					ValueReference: &models.Reference{
						Reference:    "Practitioner/" + leaderID,
						ReferencedID: leaderID,
						Type:         "Practitioner",
						External:     new(bool),
					},
				},
			},
		},
		Type:   "person",
		Actual: &tru,
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "http://interventionengine.org/fhir/cs/huddle", Code: "HUDDLE"},
			},
			Text: "Huddle",
		},
		Name: name,
	})

	return &huddle
}

// IsHuddle checks the Group's code to ensure it has the proper Huddle code
func (h *Huddle) IsHuddle() bool {
	return h.Code.MatchesCode("http://interventionengine.org/fhir/cs/huddle", "HUDDLE")
}

// ActiveDateTime returns the huddle's active datetime (or nil if there is not one)
func (h *Huddle) ActiveDateTime() *models.FHIRDateTime {
	activeDT := findExtension(h.Extension, "http://interventionengine.org/fhir/extension/group/activeDateTime")
	if activeDT != nil {
		return activeDT.ValueDateTime
	}
	return nil
}

// Leader returns the huddle's leader (or nil if there is not one)
func (h *Huddle) Leader() *models.Reference {
	leader := findExtension(h.Extension, "http://interventionengine.org/fhir/extension/group/leader")
	if leader != nil {
		return leader.ValueReference
	}
	return nil
}

// HuddleMembers returns a slice of HuddleMembers associated to this huddle
func (h *Huddle) HuddleMembers() []HuddleMember {
	members := make([]HuddleMember, len(h.Member))
	for i := range h.Member {
		members[i] = HuddleMember(h.Member[i])
	}
	return members
}

// FindHuddleMember returns the huddle member with the specified ID (or nil if the patient is not in the huddle)
func (h *Huddle) FindHuddleMember(patientID string) *HuddleMember {
	for i := range h.Member {
		if h.Member[i].Entity.ReferencedID == patientID {
			hm := HuddleMember(h.Member[i])
			return &hm
		}
	}
	return nil
}

// AddHuddleMemberDueToRiskScore adds the patient to the huddle using RISK_SCORE as the reason.  If the
// patient is already in the huddle, nothing will be updated.
func (h *Huddle) AddHuddleMemberDueToRiskScore(patientID string) {
	h.addHuddleMember(patientID, &models.CodeableConcept{
		Coding: []models.Coding{
			{System: "http://interventionengine.org/fhir/cs/huddle-member-reason", Code: "RISK_SCORE"},
		},
		Text: "Risk Score Warrants Discussion",
	})
}

// AddHuddleMemberDueToRecentEvent adds the patient to the huddle using RECENT_ENCOUNTER event code as the reason.
// If the patient is already in the huddle, nothing will be updated.
func (h *Huddle) AddHuddleMemberDueToRecentEvent(patientID string, code EventCode) {
	h.addHuddleMember(patientID, &models.CodeableConcept{
		Coding: []models.Coding{
			{System: "http://interventionengine.org/fhir/cs/huddle-member-reason", Code: "RECENT_ENCOUNTER"},
		},
		Text: code.Name,
	})
}

// AddHuddleMemberDueToRollOver adds the patient to the huddle using the ROLLOVER and previous reason.
// If the patient is already in the huddle, nothing will be updated.
func (h *Huddle) AddHuddleMemberDueToRollOver(patientID string, from time.Time, previousReason *models.CodeableConcept) {
	var reason string
	if previousReason.MatchesCode("http://interventionengine.org/fhir/cs/huddle-member-reason", "ROLLOVER") {
		reason = previousReason.Text
	} else if previousReason.MatchesCode("http://interventionengine.org/fhir/cs/huddle-member-reason", "MANUAL_ADDITION") {
		reason = fmt.Sprintf("Rolled Over from %s (Manually Added - %s)", from.Format("Jan 2"), previousReason.Text)
	} else {
		reason = fmt.Sprintf("Rolled Over from %s (%s)", from.Format("Jan 2"), previousReason.Text)
	}
	h.addHuddleMember(patientID, &models.CodeableConcept{
		Coding: []models.Coding{
			{System: "http://interventionengine.org/fhir/cs/huddle-member-reason", Code: "ROLLOVER"},
		},
		Text: reason,
	})
}

func (h *Huddle) addHuddleMember(patientID string, reason *models.CodeableConcept) {
	// First look to see if the patient is already in the group, and if so, return
	if h.FindHuddleMember(patientID) != nil {
		return
	}

	// The patient is not yet in the group, so add him/her
	h.Member = append(h.Member, models.GroupMemberComponent{
		BackboneElement: models.BackboneElement{
			Element: models.Element{
				Extension: []models.Extension{
					{
						Url:                  "http://interventionengine.org/fhir/extension/group/member/reason",
						ValueCodeableConcept: reason,
					},
				},
			},
		},
		Entity: &models.Reference{
			Reference:    "Patient/" + patientID,
			ReferencedID: patientID,
			Type:         "Patient",
			External:     new(bool),
		},
	})
}

// RemoveHuddleMember removes the requested huddle member and returns the removed member.
// If no matching member is found, it returns nil.
func (h *Huddle) RemoveHuddleMember(patientID string) *HuddleMember {
	for i := range h.Member {
		if h.Member[i].Entity.ReferencedID == patientID {
			hm := HuddleMember(h.Member[i])
			h.Member = append(h.Member[:i], h.Member[i+1:]...)
			return &hm
		}
	}
	return nil
}

// HuddleMember provides convenient functions on a GroupMemberComponent to get access to extended huddle data fields
type HuddleMember models.GroupMemberComponent

// ID returns the members ID
func (h *HuddleMember) ID() string {
	if h.Entity == nil {
		return ""
	}
	return h.Entity.ReferencedID
}

// Reason returns the reason the member was added to the huddle (or nil if the reason isn't set)
func (h *HuddleMember) Reason() *models.CodeableConcept {
	reason := findExtension(h.Extension, "http://interventionengine.org/fhir/extension/group/member/reason")
	if reason != nil {
		return reason.ValueCodeableConcept
	}
	return nil
}

// ReasonIsManuallyAdded indicates if the member reason is due to the patient being manually added to the huddle
func (h *HuddleMember) ReasonIsManuallyAdded() bool {
	reason := h.Reason()
	return reason != nil && reason.MatchesCode("http://interventionengine.org/fhir/cs/huddle-member-reason", "MANUAL_ADDITION")
}

// ReasonIsRecentEncounter indicates if the member reason is due to a recent significant encounter
func (h *HuddleMember) ReasonIsRecentEncounter() bool {
	reason := h.Reason()
	return reason != nil && reason.MatchesCode("http://interventionengine.org/fhir/cs/huddle-member-reason", "RECENT_ENCOUNTER")
}

// ReasonIsRiskScore indicates if the member reason is due to the patient's current risk score
func (h *HuddleMember) ReasonIsRiskScore() bool {
	reason := h.Reason()
	return reason != nil && reason.MatchesCode("http://interventionengine.org/fhir/cs/huddle-member-reason", "RISK_SCORE")
}

// ReasonIsRollOver indicates if the member reason is due to roll over from a previous huddle
func (h *HuddleMember) ReasonIsRollOver() bool {
	reason := h.Reason()
	return reason != nil && reason.MatchesCode("http://interventionengine.org/fhir/cs/huddle-member-reason", "ROLLOVER")
}

// Reviewed returns the date that the member was reviewed for this huddle (or nil if they haven't been reviewed)
func (h *HuddleMember) Reviewed() *models.FHIRDateTime {
	reviewed := findExtension(h.Extension, "http://interventionengine.org/fhir/extension/group/member/reviewed")
	if reviewed != nil {
		return reviewed.ValueDateTime
	}
	return nil
}

func findExtension(ext []models.Extension, extURL string) *models.Extension {
	for i := range ext {
		if ext[i].Url == extURL {
			return &ext[i]
		}
	}
	return nil
}
