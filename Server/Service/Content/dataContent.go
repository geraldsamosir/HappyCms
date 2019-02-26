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

type QueryUrlForUpdate struct {
	_Id string `schema:"id,required"`
}

type QueryUrlForFilter struct {
	_Id    string `schema:"id,"`
	Title  string `schema:"title"`
	Author string `schema:"author"`
}

func (d *DataContent) Models() *mgo.Collection {
	var Database Config.MonggodbConnector
	connect := Database.Connect()
	collection := connect.C(collectionName)
	return collection
}
