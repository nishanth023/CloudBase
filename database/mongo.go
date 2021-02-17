package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func new(url string) (*Db, error) {
	db := &Db{}
	err := db.init(url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Db struct {
	url    string
	client *mongo.Client
	cxt    context.Context
	cancel context.CancelFunc
}

func (db *Db) close() {
	defer db.cancel()
	defer func() {
		if err = db.client.Disconnect(db.cxt); err != nil {
			panic(err)
		}
	}()
}

func (db *Db) init(url string) (err error) {
	db.url = url
	db.cxt, db.cancel = context.WithTimeout(context.Background(), 10*time.Second)
	db.client, err = mongo.Connect(db.cxt, options.Client().ApplyURI(os.Getenv(db.url)))
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) InsertOne(dataBase, col string, doc interface{}) (err error) {
	collection := db.client.Database(dataBase).Collection(col)
	_, err = collection.InsertOne(db.cxt, doc)
	return err
}

func (db *Db) InsertMany(dataBase, col string, docs []interface{}) (err error) {
	collection := db.client.Database(dataBase).Collection(col)
	_, err = collection.InsertMany(db.cxt, docs)
	return err
}

func (db *Db) Query(dataBase, col string, query, field interface{}) (result interface{}, err error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err = collection.Find(db.cxt, query, options.Find().SetProjection(field))
	return
}

func (db *Db) UpdateOne(dataBase, col string, filter, update interface{}) (err error) {
	collection := db.client.Database(dataBase).Collection(col)
	_, err = collection.UpdateOne(db.cxt, filter, update)
	return nil
}

func (db *Db) UpdateMany(dataBase, col string, filter, update interface{}) (err error) {
	collection := db.client.Database(dataBase).Collection(col)
	_, err = collection.UpdateMany(db.cxt, filter, update)
	return nil
}

func (db *Db) DeleteOne(dataBase, col string, query interface{}) (err error) {
	collection := db.client.Database(dataBase).Collection(col)
	_, err = collection.DeleteOne(db.cxt, query)
	return
}

func (db *Db) DeleteMany(dataBase, col string, query interface{}) (err error) {
	collection := db.client.Database(dataBase).Collection(col)
	_, err = collection.DeleteMany(db.cxt, query)
	return
}

func (db *Db) DorpDb(dataBase string) (err error) {
	err = db.client.Database(dataBase).Drop(db.cxt)
	return err
}

func (db *Db) DorpCol(dataBase, col string) (err error) {
	err = db.client.Database(dataBase).Collection(col).Drop(db.cxt)
	return err
}
