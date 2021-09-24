package main

import (
	k "cz.sw-samuraj/oci-vault/kms"
	l "cz.sw-samuraj/oci-vault/logging"
	"github.com/sirupsen/logrus"
)

const (
	// compartment: sw-samuraj/vault-golang
	compartmentId = "ocid1.compartment.oc1..aaaaaaaangukhnc5jjl34tzm6dbh3blca3f4uviti3niavpttn6qgxddlsna"
	secretId      = "ocid1.vaultsecret.oc1.phx.amaaaaaayrywvyyatfzgkrnxmdr2crt5npuq66yrryc6qviapidlqsyeinza"
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
	kmsClient := k.GetKmsClient()

	log.Info("listing vaults...")
	vaults := k.ListVaults(kmsClient, compartmentId)

	log.Info("deleting existing vaults...")
	k.DeleteExistingVaults(kmsClient, vaults)

	log.Info("creating a new vault...")
	vaultId := k.CreateVault(kmsClient, compartmentId)

	log.Info("checking vault availability...")
	k.CheckVaultAvailability(kmsClient, vaultId)

	log.Info("getting the vaults client...")
	// vaultsClient := v.GetVaultsClient()

	log.Info("getting the secrets client...")
	// secretsClient := s.GetSecretsClient()

	log.Info("listing secrets...")
	// s.ListSecretVersions(secretsClient, secretId)
	log.Info("getting the secret...")
	// s.GetSecret(secretsClient, secretId)
}
