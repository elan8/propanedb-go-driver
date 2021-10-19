package propane

import (
	"context"

	filename "github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

func InitLog(debugMode bool) {
	if debugMode {
		filenameHook := filename.NewHook()
		filenameHook.Field = "line"
		log.AddHook(filenameHook)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

type Client struct {
	conn     *grpc.ClientConn
	dbClient DatabaseClient
}

func Connect(ctx context.Context, serverAddress string) (*Client, error) {

	client := &Client{}
	InitLog(true)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	client.conn = conn
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	client.dbClient = NewDatabaseClient(client.conn)
	return client, err

}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.conn.Close()
}

// rpc CreateDatabase(PropaneDatabase) returns (PropaneStatus) {}
func (c *Client) CreateDatabase(ctx context.Context, db *PropaneDatabase) (statusOut *PropaneStatus, err error) {
	if db.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.CreateDatabase(ctx, db)
}

func (c *Client) Put(ctx context.Context, databaseName string, object interface{}) (id string, err error) {
	if databaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return "", err
	}

	entity1 := &PropaneEntity{}
	put := &PropanePut{}
	entity := object.(protoreflect.ProtoMessage)
	any, err := anypb.New(entity)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	entity1.Data = any
	put.Entity = entity1
	put.DatabaseName = databaseName
	propaneId, err := c.dbClient.Put(ctx, put)

	return propaneId.Id, err
}

func (c *Client) Get(ctx context.Context, databaseName string, id string) (entity interface{}, err error) {
	if databaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}

	propaneId := &PropaneId{}
	propaneId.DatabaseName = databaseName
	propaneId.Id = id

	propaneEntity, err := c.dbClient.Get(ctx, propaneId)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	//log.Printf("Entity 1: %s", entity3.String())

	any := propaneEntity.Data
	entity, err = any.UnmarshalNew()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	//any, err := anypb.New(item1)
	//m := new(propane.TestEntity)
	// if err := any.UnmarshalTo(m); err != nil {
	// 	log.Fatalf("Error: %s", err)
	// 	t.Errorf("Cannot unmarshal to TestEntity")
	// }

	return entity, err
}

func (c *Client) Delete(ctx context.Context, databaseName string, id string) (err error) {
	if databaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return err
	}

	propaneId := &PropaneId{}
	propaneId.DatabaseName = databaseName
	propaneId.Id = id

	_, err = c.dbClient.Delete(ctx, propaneId)

	return err
}

func (c *Client) Search(ctx context.Context, databaseName string, entityType string, query string) (output []interface{}, err error) {
	if databaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	input := &PropaneSearch{}
	input.DatabaseName = databaseName
	input.EntityType = entityType
	input.Query = query

	propaneEntities, err := c.dbClient.Search(ctx, input)

	for _, propaneEntity := range propaneEntities.Entities {
		any := propaneEntity.Data
		entity, err := any.UnmarshalNew()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		output = append(output, entity)
	}

	return
}
