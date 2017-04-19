package helper

import (
	"errors"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoGetNextID increment the ID (by one) of the element with an _id of objectID in the giving database.collection
func MongoGetNextID(session *mgo.Session, db string, collection string, objectID string) (string, error) {
	var result bson.M
	s := session.Copy()
	defer s.Close()

	c := s.DB(db).C(collection)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: true,
	}

	_, err := c.Find(bson.M{"_id": objectID}).Apply(change, &result)
	if err != nil {
		return "", errors.New("Can't create new ID")
	}

	return strconv.FormatFloat(result["counter"].(float64), 'f', -1, 64), nil
}
