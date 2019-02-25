package Content

import (
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Config"
	"github.com/globalsign/mgo"
)

const (
	collectionName = "Content"
)

type DataContent struct {
	Title      string `json:"title" validate:"required"`
	Author     string `json:"author" validate:"required"`
	Categories string `json:"categories" validate:"required"`
	Content    string `json:"content"`
}

func (d *DataContent) Models() *mgo.Collection {
	var Database Config.MonggodbConnector
	connect := Database.Connect()
	collection := connect.C(collectionName)
	return collection
}
