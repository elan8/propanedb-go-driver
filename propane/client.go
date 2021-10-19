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
	conn         *grpc.ClientConn
	dbClient     DatabaseClient
	databaseName string
}

func Connect(ctx context.Context, serverAddress string) (*Client, error) {

	client := &Client{}
	InitLog(true)
	//var opts []grpc.DialOption
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	client.conn = conn
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	//NewDatabaseClient()
	client.dbClient = NewDatabaseClient(client.conn)

	// db := &PropaneDatabase{}
	// db.Name = databaseName
	// db.DescriptorSet = descriptorSet
	// _, err = client.dbClient.CreateDatabase(ctx, db)
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }
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
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }
	//return c.dbClient.Put(ctx, entity)
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

func (c *Client) Get(ctx context.Context, id *PropaneId) (entity *PropaneEntity, err error) {
	if id.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Get(ctx, id)
}

func (c *Client) Delete(ctx context.Context, id *PropaneId) (statusOut *PropaneStatus, err error) {
	if id.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Delete(ctx, id)
}

func (c *Client) Search(ctx context.Context, input *PropaneSearch) (output *PropaneEntities, err error) {
	if input.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Search(ctx, input)
}
