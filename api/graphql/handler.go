package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

func graphqlSchema() graphql.Schema {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				return "Hello " + name, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)
	return schema
}

func StartGraphQLServer() {
	schema := graphqlSchema()
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
