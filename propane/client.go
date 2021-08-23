package propane

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/elan8/propanedb-go-driver/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Client struct {
	conn     *grpc.ClientConn
	dbClient pb.DatabaseClient
}

func Connect(ctx context.Context, serverAddress string, databaseName string, descriptorSet *descriptorpb.FileDescriptorSet) (*Client, error) {

	client := &Client{}

	//var opts []grpc.DialOption
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	client.conn = conn
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	client.dbClient = pb.NewDatabaseClient(client.conn)

	db := &pb.PropaneDatabase{}
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

func (c *Client) Put(ctx context.Context, entity *pb.PropaneEntity) (id *pb.PropaneId, err error) {
	return c.dbClient.Put(ctx, entity)
}

func (c *Client) Get(ctx context.Context, id *pb.PropaneId) (entity *pb.PropaneEntity, err error) {
	return c.dbClient.Get(ctx, id)
}

func (c *Client) Delete(ctx context.Context, id *pb.PropaneId) (status *pb.PropaneStatus, err error) {
	return c.dbClient.Delete(ctx, id)
}
