package kinesis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	appConfig "verve/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

var kinesisClient *kinesis.Client

func LoadKinesisClient(appCnfg appConfig.Config) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appCnfg.AWS.Region))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	kinesisClient = kinesis.NewFromConfig(cfg)

	return nil
}

func SendRecord(count int) error {
	record := map[string]interface{}{
		"unique_request_count": count,
		"timestamp":            fmt.Sprintf("%d", time.Now().Unix()),
	}
	data, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %w", err)
	}

	_, err = kinesisClient.PutRecord(context.TODO(), &kinesis.PutRecordInput{
		PartitionKey: aws.String("unique-request-partition"),
		Data:         data,
	})
	if err != nil {
		return fmt.Errorf("failed to send record to Kinesis: %w", err)
	}

	log.Printf("Record sent to Kinesis: %s", data)
	return nil
}
