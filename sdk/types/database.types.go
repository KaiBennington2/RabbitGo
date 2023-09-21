package types

import "go.mongodb.org/mongo-driver/mongo/options"

type DatabaseOpts struct {
	Name string
	Opts []*options.DatabaseOptions
}
