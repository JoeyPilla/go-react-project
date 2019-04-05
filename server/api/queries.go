package api

import (
	"strings"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"songs": &graphql.Field{
			Type: graphql.NewList(songType),
			Args: graphql.FieldConfigArgument{
				"album": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				album := params.Args["album"].(string)
				filtered := Filter(songs, func(v Song) bool {
					return strings.Contains(v.Album, album)
				})
				return filtered, nil
			},
		},
		"album": &graphql.Field{
			Type: albumType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				for _, album := range albums {
					if album.ID == id {
						return album, nil
					}
				}
				return nil, nil
			},
		},
	},
})
