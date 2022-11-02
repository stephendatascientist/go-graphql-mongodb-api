package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/stephendatascientist/go-graphql-mongodb-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	Client *mongo.Client
}

func (r *CustomerRepository) CreateCustomer(customer model.NewCustomer) *model.Customer {
	customerColl := r.Client.Database("graphql-mongodb-api").Collection("customer")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := customerColl.InsertOne(ctx, bson.M{"name": customer.Name, "email": customer.Email, "mobile": customer.Email})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnCustomer := model.Customer{ID: insertedID, Name: customer.Name}

	return &returnCustomer
}

func (r *CustomerRepository) GetCustomers() []*model.Customer {
	customerColl := r.Client.Database("graphql-mongodb-api").Collection("customer")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := customerColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var customers []*model.Customer
	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}

		customer := model.Customer{ID: (sus[0].String()), Name: (sus[1].String())}

		customers = append(customers, &customer)
	}

	return customers
}

func (r *CustomerRepository) GetCustomer(id string) *model.Customer {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	customerColl := r.Client.Database("graphql-mongodb-api").Collection("customer")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := customerColl.FindOne(ctx, bson.M{"_id": ObjectID})

	customer := model.Customer{ID: id}

	res.Decode(&customer)

	return &customer
}
