package engine

const (
	TagPyro = 1 << 1
	TagDirt = 1 << 2
	TagAir  = 1 << 3
	TagCryo = 1 << 4
)

//
//import (
//	"math"
//	"strings"
//	"sync"
//)
//
//var globalTagLock sync.RWMutex
//var globalTagLatest = 0
//var globalTagMap = make(map[string]Tag)
//
//// Tag is the index of a tag in a Tags map.
//type Tag struct {
//	slot        uint16
//	indexInSlot uint8
//	indexMask   uint8
//}
//
//// TagByName gets a tag by its name.
//// If no tag with that name exists, one will be created.
//func TagByName(name string) Tag {
//	cleanName := strings.ToLower(name)
//
//	// Try to get an existing tag.
//	globalTagLock.RLock()
//	tag, exists := globalTagMap[cleanName]
//	globalTagLock.RUnlock()
//	if exists {
//		return tag
//	}
//
//	// Require the lock as exclusive.
//	globalTagLock.Lock()
//	defer globalTagLock.Unlock()
//
//	// If it exists now, some other thread made it.
//	tag, exists = globalTagMap[cleanName]
//	if exists {
//		return tag
//	}
//
//	// If it doesn't, we create it.
//	tag = Tag{
//		slot:        uint16(globalTagLatest / 8),
//		indexInSlot: uint8(globalTagLatest % 8),
//		indexMask:   uint8(math.Pow(2, float64(globalTagLatest%8))),
//	}
//
//	globalTagMap[cleanName] = tag
//	globalTagLatest += 1
//	return tag
//}
//
//// TagSet is a set of tags.
//type TagSet struct {
//	tags []uint8
//}
//
//func (tags *TagSet) Has(tag Tag) bool {
//	if int(tag.slot) > len(tags.tags) {
//		return false
//	}
//
//	return (tags.tags[tag.slot] & tag.indexMask) != 0
//}
//
//func (tags *TagSet) Set(tag Tag, value bool) {
//	for int(tag.slot) < len(tags.tags) {
//
//	}
//
//}
//func (receiver) name() {
//
//}
