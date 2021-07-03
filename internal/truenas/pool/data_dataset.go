package pool

import (
	"context"
	"time"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/pool"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
)

func DataDatasetResource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataDatasetRead,
		Schema: map[string]*schema.Schema{
			"truenas_pool_dataset": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"encrypted": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"encryption_root": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_loaded": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mountpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deduplication": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aclmode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"acltype": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"xattr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"atime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"casesensitivity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"exec": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sync": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compression": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compressratio": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"origin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quota": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"refquota": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reservation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"refreservation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"copies": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snapdir": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"readonly": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recordsize": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"encryption_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"used": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"special_small_block_size": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pbkdf2iters": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locked": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenDatasetData(list []*pool.Dataset) []interface{} {
	if list == nil {
		return make([]interface{}, 0)
	}

	l := make([]interface{}, len(list), len(list))
	for i, dataset := range list {
		d := make(map[string]interface{})
		d["id"] = dataset.ID
		d["type"] = dataset.Type
		d["name"] = dataset.Name
		d["pool"] = dataset.Pool
		d["encrypted"] = dataset.Encrypted
		d["encryption_root"] = dataset.EncryptionRoot
		d["key_loaded"] = dataset.KeyLoaded
		d["mountpoint"] = dataset.Mountpoint
		d["deduplication"] = dataset.Deduplication.String()
		d["aclmode"] = dataset.Aclmode.String()
		d["acltype"] = dataset.Acltype.String()
		d["xattr"] = dataset.Xattr.String()
		d["atime"] = dataset.Atime.String()
		d["casesensitivity"] = dataset.Casesensitivity.String()
		d["exec"] = dataset.Exec.String()
		d["sync"] = dataset.Sync.String()
		d["compression"] = dataset.Compression.String()
		d["compressratio"] = dataset.Compressratio.String()
		d["origin"] = dataset.Origin.String()
		d["quota"] = dataset.Quota.String()
		d["refquota"] = dataset.Refquota.String()
		d["reservation"] = dataset.Reservation.String()
		d["refreservation"] = dataset.Refreservation.String()
		d["copies"] = dataset.Copies.String()
		d["snapdir"] = dataset.Snapdir.String()
		d["readonly"] = dataset.Readonly.String()
		d["recordsize"] = dataset.Recordsize.String()
		d["key_format"] = dataset.KeyFormat.String()
		d["encryption_algorithm"] = dataset.EncryptionAlgorithm.String()
		d["used"] = dataset.Used.String()
		d["available"] = dataset.Available.String()
		d["special_small_block_size"] = dataset.SpecialSmallBlockSize.String()
		d["pbkdf2iters"] = dataset.Pbkdf2Iters.String()
		d["locked"] = dataset.Locked

		l[i] = d
	}
	return l
}

func dataDatasetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	list, err := m.(v2.Client).PoolDataset().List(ctx)
	if err != nil {
		return diag.Errorf("failed to get dataset list: %s", err)
	}

	err = d.Set("truenas_pool_dataset", flattenDatasetData(list))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(time.Now().String())

	return nil
}
