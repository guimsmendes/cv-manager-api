package gateway

import "gopkg.in/mgo.v2/bson"

type CVRepository interface {
	DbInsert(collection string, insert interface{}) error
	DbFindOne(collection string, findBson bson.M, selectBson bson.M) (map[string]interface{},error)
	DbFindAll(collection string, findBson bson.M, selectBson bson.M) (map[string]interface{},error)
	DbUpdate(collection string, selector bson.M, update bson.M) error
	DbRemoveOne(collection string, selector bson.M)  error
}