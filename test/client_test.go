package test

import (
	"context"
	"io/ioutil"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/elan8/propanedb-go-driver/propane"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	// pool, err := dockertest.NewPool("")
	// if err != nil {
	// 	log.Fatalf("Could not connect to docker: %s", err)
	// }

	// // pulls an image, creates a container based on it and runs it
	// resource, err := pool.Run("jevon82/propanedb", "latest", []string{"ROOT_PASSWORD=secret"})
	// if err != nil {
	// 	log.Fatalf("Could not start resource: %s", err)
	// }

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	// if err := pool.Retry(func() error {
	// 	var err error
	// 	db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return db.Ping()
	// }); err != nil {
	// 	log.Fatalf("Could not connect to docker: %s", err)
	// }

	// code := m.Run()

	// // You can't defer this because os.Exit doesn't care for defer
	// if err := pool.Purge(resource); err != nil {
	// 	log.Fatalf("Could not purge resource: %s", err)
	// }

	// os.Exit(code)
}

func TestConnect(t *testing.T) {
	// db.Query()
	ctx := context.Background()
	//ds := descriptorpb.FileDescriptorSet{}

	b, err := ioutil.ReadFile("../pb/test.bin")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	proto.Unmarshal(b, fds)

	client, err := propane.Connect(ctx, "localhost:50051", "test", fds)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Print(client)
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
