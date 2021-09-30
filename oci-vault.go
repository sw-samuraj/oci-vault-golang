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
	compartmentId = "ocid1.compartment.oc1..aaaaaaaangukhnc5jjl34tzm6dbh3blca3f4uviti3niavpttn6qgxddlsna"
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
	// vaults := k.ListVaults(kmsVaultClient, compartmentId)
	k.ListVaults(kmsVaultClient, compartmentId)

	// TODO Guido: temporarily disable vault creation because of limits
	// log.Info("deleting existing vaults...")
	// k.DeleteExistingVaults(kmsVaultClient, vaults)

	// log.Info("creating a new vault...")
	// vaultId := k.CreateVault(kmsVaultClient, compartmentId)

	// vi := "ocid1.vault.oc1.eu-frankfurt-1.b5qvell7aaaao.abtheljsp3iybqbx4awmcdyzisocmnjsa6ypqv3nu24v4qetgrjix4mokzaq"
	vi := "ocid1.vault.oc1.eu-frankfurt-1.b5qvelxgaafak.abtheljrn2uspkm4ojrnpuwgkd7syoohv4bel2qx5jlftt7ywofipoh2i5ja"
	vaultId := &vi
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
