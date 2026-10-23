package domain

import (
	"testing"
	"time"
)

func TestRegisterAndAppendTraceEvent(t *testing.T) {
	service := NewTraceService(NewMemoryStore(), func() time.Time { return time.Unix(0, 0).UTC() })

	asset, err := service.RegisterAsset("asset-1", "hash-1")
	if err != nil {
		t.Fatalf("register asset: %v", err)
	}
	if asset.Status != StatusRegistered {
		t.Fatalf("status = %s, want %s", asset.Status, StatusRegistered)
	}

	updated, err := service.AppendTraceEvent(TraceEvent{EventID: "evt-1", AssetID: "asset-1", EventType: "CERTIFIED", ActorOrgID: "cert-org", PayloadHash: "hash-2", OccurredAt: time.Unix(1, 0).UTC()})
	if err != nil {
		t.Fatalf("append event: %v", err)
	}
	if updated.Status != StatusCertified {
		t.Fatalf("status = %s, want %s", updated.Status, StatusCertified)
	}

	_, events, err := service.GetAssetState("asset-1")
	if err != nil {
		t.Fatalf("get state: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("events = %d, want 1", len(events))
	}
}

func TestAppendRejectsUnknownAsset(t *testing.T) {
	service := NewTraceService(NewMemoryStore(), nil)
	_, err := service.AppendTraceEvent(TraceEvent{EventID: "evt-1", AssetID: "missing", EventType: "SOLD", ActorOrgID: "dealer", PayloadHash: "hash", OccurredAt: time.Now()})
	if err == nil {
		t.Fatal("expected error")
	}
}
