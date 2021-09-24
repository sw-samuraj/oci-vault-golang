package secrets

import (
	"context"
	l "cz.sw-samuraj/oci-vault/logging"
	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/secrets"
)

func GetSecretsClient() secrets.SecretsClient {
	log := l.FuncLog("GetSecretsClient")
	client, err := secrets.NewSecretsClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatalf("can't get a secrets client: %s", err)
	}
	log.Info("secrets client has been obtained")
	return client
}

func ListSecretVersions(client secrets.SecretsClient, secretId string) {
	log := l.FuncLog("ListSecretVersions")

	request := secrets.ListSecretBundleVersionsRequest{
		OpcRequestId: common.String("42-get-my-secret"),
		SecretId:     common.String(secretId),
	}

	log.Info("calling the secret service...")
	response, err := client.ListSecretBundleVersions(context.Background(), request)
	if err != nil {
		log.Fatalf("can't get a response from the secret service: %s", err)
	}

	log.Info(response.Items)
}

func GetSecret(client secrets.SecretsClient, secretId string) {
	logger := l.FuncLog("GetSecret")

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
