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

func (s *CVRepository) DbFindOne(collection string, findBson bson.M) (bson.M, error) {
	c := s.C(collection)
	var doc bson.M
	err := c.Find(findBson).One(&doc)
	return doc, err
}

func (s *CVRepository) DbFindAll(collection string, findBson bson.M) ([]bson.M, error) {
	c := s.C(collection)
	var docList []bson.M
	err := c.Find(findBson).All(&docList)
	return docList, err
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

func (s *CVRepository) DbRemoveOne(collection string, findBson bson.M) error {
	c := s.C(collection)
	return c.Remove(findBson)
}

func toDoc(v interface{}) (doc *bson.M) {
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
