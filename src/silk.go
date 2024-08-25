package silk

import (
	"bytes"
	"encoding/gob"
	"os"

	"github.com/infoforcefeed/olegdb/pkg/goleg"
)

// declare magic feature choices
var (
	FEATS = Feats{
		//Append Only
		AppendOnly: 1 << 0,
		//LZ4
		Lz4: 1 << 2,
		//Splay Tree
		SplayTree: 1 << 1,
		//AOL FFLUSH
		AolFflush: 1 << 3,
	}
)

type Feats struct {
	AppendOnly int
	Lz4        int
	SplayTree  int
	AolFflush  int
}

// onto the main struct
type Silk struct { //wraps a single database, does not delete
	database goleg.Database
	open     bool
	feats    Feats
}

func New() *Silk {
	return &Silk{
		database: goleg.Database{},
		open:     false,
		feats:    FEATS,
	}
}

// give me a random database here and switch me to it
func (silk *Silk) NewDatabase(basedir string, feats int) (string, error) {
	silk.CloseDatabase()
	randomdir, err := os.CreateTemp(basedir, "silk")
	if err != nil {
		return "", err
	} else {
		err = silk.OpenDatabase(basedir, randomdir.Name(), feats)
		if err != nil {
			return "", err
		} else {
			return randomdir.Name(), err
		}
	}
}

// open a database in a directory I know exists (currently creates a database if it doesn't exist, as long as the dir exists)
func (silk *Silk) OpenDatabase(dir string, name string, feats int) error {
	database, err := goleg.Open(dir, name, feats)
	if err != nil {
		return err
	} else {
		silk.CloseDatabase()
		silk.database = database
		silk.open = true
		return err
	}
}

// close a database, used internally and externally
func (silk *Silk) CloseDatabase() {
	if silk.open {
		silk.database.Close()
		silk.open = false
	}
}

// find me all valid databases in a directory
func (silk *Silk) FindDiskDatabase(dir string) {

}

func (silk *Silk) PushNode(node RelationalNode) {
	//this is a lot of allocations, need to figure out how to use existing encoders.
	// Maybe with channel and loop? #TODO
	tmpBuffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&tmpBuffer)
	//encode
	encoder.Encode(node)
	silk.database.Jar(node.Key, tmpBuffer.Bytes())
}

func (silk *Silk) PullNode(uuid string) RelationalNode {
	//this is a lot of allocations, need to figure out how to use existing decoders.
	// Maybe with channel and loop? #TODO
	tmpBuffer := bytes.NewBuffer(silk.database.Unjar(uuid))
	decoder := gob.NewDecoder(tmpBuffer)
	//decode
	var node RelationalNode
	decoder.Decode(&node)
	return node
}
