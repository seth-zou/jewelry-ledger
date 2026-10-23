package domain

import (
	"errors"
	"time"
)

type LedgerStore interface {
	GetAsset(assetID string) (JewelryAsset, bool)
	PutAsset(asset JewelryAsset)
	AppendEvent(event TraceEvent)
	Events(assetID string) []TraceEvent
}

type TraceService struct {
	store LedgerStore
	clock func() time.Time
}

func NewTraceService(store LedgerStore, clock func() time.Time) TraceService {
	if clock == nil {
		clock = time.Now
	}
	return TraceService{store: store, clock: clock}
}

func (s TraceService) RegisterAsset(assetID string, payloadHash string) (JewelryAsset, error) {
	if assetID == "" || payloadHash == "" {
		return JewelryAsset{}, errors.New("asset id and payload hash are required")
	}
	if _, exists := s.store.GetAsset(assetID); exists {
		return JewelryAsset{}, errors.New("asset already exists")
	}
	asset := JewelryAsset{AssetID: assetID, PayloadHash: payloadHash, Status: StatusRegistered, UpdatedAt: s.clock()}
	s.store.PutAsset(asset)
	return asset, nil
}

func (s TraceService) AppendTraceEvent(event TraceEvent) (JewelryAsset, error) {
	if err := event.Validate(); err != nil {
		return JewelryAsset{}, err
	}
	asset, exists := s.store.GetAsset(event.AssetID)
	if !exists {
		return JewelryAsset{}, errors.New("asset not found")
	}
	asset.Status = statusForEvent(event.EventType, asset.Status)
	asset.UpdatedAt = s.clock()
	s.store.PutAsset(asset)
	s.store.AppendEvent(event)
	return asset, nil
}

func (s TraceService) GetAssetState(assetID string) (JewelryAsset, []TraceEvent, error) {
	asset, exists := s.store.GetAsset(assetID)
	if !exists {
		return JewelryAsset{}, nil, errors.New("asset not found")
	}
	return asset, s.store.Events(assetID), nil
}

func statusForEvent(eventType string, current AssetStatus) AssetStatus {
	switch eventType {
	case "CERTIFIED":
		return StatusCertified
	case "SHIPPED":
		return StatusInTransit
	case "SOLD":
		return StatusSold
	case "REVOKED":
		return StatusRevoked
	default:
		return current
	}
}
