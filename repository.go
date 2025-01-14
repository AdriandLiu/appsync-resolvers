package resolvers

import (
	"fmt"
	"reflect"
)

// Repository stores all resolvers
type Repository map[string]resolver

// Add stores a new resolver
func (r Repository) Add(resolve string, handler interface{}) error {
	err := validators.run(reflect.TypeOf(handler))

	if err == nil {
		r[resolve] = resolver{handler}
	}

	return err
}

// Handle responds to the AppSync request
func (r Repository) Handle(in invocation) (interface{}, error) {
	fmt.Println(in)
	fmt.Println(in.Resolve)
	fmt.Println(in.Context)
	fmt.Println(in.payload())
	handler, found := r[in.Resolve]

	if found {
		return handler.call(in.payload())
	}

	return nil, fmt.Errorf("No resolver found: %s", in.Resolve)
}
