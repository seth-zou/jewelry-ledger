package domain

type MemoryStore struct {
	assets map[string]JewelryAsset
	events map[string][]TraceEvent
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{assets: map[string]JewelryAsset{}, events: map[string][]TraceEvent{}}
}

func (s *MemoryStore) GetAsset(assetID string) (JewelryAsset, bool) {
	asset, ok := s.assets[assetID]
	return asset, ok
}

func (s *MemoryStore) PutAsset(asset JewelryAsset) {
	s.assets[asset.AssetID] = asset
}

func (s *MemoryStore) AppendEvent(event TraceEvent) {
	s.events[event.AssetID] = append(s.events[event.AssetID], event)
}

func (s *MemoryStore) Events(assetID string) []TraceEvent {
	return append([]TraceEvent(nil), s.events[assetID]...)
}
