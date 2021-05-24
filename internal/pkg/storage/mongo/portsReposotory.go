package mongo

import (
	"context"
	"fmt"
	"log"
	"ports/internal/pkg/pb"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	DSN string `envconfig:"MONGO_DSN"`
}

// FIXME all mongo db initialization should be extracted to other package
var defaultConfig Config

func init() {
	err := envconfig.Process("ports-mongo-storage", &defaultConfig)
	if err != nil {
		log.Fatalf("Can't configure app")
	}

}

func NewPortsRepository() (repo PortsRepository, err error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return repo, err
	}

	return PortsRepository{
		ctx:  ctx,
		Coll: client.Database("ports").Collection("ports"),
	}, nil
}

type PortsRepository struct {
	ctx  context.Context
	Coll *mongo.Collection
}

func (r *PortsRepository) Get(code string) (port *pb.Port, found bool) {
	err := r.Coll.FindOne(r.ctx, bson.M{"code": code}).Decode(&port)
	return port, err == nil || err != mongo.ErrNoDocuments
}
func (r *PortsRepository) List(codes ...string) (ports []*pb.Port, err error) {
	return r.fetch(bson.M{"code": bson.M{"$in": codes}})
}

func (r *PortsRepository) fetch(query bson.M) (ports []*pb.Port, err error) {
	cur, err := r.Coll.Find(r.ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(r.ctx)
	for cur.Next(r.ctx) {
		var port pb.Port
		err := cur.Decode(&port)
		if err != nil {
			return ports, fmt.Errorf("getting ports decoding error: %w", err)
		}
		ports = append(ports, &port)
	}
	if err := cur.Err(); err != nil {
		return ports, fmt.Errorf("getting ports error: %w", err)
	}
	return
}

func (r *PortsRepository) Delete(code string) error {
	_, err := r.Coll.DeleteOne(r.ctx, bson.M{"code": code})
	return err
}
func (r *PortsRepository) Save(port *pb.Port) error {
	_, err := r.Coll.ReplaceOne(r.ctx, bson.M{"code": port.Code}, port, options.Replace().SetUpsert(true))
	// _, err := r.Coll.UpdateOne(r.ctx, bson.M{"code": port.Code}, port, options.Update().SetUpsert(true))
	return err
}
