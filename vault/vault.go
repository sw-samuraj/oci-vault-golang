package vault

import (
	l "cz.sw-samuraj/oci-vault/logging"
	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/vault"
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
