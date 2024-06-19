package models

import (
	"encoding/json"
)

// UpdateSubscriptionAffiliationIdRequest represents a UpdateSubscriptionAffiliationIdRequest struct.
// Request for updating a Subscription Affiliation Id
type UpdateSubscriptionAffiliationIdRequest struct {
	GatewayAffiliationId string `json:"gateway_affiliation_id"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionAffiliationIdRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionAffiliationIdRequest objects.
func (u *UpdateSubscriptionAffiliationIdRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionAffiliationIdRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionAffiliationIdRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["gateway_affiliation_id"] = u.GatewayAffiliationId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionAffiliationIdRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionAffiliationIdRequest objects.
func (u *UpdateSubscriptionAffiliationIdRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		GatewayAffiliationId string `json:"gateway_affiliation_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.GatewayAffiliationId = temp.GatewayAffiliationId
	return nil
}
