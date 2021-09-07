package propane

import (
	"context"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Client struct {
	conn     *grpc.ClientConn
	dbClient DatabaseClient
}

func Connect(ctx context.Context, serverAddress string, databaseName string, descriptorSet *descriptorpb.FileDescriptorSet) (*Client, error) {

	client := &Client{}

	//var opts []grpc.DialOption
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	client.conn = conn
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	client.dbClient = NewDatabaseClient(client.conn)

	db := &PropaneDatabase{}
	db.Name = databaseName
	db.DescriptorSet = descriptorSet
	_, err = client.dbClient.CreateDatabase(ctx, db)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	return client, err

}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.conn.Close()
}

func (c *Client) Put(ctx context.Context, entity *PropaneEntity) (id *PropaneId, err error) {
	return c.dbClient.Put(ctx, entity)
}

func (c *Client) Get(ctx context.Context, id *PropaneId) (entity *PropaneEntity, err error) {
	return c.dbClient.Get(ctx, id)
}

func (c *Client) Delete(ctx context.Context, id *PropaneId) (status *PropaneStatus, err error) {
	return c.dbClient.Delete(ctx, id)
}
