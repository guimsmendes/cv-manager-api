package repository

import (
	"cv-manager-api/src/core/gateway"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CVRepository struct { *mgo.Database }

func (s *CVRepository) DbInsert(collection string, insert interface{}) error {
	c := s.C(collection)
	err := c.Insert(toDoc(insert))
	return err
}

func (s *CVRepository) DbFindOne(collection string, findBson bson.M, selectBson bson.M) (map[string]interface{}, error) {
	c := s.C(collection)
	getMap := make(map[string]interface{})
	err := c.Find(findBson).Select(selectBson).One(&getMap)
	return getMap, err
}

func (s *CVRepository) DbFindAll(collection string, findBson bson.M, selectBson bson.M) (map[string]interface{}, error) {
	c := s.C(collection)
	getMap := make(map[string]interface{})
	err := c.Find(findBson).Select(selectBson).All(&getMap)
	return getMap, err
}

func (s *CVRepository) DbUpdate(collection string, selector bson.M, update bson.M) error {
	c := s.C(collection)
	setBson := bson.M{}
	setBson["$set"] = update
	//
	updateError := c.Update(selector, setBson)
	//
	return updateError
}

func (s *CVRepository) DbRemoveOne(collection string, selector bson.M) error {
	c := s.C(collection)
	removeError := c.Remove(selector)
	return removeError
}

func toDoc(v interface{}) (doc *bson.DocElem) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func NewCVRepository(connection *mgo.Database) gateway.CVRepository {
	return &CVRepository {connection}
}
