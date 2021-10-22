package propane

import (
	"context"

	filename "github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
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
	conn             *grpc.ClientConn
	dbClient         DatabaseClient
	selectedDatabase string
}

func Connect(ctx context.Context, serverAddress string) (*Client, error) {
	client := &Client{}
	InitLog(true)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	client.conn = conn
	if err != nil {
		return nil, err
	}
	client.dbClient = NewDatabaseClient(client.conn)
	return client, err
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.conn.Close()
}

func (c *Client) SelectDatabase(ctx context.Context, databaseName string) error {
	c.selectedDatabase = databaseName

	return nil
}

func (c *Client) CreateDatabase(ctx context.Context, databaseName string, fileDescriptorSet *descriptorpb.FileDescriptorSet) (err error) {
	if databaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return err
	}
	db := &PropaneDatabase{}
	db.DatabaseName = databaseName
	db.DescriptorSet = fileDescriptorSet
	_, err = c.dbClient.CreateDatabase(ctx, db)
	return err
}

func (c *Client) Put(ctx context.Context, object interface{}) (id string, err error) {

	ctx = metadata.AppendToOutgoingContext(ctx, "database-name", c.selectedDatabase)

	entity1 := &PropaneEntity{}
	put := &PropanePut{}
	entity := object.(protoreflect.ProtoMessage)
	any, err := anypb.New(entity)
	if err != nil {
		return "", err
	}
	entity1.Data = any
	put.Entity = entity1
	//put.DatabaseName = databaseName
	propaneId, err := c.dbClient.Put(ctx, put)
	if err != nil {
		return "", err
	}

	return propaneId.Id, err
}

func (c *Client) Get(ctx context.Context, id string) (entity interface{}, err error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "database-name", c.selectedDatabase)

	propaneId := &PropaneId{}
	//propaneId.DatabaseName = databaseName
	propaneId.Id = id

	propaneEntity, err := c.dbClient.Get(ctx, propaneId)
	if err != nil {
		return nil, err
	}
	any := propaneEntity.Data
	entity, err = any.UnmarshalNew()
	if err != nil {
		return nil, err
	}
	return entity, err
}

func (c *Client) Delete(ctx context.Context, id string) (err error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "database-name", c.selectedDatabase)

	propaneId := &PropaneId{}
	//propaneId.DatabaseName = databaseName
	propaneId.Id = id

	_, err = c.dbClient.Delete(ctx, propaneId)

	return err
}

func (c *Client) Search(ctx context.Context, entityType string, query string) (output []interface{}, err error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "database-name", c.selectedDatabase)
	input := &PropaneSearch{}
	//input.DatabaseName = databaseName
	input.EntityType = entityType
	input.Query = query

	propaneEntities, err := c.dbClient.Search(ctx, input)
	if err != nil {
		return nil, err
	}

	for _, propaneEntity := range propaneEntities.Entities {
		any := propaneEntity.Data
		entity, err := any.UnmarshalNew()
		if err != nil {
			return nil, err
		}
		output = append(output, entity)
	}

	return output, nil
}
