package main

type Config struct {
	Elasticsearch Elasticsearch
	AWS           AWS
}

type Elasticsearch struct {
	Endpoint string
}

type AWS struct {
	Region  string
	Service string
	RoleARN string
}

type User struct {
	Password     string
	BackendRoles []string `yaml:"backend_roles" json:"backend_roles"`
}
