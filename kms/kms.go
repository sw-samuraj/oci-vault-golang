package kms

import (
	"context"
	l "cz.sw-samuraj/oci-vault/logging"
	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/keymanagement"
)

func GetKmsClient() keymanagement.KmsVaultClient {
	log := l.FuncLog("GetKmsClient")
	client, err := keymanagement.NewKmsVaultClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatalf("can't get a vault kms client: %s", err)
	}
	log.Info("vaults kms client has been obtained")
	return client
}

func ListVaults(client keymanagement.KmsVaultClient, compartmentId string) []keymanagement.VaultSummary {
	log := l.FuncLog("ListVaults")

	request := keymanagement.ListVaultsRequest{
		OpcRequestId:  common.String("42-list-my-vaults"),
		CompartmentId: common.String(compartmentId),
	}

	log.Info("calling the vault kms service...")
	response, err := client.ListVaults(context.Background(), request)
	if err != nil {
		log.Fatalf("can't get a response from the vault kms service: %s", err)
	}

	log.Info(response.Items)
	return response.Items
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
			log.Infof("vault will be scheduled for delete: %s (%s)", v.DisplayName, v.Id)
			toDelete = append(toDelete, v)
		}
	}

	if len(toDelete) == 0 {
		log.Info("no vaults in the state active, doing nothing")
		return
	}

	for _, v := range toDelete {
		log.Infof("vault will be scheduled for delete: %s (%s)", v.DisplayName, v.Id)

		request := keymanagement.ScheduleVaultDeletionRequest{
			OpcRequestId: common.String("42-delete-my-vault"),
			VaultId:      v.Id,
		}

		log.Info("calling the vault kms service...")
		response, err := client.ScheduleVaultDeletion(context.Background(), request)
		if err != nil {
			log.Fatalf("can't get a response from the secret service: %s", err)
		}
		log.Infof("vault has been scheduled for delete: %s (%s)", response.DisplayName, response.Id)
	}

	log.Infof("%d vaults have been scheduled for deletion", len(toDelete))
}
