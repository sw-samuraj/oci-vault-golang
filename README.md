# OCI Vault and Golang

This is an supporting project for the article about managing
[OCI Vault](//www.oracle.com/security/cloud-security/key-management/)
service via Golang SDK.

You can read the article here:
[OCI Vault & Secrets service via Golang](//medium.com/@sw_samuraj/)

## How to build the example

```
go mod tidy
go build
```

## Necessary policies

Before you run the example, you need to have following policies
in your tenancy:

```
Allow group vault-admins to manage vaults in compartment <your-compartment>
Allow group vault-admins to manage keys in compartment <your-compartment>
Allow group vault-admins to manage secret-family in compartment <your-compartment>
```

## How to run the example

```
./oci-vault
```

## What it does?

1. Lists all the vaults in given compartment (see the `compartmentId` constant).
1. Schedules for delete all the _active_ vaults (in the given compartment).
   **Be careful to not accidentaly destroy any important vault!!!** I recommend
   to play in dedicated and empty compartment!
1. Creates a new vault.
1. Creates a new master key.
1. Creates a new secret.
1. Lists versions for newly created secret (there will be only one --- current/latest).
1. Reads the secret.

## How it should look like?

When everything works fine, you should see following output:

```
INFO[2021-10-01T15:09:33+02:00] starting the vault service example            func=main
INFO[2021-10-01T15:09:33+02:00] getting the vaults kms client...              func=main
INFO[2021-10-01T15:09:33+02:00] vaults kms client has been obtained           func=GetKmsVaultClient
INFO[2021-10-01T15:09:33+02:00] listing vaults...                             func=main
INFO[2021-10-01T15:09:33+02:00] calling the kms vault service...              func=ListVaults
INFO[2021-10-01T15:09:35+02:00] vaults in compartment:                        func=ListVaults
INFO[2021-10-01T15:09:35+02:00] deleting existing vaults...                   func=main
INFO[2021-10-01T15:09:35+02:00] no vaults in the compartment, doing nothing   func=DeleteExistingVaults
INFO[2021-10-01T15:09:35+02:00] creating a new vault...                       func=main
INFO[2021-10-01T15:09:35+02:00] calling the kms vault service...              func=CreateVault
INFO[2021-10-01T15:09:36+02:00] vault has been created: { CompartmentId=ocid1.compartment.oc1..aaaaaaaavwsgxcpgcfp7hn4ojqej26kofxjtqtqs4bertq2qfq5s5qycvxzq CryptoEndpoint=<nil> DisplayName=sw-samuraj-vault Id=ocid1.vault.oc1.eu-frankfurt-1.b5qvocelaaaao.abtheljs5kjow2l2xkiv2w3c4daptizehxbjavnvnmr42qaqqic5xcyl3c7a LifecycleState=CREATING ManagementEndpoint=<nil> TimeCreated=2021-10-01 13:09:31.869 +0000 UTC VaultType=DEFAULT WrappingkeyId= DefinedTags=map[] FreeformTags=map[] TimeOfDeletion=<nil> RestoredFromVaultId=<nil> ReplicaDetails=<nil> IsPrimary=true }  func=CreateVault
INFO[2021-10-01T15:09:36+02:00] checking vault availability, waiting for management endpoint...  func=main
INFO[2021-10-01T15:09:36+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:09:52+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:10:08+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:10:24+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:10:39+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:10:55+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:11:11+02:00] vault is still in the state CREATING, waiting for 15s  func=GetManagementEndpoint
INFO[2021-10-01T15:11:27+02:00] vault is in the state: ACTIVE                 func=GetManagementEndpoint
INFO[2021-10-01T15:11:27+02:00] management endpoint has been received: https://b5qvocelaaaao-management.kms.eu-frankfurt-1.oraclecloud.com  func=GetManagementEndpoint
INFO[2021-10-01T15:11:27+02:00] getting the kms management client...          func=main
INFO[2021-10-01T15:11:27+02:00] kms management client has been obtained       func=GetKmsManagementClient
INFO[2021-10-01T15:11:27+02:00] creating a new master key...                  func=main
INFO[2021-10-01T15:11:27+02:00] calling the kms management service...         func=CreateMasterKey
INFO[2021-10-01T15:11:29+02:00] master key has been created: { CompartmentId=ocid1.compartment.oc1..aaaaaaaavwsgxcpgcfp7hn4ojqej26kofxjtqtqs4bertq2qfq5s5qycvxzq CurrentKeyVersion=ocid1.keyversion.oc1.eu-frankfurt-1.b5qvocelaaaao.bc4tmflqr7iaa.abtheljtadfavvzn55zp7qmc25y7ursoltsgol5jy5hdax2c2mind3isl4vq DisplayName=sw-samuraj-master-key Id=ocid1.key.oc1.eu-frankfurt-1.b5qvocelaaaao.abtheljtuqxwtfwpgq6odjl4nuk2k5un6wq574rze6hkmmfiwrbgeswfwaxa KeyShape={ Algorithm=AES Length=32 CurveId= } LifecycleState=CREATING TimeCreated=2021-10-01 13:11:25.44 +0000 UTC VaultId=ocid1.vault.oc1.eu-frankfurt-1.b5qvocelaaaao.abtheljs5kjow2l2xkiv2w3c4daptizehxbjavnvnmr42qaqqic5xcyl3c7a DefinedTags=map[] FreeformTags=map[] ProtectionMode=HSM TimeOfDeletion=<nil> RestoredFromKeyId=<nil> ReplicaDetails=<nil> IsPrimary=true }  func=CreateMasterKey
INFO[2021-10-01T15:11:29+02:00] checking key availability...                  func=main
INFO[2021-10-01T15:11:30+02:00] key is still in the state CREATING, waiting for 10s  func=CheckKeyAvailability
INFO[2021-10-01T15:11:41+02:00] key is still in the state CREATING, waiting for 10s  func=CheckKeyAvailability
INFO[2021-10-01T15:11:51+02:00] key is still in the state CREATING, waiting for 10s  func=CheckKeyAvailability
INFO[2021-10-01T15:12:02+02:00] key is still in the state CREATING, waiting for 10s  func=CheckKeyAvailability
INFO[2021-10-01T15:12:13+02:00] key is in the state: ENABLED                  func=CheckKeyAvailability
INFO[2021-10-01T15:12:13+02:00] getting the vaults client...                  func=main
INFO[2021-10-01T15:12:13+02:00] vaults client has been obtained               func=GetVaultsClient
INFO[2021-10-01T15:12:13+02:00] creating a new secret...                      func=main
INFO[2021-10-01T15:12:13+02:00] calling the vault service...                  func=CreateSecret
INFO[2021-10-01T15:12:16+02:00] secret has been created: { CompartmentId=ocid1.compartment.oc1..aaaaaaaavwsgxcpgcfp7hn4ojqej26kofxjtqtqs4bertq2qfq5s5qycvxzq Id=ocid1.vaultsecret.oc1.eu-frankfurt-1.amaaaaaayrywvyyakk7nxw2thavsloirwji7buk4zwx2z34ufvbhivoikmuq LifecycleState=CREATING SecretName=sw-samuraj-private-key-20211001151013 TimeCreated=2021-10-01 13:12:11.8 +0000 UTC VaultId=ocid1.vault.oc1.eu-frankfurt-1.b5qvocelaaaao.abtheljs5kjow2l2xkiv2w3c4daptizehxbjavnvnmr42qaqqic5xcyl3c7a CurrentVersionNumber=1 DefinedTags=map[] Description=<nil> FreeformTags=map[] KeyId=ocid1.key.oc1.eu-frankfurt-1.b5qvocelaaaao.abtheljtuqxwtfwpgq6odjl4nuk2k5un6wq574rze6hkmmfiwrbgeswfwaxa LifecycleDetails=<nil> Metadata=map[] SecretRules=[] TimeOfCurrentVersionExpiry=<nil> TimeOfDeletion=<nil> }  func=CreateSecret
INFO[2021-10-01T15:12:16+02:00] getting the secrets client...                 func=main
INFO[2021-10-01T15:12:16+02:00] secrets client has been obtained              func=GetSecretsClient
INFO[2021-10-01T15:12:16+02:00] listing secrets...                            func=main
INFO[2021-10-01T15:12:16+02:00] calling the secret service...                 func=ListSecretVersions
INFO[2021-10-01T15:12:18+02:00] [{ SecretId=ocid1.vaultsecret.oc1.eu-frankfurt-1.amaaaaaayrywvyyakk7nxw2thavsloirwji7buk4zwx2z34ufvbhivoikmuq VersionNumber=1 TimeCreated=2021-10-01 13:12:11.8 +0000 UTC VersionName=<nil> TimeOfDeletion=<nil> TimeOfExpiry=<nil> Stages=[CURRENT LATEST] }]  func=ListSecretVersions
INFO[2021-10-01T15:12:18+02:00] getting the secret...                         func=main
INFO[2021-10-01T15:12:18+02:00] calling the secret service...                 func=GetSecret
INFO[2021-10-01T15:12:19+02:00] secret content: { Content=bXktc2VjcmV0LWtleQ== }  func=GetSecret
INFO[2021-10-01T15:12:19+02:00] decoded secret: my-secret-key                 func=GetSecret
```
