package vault

import (
	_ "encoding/json"

	"github.com/hashicorp/terraform/helper/schema"

	_ "github.com/hashicorp/vault/api"
)

func genericSecretResource() *schema.Resource {
	return &schema.Resource{
		Create: genericSecretResourceCreate,
		Update: genericSecretResourceUpdate,
		Delete: genericSecretResourceDelete,
		Read:   genericSecretResourceRead,

		Schema: map[string]*schema.Schema{

			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Full path where the generic secret will be written.",
			},

			// Data is passed as JSON so that an arbitrary structure is
			// possible, rather than forcing e.g. all values to be strings.
			"data_json": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "JSON-encoded secret data to write.",
				StateFunc:   obscureSecretInState,
			},
		},
	}
}

func genericSecretResourceCreate(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*api.Client)

	//client.Logical().Write()

	return nil
}

func genericSecretResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*api.Client)

	//client.Logical().Write()

	return nil
}

func genericSecretResourceDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*api.Client)

	return nil
}

func genericSecretResourceRead(d *schema.ResourceData, meta interface{}) error {
	// We don't actually attempt to read back the secret data
	// here, but instead just verify that the key still exists
	// in the backend so it can be re-created if it's deleted.
	//
	// This allows generic secrets to be created using a token
	// that has only *write* access to the store.

	return nil
}
