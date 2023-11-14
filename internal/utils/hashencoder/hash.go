package hashencoder

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type HashingManager struct {
	hash hash.Hash
}

func New() *HashingManager {
	return &HashingManager{
		hash: sha256.New(),
	}
}

func (h *HashingManager) Encode(key string) string {
	src := []byte(key)
	h.hash.Write(src)
	dst := h.hash.Sum(nil)
	h.hash.Write(dst)
	dst = h.hash.Sum(nil)
	h.hash.Reset()
	return hex.EncodeToString(dst)
}
