package silk

import (
	"github.com/google/uuid"
)

type RelationalNode struct { //anything not public will not get encoded without a custom en/decoder (which defeats the purpose of using gob)
	Chunks    []uuid.UUID
	Key       string //UWUID
	Name      string
	Links     []string
	Backlinks []string
	Tags      []string
}

func NewRelationalNode(value string) *RelationalNode {
	//maybe take a list for tags and links?
	key, uuids := KeyFromValue(value)
	return &RelationalNode{
		Chunks:    uuids,
		Key:       key,
		Name:      value,
		Links:     nil,
		Backlinks: nil,
		Tags:      nil,
	}
}

// returns a big uuid-like string and an array of component uuids from a seed string
// if there is somehow a collision please contact me so I can quit as a developer
func KeyFromValue(value string) (string, []uuid.UUID) {
	//grab up to 5 chunks of up to 64 bytes, left to right
	var chunkBytes [][]byte
	bytes := []byte(value)
	bytes_len := len(bytes)
chunk:
	for chunk_no := 0; chunk_no < 5; chunk_no++ {
		var chunk []byte
		for single_chunk_no := 0; single_chunk_no < 64; single_chunk_no++ {
			index := (chunk_no * 64) + single_chunk_no
			if bytes_len > index {
				chunk = append(chunk, bytes[index])
			} else {
				chunkBytes = append(chunkBytes, chunk)
				break chunk
			}
		}
		chunkBytes = append(chunkBytes, chunk)
	}
	//md5 each chunk
	var chunkBytesHashed []uuid.UUID
	for _, chunk := range chunkBytes {
		chunkBytesHashed = append(chunkBytesHashed, uuid.NewMD5(uuid.Nil, chunk))
	}
	// an entire uuid just for emergencies
	chunkBytesHashed = append(chunkBytesHashed, uuid.NewSHA1(uuid.NameSpaceURL, bytes))
	//combine UUIDs for string
	var UWUID string
	for _, hash := range chunkBytesHashed {
		if len(UWUID) != 0 {
			UWUID = UWUID + "-" + hash.String()
		} else {
			UWUID = UWUID + hash.String()
		}
	}
	return UWUID, chunkBytesHashed
}
