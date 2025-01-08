package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewAuditEventsAdminApiActivity(webhook webhook.ValidWebhook) (*AuditEventsAdminApiActivity, error) {
	payload := &AuditEventsAdminApiActivityPayload{}
	info, err := webhook.GetInfo()
	if err != nil {
		return nil, err
	}
	reader, err := webhook.GetReader()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(payload); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}
	return &AuditEventsAdminApiActivity{
		info: info,
		data: *payload,
	}, nil
}

type AuditEventsAdminApiActivity struct {
	info webhook.Info
	data AuditEventsAdminApiActivityPayload
}

func (webhook *AuditEventsAdminApiActivity) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *AuditEventsAdminApiActivity) GetData() (AuditEventsAdminApiActivityPayload, error) {
	return webhook.data, nil
}

type AuditEventsAdminApiActivityPayload struct {
	Events []struct {
		Event struct {
			Action string `json:"action"`
			Actor  struct {
				ActorIdentifier int64 `json:"actor_identifier"`
				ActorMetadata   struct {
					APIVersionRequested string `json:"api_version_requested"`
					APIVersionServed    string `json:"api_version_served"`
					AppName             string `json:"app_name"`
					EcosystemCategory   string `json:"ecosystem_category"`
				} `json:"actor_metadata"`
				ActorType  string `json:"actor_type"`
				OnBehalfOf struct {
					UserEmail string `json:"user_email"`
				} `json:"on_behalf_of"`
			} `json:"actor"`
			AdditionalMetadata struct {
				ErrorCodes    []interface{} `json:"error_codes"`
				MutationNames []interface{} `json:"mutation_names"`
				Query         string        `json:"query"`
				RequestType   string        `json:"request_type"`
				Variables     struct {
					First int64 `json:"first"`
				} `json:"variables"`
			} `json:"additional_metadata"`
			Context struct {
				ContextIdentifier string `json:"context_identifier"`
				ContextMetadata   struct {
					APIRequestFailed bool    `json:"api_request_failed"`
					ClientIP         string  `json:"client_ip"`
					ContentType      string  `json:"content_type"`
					ResponseCode     int64   `json:"response_code"`
					ResponseTimeMs   float64 `json:"response_time_ms"`
					UserAgent        string  `json:"user_agent"`
				} `json:"context_metadata"`
				ContextType string `json:"context_type"`
			} `json:"context"`
			Subject struct {
				SubjectIdentifier string   `json:"subject_identifier"`
				SubjectMetadata   struct{} `json:"subject_metadata"`
				SubjectType       string   `json:"subject_type"`
			} `json:"subject"`
			Timestamp string `json:"timestamp"`
		} `json:"event"`
		Time int64 `json:"time"`
	} `json:"events"`
}
