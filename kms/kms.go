package kms

import (
	"context"
	l "cz.sw-samuraj/oci-vault/logging"
	"fmt"
	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/keymanagement"
	"time"
)

func GetKmsVaultClient() keymanagement.KmsVaultClient {
	log := l.FuncLog("GetKmsVaultClient")
	client, err := keymanagement.NewKmsVaultClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatalf("can't get a kms vault client: %s", err)
	}
	log.Info("vaults kms client has been obtained")
	return client
}

func GetKmsManagementClient(endpoint *string) keymanagement.KmsManagementClient {
	log := l.FuncLog("GetKmsManagementClient")
	client, err := keymanagement.NewKmsManagementClientWithConfigurationProvider(common.DefaultConfigProvider(), *endpoint)
	if err != nil {
		log.Fatalf("can't get a kms management client: %s", err)
	}
	log.Info("kms management client has been obtained")
	return client
}

func ListVaults(client keymanagement.KmsVaultClient, compartmentId string) []keymanagement.VaultSummary {
	log := l.FuncLog("ListVaults")

	request := keymanagement.ListVaultsRequest{
		OpcRequestId:  common.String("42-list-my-vaults"),
		CompartmentId: common.String(compartmentId),
	}

	log.Info("calling the kms vault service...")
	response, err := client.ListVaults(context.Background(), request)
	if err != nil {
		log.Fatalf("can't get a response from the kms vault service: %s", err)
	}

	log.Infof("vaults in compartment: %s", formatVaults(response.Items))
	return response.Items
}

func formatVaults(vaults []keymanagement.VaultSummary) string {
	vaultsString := ""
	for _, v := range vaults {
		vaultId := fmt.Sprintf("%s...%s", (*v.Id)[:42], (*v.Id)[92:])
		timeCreated := v.TimeCreated.Format(time.RFC3339)
		vaultsString += fmt.Sprintf("\n    %s, %s, %s, %s", *v.DisplayName, v.LifecycleState, timeCreated, vaultId)
	}
	return vaultsString
}

func DeleteExistingVaults(client keymanagement.KmsVaultClient, vaults []keymanagement.VaultSummary) {
	log := l.FuncLog("DeleteExistingVaults")

	if len(vaults) == 0 {
		log.Info("no vaults in the compartment, doing nothing")
		return
	}

	var toDelete []keymanagement.VaultSummary
	for _, v := range vaults {
		if v.LifecycleState == keymanagement.VaultSummaryLifecycleStateActive {
			log.Infof("vault will be scheduled for delete: %s (%s)", *v.DisplayName, *v.Id)
			toDelete = append(toDelete, v)
		}
	}

	if len(toDelete) == 0 {
		log.Info("no vaults in the state active, doing nothing")
		return
	}

	for _, v := range toDelete {
		request := keymanagement.ScheduleVaultDeletionRequest{
			OpcRequestId: common.String("42-delete-my-vault"),
			VaultId:      v.Id,
		}

		log.Info("calling the kms vault service...")
		response, err := client.ScheduleVaultDeletion(context.Background(), request)
		if err != nil {
			log.Fatalf("can't get a response from the secret service: %s", err)
		}
		log.Infof("vault has been scheduled for delete: %s (%s)", *response.DisplayName, *response.Id)
	}

	log.Infof("%d vaults have been scheduled for deletion", len(toDelete))
}

func CreateVault(client keymanagement.KmsVaultClient, compartmentId string) *string {
	log := l.FuncLog("CreateVault")

	request := keymanagement.CreateVaultRequest{
		OpcRequestId: common.String("42-create-my-vault"),
		CreateVaultDetails: keymanagement.CreateVaultDetails{
			CompartmentId: common.String(compartmentId),
			DisplayName:   common.String("sw-samuraj-vault"),
			VaultType:     keymanagement.CreateVaultDetailsVaultTypeDefault,
		},
	}

	log.Info("calling the kms vault service...")
	response, err := client.CreateVault(context.Background(), request)
	if err != nil {
		log.Fatalf("can't get a response from the kms vault service: %s", err)
	}

	log.Infof("vault has been created: %s", response.Vault)
	return response.Id
}

func GetManagementEndpoint(client keymanagement.KmsVaultClient, vaultId *string) *string {
	log := l.FuncLog("GetManagementEndpoint")

	request := keymanagement.GetVaultRequest{
		OpcRequestId: common.String("42-get-my-vault"),
		VaultId:      vaultId,
	}

	for {
		response, err := client.GetVault(context.Background(), request)
		if err != nil {
			log.Errorf("can't get a response from the kms vault service: %s", err)
		}
		if response.LifecycleState == keymanagement.VaultLifecycleStateActive {
			log.Infof("vault is in the state: %s", response.LifecycleState)
			log.Infof("management endpoint has been received: %s", *response.ManagementEndpoint)
			return response.ManagementEndpoint
		} else {
			log.Infof("vault is still in the state %s, waiting for 15s", response.LifecycleState)
			time.Sleep(15 * time.Second)
		}
	}
}

func CreateMasterKey(client keymanagement.KmsManagementClient, compartmentId string) *string {
	log := l.FuncLog("CreateMasterKey")

	request := keymanagement.CreateKeyRequest{
		OpcRequestId: common.String("42-create-my-key"),
		CreateKeyDetails: keymanagement.CreateKeyDetails{
			CompartmentId: common.String(compartmentId),
			DisplayName:   common.String("sw-samuraj-master-key"),
			KeyShape: &keymanagement.KeyShape{
				Algorithm: keymanagement.KeyShapeAlgorithmAes,
				Length:    common.Int(32),
			},
			ProtectionMode: keymanagement.CreateKeyDetailsProtectionModeHsm,
		},
	}

	log.Info("calling the kms management service...")
	response, err := client.CreateKey(context.Background(), request)
	if err != nil {
		log.Fatalf("can't get a response from the kms management service: %s", err)
	}

	log.Infof("master key has been created: %s", response.Key)
	return response.Id
}
