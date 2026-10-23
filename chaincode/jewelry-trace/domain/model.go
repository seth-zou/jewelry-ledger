package domain

import (
	"errors"
	"time"
)

type AssetStatus string

const (
	StatusRegistered AssetStatus = "REGISTERED"
	StatusCertified  AssetStatus = "CERTIFIED"
	StatusInTransit  AssetStatus = "IN_TRANSIT"
	StatusSold       AssetStatus = "SOLD"
	StatusRevoked    AssetStatus = "REVOKED"
)

type JewelryAsset struct {
	AssetID     string
	PayloadHash string
	Status      AssetStatus
	UpdatedAt   time.Time
}

type TraceEvent struct {
	EventID     string
	AssetID     string
	EventType   string
	ActorOrgID  string
	PayloadHash string
	OccurredAt  time.Time
}

func (e TraceEvent) Validate() error {
	if e.EventID == "" || e.AssetID == "" || e.EventType == "" || e.ActorOrgID == "" || e.PayloadHash == "" {
		return errors.New("trace event requires event id, asset id, type, actor org, and payload hash")
	}
	return nil
}
