package api

import "github.com/graphql-go/graphql"

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createSong": &graphql.Field{
			Type: songType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"album": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"duration": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var song Song
				song.ID = params.Args["id"].(string)
				song.Album = params.Args["album"].(string)
				song.Title = params.Args["title"].(string)
				song.Duration = params.Args["duration"].(string)
				songs = append(songs, song)
				return song, nil
			},
		},
	},
})
