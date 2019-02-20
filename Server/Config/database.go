package Config

import (
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

type MonggodbConnector struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
	url      string
}

func (m *MonggodbConnector) Connect() *mgo.Database {
	m.Host = viper.GetString("database_host")
	m.User = viper.GetString("database_user")
	m.Password = viper.GetString("database_user_password")
	m.Database = viper.GetString("database_name")
	m.Port = viper.GetInt("database_port")
	if m.User == "" || m.Password == "" {
		m.url = "mongodb://" + m.Host + ":" + strconv.Itoa(m.Port)
	} else {
		m.url = "mongodb://" + m.User + ":" + m.Password + "@" + m.Host + ":" + strconv.Itoa(m.Port)
	}
	session, err := mgo.Dial(m.url)
	if err != nil {
		panic(err)
	}
	Db := session.DB(viper.GetString("database_name"))
	return Db
}
