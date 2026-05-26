package util

import "sync"

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func (s *SyncMap[K, V]) Store(key K, value V) { _ = "STUB: not implemented"; return }

func (s *SyncMap[K, V]) CompareAndDelete(key K, value V) { _ = "STUB: not implemented"; return }

func (s *SyncMap[K, V]) Load(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (s *SyncMap[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (s *SyncMap[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) { _ = "STUB: not implemented"; return }
