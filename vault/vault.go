package vault

import (
	"context"
	l "cz.sw-samuraj/oci-vault/logging"
	"encoding/base64"
	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/vault"
	"time"
)

func GetVaultsClient() vault.VaultsClient {
	log := l.FuncLog("GetVaultsClient")
	client, err := vault.NewVaultsClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatalf("can't get a vaults client: %s", err)
	}
	log.Info("vaults client has been obtained")
	return client
}

func CreateSecret(client vault.VaultsClient, compartmentId string, vaultId *string, keyId *string, secret string) *string {
	log := l.FuncLog("CreateSecret")

	request := vault.CreateSecretRequest{
		OpcRequestId: common.String("42-create-my-secret"),
		CreateSecretDetails: vault.CreateSecretDetails{
			CompartmentId: common.String(compartmentId),
			SecretName:    common.String("sw-samuraj-private-key-" + getTimestampString()),
			VaultId:       vaultId,
			KeyId:         keyId,
			SecretContent: vault.Base64SecretContentDetails{
				Content: common.String(base64.StdEncoding.EncodeToString([]byte(secret))),
			},
		},
	}

	log.Info("calling the vault service...")
	response, err := client.CreateSecret(context.Background(), request)
	if err != nil {
		log.Fatalf("can't get a response from the vault service: %s", err)
	}

	log.Infof("secret has been created: %s", response.Secret)
	return response.Id
}

func getTimestampString() string {
	return time.Now().Format("20060102150105")
}
