package gg

import "github.com/graphql-go/graphql"

func IntArg(name string) graphql.FieldConfigArgument  {
	return graphql.FieldConfigArgument{name:&graphql.ArgumentConfig{Type:graphql.Int}}
}
func StringArg(name string )  graphql.FieldConfigArgument  {
	return graphql.FieldConfigArgument{name:&graphql.ArgumentConfig{Type:graphql.Int}}
}