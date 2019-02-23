package Content

import (
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Config"
	"github.com/globalsign/mgo"
)

const (
	collectionName = "Artikel"
)

type DataContent struct {
	Title      string
	Author     string
	Categories string
	Content    string
}

func (d *DataContent) Models() *mgo.Collection {
	var Database Config.MonggodbConnector
	connect := Database.Connect()
	collection := connect.C(collectionName)
	return collection
}
