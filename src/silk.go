package silk

import (
	"github.com/infoforcefeed/olegdb/pkg/goleg"
)

type Silk struct {
	database goleg.Database
}

func New() *Silk {
	return &Silk{
		database: goleg.Database{},
	}
}

func (silk *Silk) NewDatabase() {

}

func (silk *Silk) OpenDatabase() {

}

func (silk *Silk) CloseDatabase() {

}
