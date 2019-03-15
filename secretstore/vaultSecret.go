package secretstore

import (
	"github.com/hashicorp/vault/api"
	"strings"
	"fmt"
)

// VaultSecret contains a key value Vaultsecret
type VaultSecret struct {
	key string
	value string
	client *api.Client
}

// NewVaultSecret is the default constructor for VaultSecret
func NewVaultSecret() *VaultSecret {
	return &VaultSecret{
	}
}

// CreateVaultSecret creates a VaultSecret from key and value
func CreateVaultSecret(key, value string) (*VaultSecret) {
	return &VaultSecret{
		key: key,
		value: value,
		client: nil,
	}
}

// VaultSecretFromKey creates a Vaultsecret from a key, it does not fetch the associated value.
func VaultSecretFromKey(key string) (*VaultSecret) {
	return &VaultSecret{
		key: key,
		value: "",
		client: nil,
	}
}

// SetKey setter of attribute key for VaultSecret struct object
func (sec *VaultSecret) SetKey(key string) *VaultSecret {
	sec.key = key
	return sec
}

// SetValue setter of attribute value for VaultSecret struct object
func (sec *VaultSecret) SetValue(value string) *VaultSecret {
	sec.value = value
	return sec	
}

// SetClient setter of attribute client for VaultSecret struct object
func (sec *VaultSecret) SetClient(client *api.Client) *VaultSecret {
	sec.client = client
	return sec	
}

// SaveNew stores a new Vaultsecret in the vault
func (sec *VaultSecret) SaveNew() (res *api.VaultSecret, err error) {

	fetched, err := sec.GetValue()
	if fetched != "" {
		return nil, fmt.Errorf("This Vaultsecret already exists : " + sec.key)
	}

	return sec.Update()
}

// GetValue fetch the value from AWS VaultSecretManager by key
func (sec *VaultSecret) GetValue() (string, error) {

	log := sec.client.Logical()
	res, err := log.Read(
		strings.Join([]string{"Vaultsecret", sec.key}, "/"),
	)

	if err != nil {
		return "", err
	}

	sec.value = res.Data["value"].(string)

	return sec.value, nil
}

// Update the Vaultsecret value stored in the aws Vaultsecret manager
func (sec *VaultSecret) Update() (error) {

	log := sec.client.Logical()
	res, err := log.Write(
		strings.Join([]string{"Vaultsecret", sec.key}, "/"),
		map[string]interface{}{ "value": sec.value },
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}


// Delete remove the key from the Vaultsecret manager
func (sec *VaultSecret) Delete() (error) {

	log := sec.client.Logical()
	res, err := log.Delete(
		strings.Join([]string{"Vaultsecret", sec.key}, "/"),
	)

	if err != nil {
		return nil, err
	}

	return res, nil	
} 

// List retrieve all the keys availables in the Vaultsecret manager
func (sec *VaultSecret) List() ([]string, error) {

	log := sec.client.Logical()
	res, err := log.List("Vaultsecret")

	if err != nil {
		return nil, err
	}

	list := make([]string, len(res.Data))
	i := 0
	for address := range res.Data {
		list[i] = address
		i++
	}

	return list, nil	
} 