package propane

import (
	"context"

	"github.com/elan8/propanedb-go-driver/pb"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	conn     *grpc.ClientConn
	dbClient pb.DatabaseClient
}

func Connect(ctx context.Context, serverAddress string) (*Client, error) {

	client := &Client{}

	//var opts []grpc.DialOption
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	client.conn = conn
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	//pb.NewDatabaseClient()
	client.dbClient = pb.NewDatabaseClient(client.conn)

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
func (c *Client) CreateDatabase(ctx context.Context, db *pb.PropaneDatabase) (statusOut *pb.PropaneStatus, err error) {

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

func (c *Client) Put(ctx context.Context, put *pb.PropanePut) (id *pb.PropaneId, err error) {
	if put.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Put(ctx, put)
}

func (c *Client) Get(ctx context.Context, id *pb.PropaneId) (entity *pb.PropaneEntity, err error) {
	if id.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Get(ctx, id)
}

func (c *Client) Delete(ctx context.Context, id *pb.PropaneId) (statusOut *pb.PropaneStatus, err error) {
	if id.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Delete(ctx, id)
}

func (c *Client) Search(ctx context.Context, input *pb.PropaneSearch) (output *pb.PropaneEntities, err error) {
	if input.DatabaseName == "" {
		err := status.Error(codes.NotFound, "Database name is empty")
		return nil, err
	}
	return c.dbClient.Search(ctx, input)
}
