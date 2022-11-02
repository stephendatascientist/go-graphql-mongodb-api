package graph

import "github.com/stephendatascientist/go-graphql-mongodb-api/repositories"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repositories.CustomerRepository
}
