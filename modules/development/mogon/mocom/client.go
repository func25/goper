package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mocom only support one client, one database for now
var client *mongo.Client // default client
var db Database          // default database

type Database struct {
	*mongo.Database
}

// Connect mongodb://localhost:27017/?w=majority&retryWrites=false
func ConnectWithURI(ctx context.Context, uri string, dbName string, opts ...*options.ClientOptions) (*mongo.Client, error) {
	var err error
	client, err = mongo.Connect(ctx, append(opts, options.Client().ApplyURI(uri))...)
	if err != nil {
		return nil, err
	}
	db.Database = client.Database(dbName)
	return client, err
}

func Setup(database *mongo.Database) {
	db.Database = database
	client = db.Client()
}

func GetClient() *mongo.Client {
	return client
}

func GetDatabase() Database {
	return db
}

func CollRead(collName string) *mongo.Collection {
	return db.CollRead(collName)
}

func (db Database) CollRead(collName string) *mongo.Collection {
	return db.Collection(collName, options.Collection().SetReadPreference(readpref.Nearest()))
}

func CollWrite(collName string) *mongo.Collection {
	return db.CollWrite(collName)
}

func (db Database) CollWrite(collName string) *mongo.Collection {
	return db.Collection(collName, options.Collection().SetReadPreference(readpref.Primary()))
}
