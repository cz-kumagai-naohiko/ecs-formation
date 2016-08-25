package types

import "github.com/aws/aws-sdk-go/service/ecs"

type Cluster struct {
	Name     string
	Services map[string]Service
}

type Service struct {
	Name             string
	TaskDefinition   string         `yaml:"task_definition"`
	DesiredCount     int64          `yaml:"desired_count"`
	KeepDesiredCount bool           `yaml:"keep_desired_count"`
	LoadBalancers    []LoadBalancer `yaml:"load_balancers"`
	Role             string         `yaml:"role"`
}

type LoadBalancer struct {
	Name          string `yaml:"name"`
	ContainerName string `yaml:"container_name"`
	ContainerPort int64  `yaml:"container_port"`
}

type ServiceUpdatePlan struct {
	Name            string
	InstanceARNs    []*string
	CurrentServices map[string]*ecs.Service
	NewServices     map[string]*Service
}