package petsdb

import (
	
	"context"
	"fmt"
	"log"
	"os"
    "time"

    "cloud.google.com/go/datastore"
  //  "github.com/google/uuid"
)

var projectID string

// Pet model stored in Datastore
type Pet struct {
	Added   time.Time `datastore:"added"`
	Caption string    `datastore:"caption"`
	Email   string    `datastore:"email"`
	Image   string    `datastore:"image"`
	Likes   int       `datastore:"likes"`
	Owner   string    `datastore:"owner"`
	Petname string    `datastore:"petname"`
	Name    string     // The ID used in the Datastore.
}


func connectorDatastore()(context.Context, *datastore.Client) {
	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_CLOUD_PROJECT"`)
	}
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
	}
	return ctx, client
}

// GetPets Returns all pets from datastore
func GetPets() ([]Pet, error) {
	var pets []Pet
	ctx, client := connectorDatastore()

	// Create a query to fetch all Pet entities
	query := datastore.NewQuery("Pet") //.Order("-likes")
	keys, err := client.GetAll(ctx, query, &pets)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set the id field on each Task from the corresponding key.
	for i, key := range keys {
		pets[i].Name = key.Name
	}

	client.Close()
	return pets, nil
}

func CreatePet(pet Pet){
	ctx, client := connectorDatastore()
	key := datastore.IncompleteKey("Pet", nil)
	//newID := uuid.New().String()
	//pet.Name = datastore.NameKey("Pet", "StringId", newID)
	_, err := client.Put(ctx, key, &pet)
	if err != nil {
		fmt.Println(err)
	}
	client.Close()
}
