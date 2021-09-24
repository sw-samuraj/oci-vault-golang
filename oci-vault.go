package main

import (
	"context"
	"github.com/oracle/oci-go-sdk/v46/common"
	"github.com/oracle/oci-go-sdk/v46/secrets"
	log "github.com/sirupsen/logrus"
)

const (
	secretId = "ocid1.vaultsecret.oc1.phx.amaaaaaayrywvyyatfzgkrnxmdr2crt5npuq66yrryc6qviapidlqsyeinza"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	logger := funcLog("main")
	logger.Info("starting the secret service example")

	client := getSecretsClient()

	listSecretVersions(client)
	getSecret(client)
}

func listSecretVersions(client secrets.SecretsClient) {
	logger := funcLog("listSecretVersions")

	request := secrets.ListSecretBundleVersionsRequest{
		OpcRequestId: common.String("42-get-my-secret"),
		SecretId:     common.String(secretId),
	}

	logger.Info("calling the secret service...")
	response, err := client.ListSecretBundleVersions(context.Background(), request)
	if err != nil {
		logger.Fatalf("can't get a response from the secret service: %s", err)
	}

	logger.Info(response.Items)
}

func getSecret(client secrets.SecretsClient) {
	logger := funcLog("getSecretBundle")

	request := secrets.GetSecretBundleRequest{
		OpcRequestId: common.String("42-get-my-secret"),
		SecretId:     common.String(secretId),
		Stage:        secrets.GetSecretBundleStageLatest,
	}

	logger.Info("calling the secret service...")
	response, err := client.GetSecretBundle(context.Background(), request)
	if err != nil {
		logger.Fatalf("can't get a response from the secret service: %s", err)
	}

	logger.Info(response.SecretBundleContent)
}

func getSecretsClient() secrets.SecretsClient {
	logger := funcLog("getSecretsClient")
	client, err := secrets.NewSecretsClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		logger.Fatalf("can't get a secrets client: %s", err)
	}
	logger.Info("default secret client has been obtained")
	return client
}

func funcLog(f string) *log.Entry {
	return log.WithFields(log.Fields{
		"func": f,
	})
}
