package main

import (
	k "cz.sw-samuraj/oci-vault/kms"
	l "cz.sw-samuraj/oci-vault/logging"
	s "cz.sw-samuraj/oci-vault/secrets"
	v "cz.sw-samuraj/oci-vault/vault"
	"github.com/sirupsen/logrus"
)

const (
	// compartment: sw-samuraj/vault-golang
	// compartmentId = "ocid1.compartment.oc1..aaaaaaaangukhnc5jjl34tzm6dbh3blca3f4uviti3niavpttn6qgxddlsna"
	compartmentId = "ocid1.compartment.oc1..aaaaaaaavwsgxcpgcfp7hn4ojqej26kofxjtqtqs4bertq2qfq5s5qycvxzq"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	log := l.FuncLog("main")
	log.Info("starting the vault service example")

	log.Info("getting the vaults kms client...")
	kmsVaultClient := k.GetKmsVaultClient()

	log.Info("listing vaults...")
	vaults := k.ListVaults(kmsVaultClient, compartmentId)

	log.Info("deleting existing vaults...")
	k.DeleteExistingVaults(kmsVaultClient, vaults)

	log.Info("creating a new vault...")
	vaultId := k.CreateVault(kmsVaultClient, compartmentId)

	log.Info("checking vault availability, waiting for management endpoint...")
	managementEndpoint := k.GetManagementEndpoint(kmsVaultClient, vaultId)

	log.Info("getting the kms management client...")
	kmsManagementClient := k.GetKmsManagementClient(managementEndpoint)

	log.Info("creating a new master key...")
	keyId := k.CreateMasterKey(kmsManagementClient, compartmentId)

	log.Info("checking key availability...")
	k.CheckKeyAvailability(kmsManagementClient, keyId)

	log.Info("getting the vaults client...")
	vaultsClient := v.GetVaultsClient()

	secret := "my-secret-key"
	log.Info("creating a new secret...")
	secretId := v.CreateSecret(vaultsClient, compartmentId, vaultId, keyId, secret)

	log.Info("getting the secrets client...")
	secretsClient := s.GetSecretsClient()

	log.Info("listing secrets...")
	s.ListSecretVersions(secretsClient, secretId)

	log.Info("getting the secret...")
	s.GetSecret(secretsClient, secretId)
}
