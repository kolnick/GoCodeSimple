package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

// go get go.mongodb.org/mongo-driver/mongo
type Persons struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string
	Age  int
}

var client *mongo.Client

func initCloudMongoDB() error {
	// 设置连接选项
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://kolnick:122251238@atlascluster.vzx0lch.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	// 检查连接
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")

	return nil
}

func initLocalMongoDB() {
	// 设置连接选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, _ = mongo.Connect(context.Background(), clientOptions)
}
func closeMongoDb() {
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}

func TestGetCollection(t *testing.T) {
	client.Database("db").Collection("persons")
	defer closeMongoDb()
}

func TestInsertCollection(t *testing.T) {
	initLocalMongoDB()
	defer closeMongoDb()
	collection := client.Database("db").Collection("persons")

	for i := 0; i < 5; i++ {
		persons := Persons{Name: "kolnick", Age: 20 + i}
		_, err := collection.InsertOne(context.Background(), persons)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBatchInset(t *testing.T) {
	initLocalMongoDB()
	defer closeMongoDb()
	collection := client.Database("db").Collection("persons")

	var documents []interface{}
	for i := 0; i < 5; i++ {
		documents = append(documents, Persons{Name: fmt.Sprintf("kolnick %d", i+1)})
	}
	result, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted documents with IDs:", result.InsertedIDs)
}

func TestFindCollection(t *testing.T) {
	initLocalMongoDB()
	defer closeMongoDb()
	collection := client.Database("db").Collection("persons")
	filter := bson.D{{"age", bson.D{{"$gte", 20}, {"$lte", 23}}}}
	cur, err := collection.Find(context.Background(), filter)
	forEachCollection(err, cur)
}

func forEachCollection(err error, cur *mongo.Cursor) {
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var list []*Persons
	err = cur.All(context.Background(), &list)
	if err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.Background())

	for _, one := range list {
		fmt.Printf("_id:%s name:%s, age:%d\n", one.Id, one.Name, one.Age)
	}
}

func TestUpdateCollection(t *testing.T) {
	initLocalMongoDB()
	defer closeMongoDb()
	collection := client.Database("db").Collection("persons")
	filter := bson.D{{"name", bson.D{{"$eq", "kolnick5"}}}}
	update := bson.D{{"$set", bson.D{{"age", 33}}}}
	collection.UpdateMany(context.Background(), filter, update)
}

func TestRemoveCollection(t *testing.T) {
	initLocalMongoDB()
	defer closeMongoDb()
	collection := client.Database("db").Collection("persons")
	filter := bson.D{{"name", bson.D{{"$eq", "kolnick4"}}}}
	collection.DeleteMany(context.Background(), filter)
}

func TestLike(t *testing.T) {
	initLocalMongoDB()
	defer closeMongoDb()
	collection := client.Database("db").Collection("persons")
	filter := bson.M{"name": bson.M{"$regex": "^kolnick"}}
	cur, err := collection.Find(context.Background(), filter)
	forEachCollection(err, cur)
}

func TestReplaceOne(t *testing.T) {

}
