package github_test

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/elan8/propanedb-go-driver/propane"
	"github.com/ory/dockertest"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
)

var port string

var docker = flag.Bool("docker", false, "Enable docker test")

func TestMain(m *testing.M) {

	flag.Parse()
	fmt.Println("dockertestEnabled:", *docker)

	if *docker {
		log.Print("Use dockertest")
		// uses a sensible default on windows (tcp/http) and linux/osx (socket)
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}

		// // pulls an image, creates a container based on it and runs it
		resource, err := pool.Run("ghcr.io/elan8/propanedb", "latest", nil)

		log.Print("Start container")
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}

		//resource.GetPort("50051/tcp")

		time.Sleep(4 * time.Second)

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		if err := pool.Retry(func() error {
			ctx := context.Background()
			port = resource.GetPort("50051/tcp")
			log.Print("Port=" + port)
			_, err := propane.Connect(ctx, "localhost:"+port)
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}

		code := m.Run()

		// // You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
		os.Exit(code)
	} else {
		log.Print("Use external PropaneDB instance")
		port = "50051"
		code := m.Run()
		os.Exit(code)
	}

}
func TestConnect(t *testing.T) {

	databaseName := "test"
	ctx := context.Background()

	b, err := ioutil.ReadFile("../propane/test.bin")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	proto.Unmarshal(b, fds)

	client, err := propane.Connect(ctx, "localhost:"+port)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	db := &propane.PropaneDatabase{}
	db.DatabaseName = databaseName
	db.DescriptorSet = fds

	_, err = client.CreateDatabase(ctx, db)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	//add item 1
	item1 := &propane.TestEntity{}
	item1.Description = "Test 1"

	// entity1 := &propane.PropaneEntity{}
	// put := &propane.PropanePut{}
	// any, err := anypb.New(item1)
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }
	// entity1.Data = any
	// put.Entity = entity1
	// put.DatabaseName = databaseName
	id1, err := client.Put(ctx, databaseName, item1)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Print("Id1=" + id1)

	//add item 2
	item2 := &propane.TestEntity{}
	item2.Description = "Test 2"

	// entity2 := &propane.PropaneEntity{}
	// put2 := &propane.PropanePut{}
	// any2, err := anypb.New(item2)
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }
	// entity2.Data = any2
	// put2.Entity = entity2
	// put2.DatabaseName = databaseName
	id2, err := client.Put(ctx, databaseName, item2)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Print("Id2=" + id2)

	//get item 1

	propaneId := &propane.PropaneId{}
	propaneId.DatabaseName = databaseName
	propaneId.Id = id1
	entity3, err := client.Get(ctx, propaneId)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("Entity 1: %s", entity3.String())

	any, err := anypb.New(item1)
	m := new(propane.TestEntity)
	if err := any.UnmarshalTo(m); err != nil {
		log.Fatalf("Error: %s", err)
		t.Errorf("Cannot unmarshal to TestEntity")
	}

	if m.Description != "Test 1" {
		t.Errorf("expected 'Test 1', got '%s'", m.Description)
	}

	//get all items
	input := propane.PropaneSearch{}
	input.DatabaseName = databaseName
	input.EntityType = "test.TestEntity"
	input.Query = "*"

	entities, err := client.Search(ctx, &input)

	if len(entities.Entities) != 2 {
		t.Errorf("expected length = 2, got '%d'", len(entities.Entities))
	}

	log.Print("Get all entities")
	for _, myEntity := range entities.GetEntities() {
		log.Printf("Entity: %s", myEntity.String())
	}

	//delete item 1

	status, err := client.Delete(ctx, propaneId)
	log.Printf("Delete Status: %s", status.GetStatusMessage())
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}