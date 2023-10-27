// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Asset struct {
	ID string `json:"id"`
}

type Config struct {
	Scope           []interface{} `json:"scope,omitempty"`
	Ports           []*int        `json:"ports,omitempty"`
	Blacklist       []*string     `json:"blacklist,omitempty"`
	Domains         []*string     `json:"domains,omitempty"`
	Resolvers       []*string     `json:"resolvers,omitempty"`
	Ips             []*string     `json:"ips,omitempty"`
	Cirds           []interface{} `json:"cirds,omitempty"`
	Transformations []interface{} `json:"transformations,omitempty"`
	Database        []interface{} `json:"database,omitempty"`
	BruteForce      *bool         `json:"bruteForce,omitempty"`
	Alterations     *bool         `json:"alterations,omitempty"`
}

type CreateAssetInput struct {
	SessionToken string      `json:"sessionToken"`
	AssetName    *string     `json:"assetName,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

type CreateSessionInput struct {
	Config *Config `json:"config"`
}

type CreateSessionJSONInput struct {
	Config string `json:"config"`
}

type Session struct {
	Token string `json:"token"`
}
