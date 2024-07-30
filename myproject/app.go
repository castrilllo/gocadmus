package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// App struct
type App struct {
	ctx context.Context
	ec2Client *ec2.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an EC2 client
	a.ec2Client = ec2.NewFromConfig(cfg)
}

// ListRunningEC2Instances lists the running EC2 instances and returns them as a JSON object
func (a *App) ListRunningEC2Instances() ([]InstanceInfo, error) {
	input := &ec2.DescribeInstancesInput{}

	result, err := a.ec2Client.DescribeInstances(a.ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instances: %w", err)
	}

	var instances []InstanceInfo
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if instance.State.Name == types.InstanceStateNameRunning {
				instanceName := "N/A"
				for _, tag := range instance.Tags {
					if *tag.Key == "Name" {
						instanceName = *tag.Value
					}
				}
				instances = append(instances, InstanceInfo{
					ID:   *instance.InstanceId,
					Name: instanceName,
				})
			}
		}
	}

	return instances, nil
}

// InstanceInfo represents the information of an EC2 instance
type InstanceInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
