package pool

import (
	"context"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/pool"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
)

func DatasetResource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasetRead,
		CreateContext: datasetCreate,
		UpdateContext: datasetUpdate,
		DeleteContext: datasetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: datasetImport,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"encryption_root": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_loaded": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mountpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deduplication": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"aclmode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"acltype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"xattr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"atime": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"casesensitivity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exec": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compressratio": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quota": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"refquota": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reservation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"refreservation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"copies": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapdir": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"readonly": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recordsize": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"used": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"available": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"special_small_block_size": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pbkdf2iters": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func a(list []*pool.Dataset) []interface{} {
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

func datasetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	dataset, err := m.(v2.Client).PoolDataset().Get(ctx, d.Id())
	if err != nil {
		return diag.Errorf("failed to get dataset: %s", err)
	}

	d.Set("id", dataset.ID)
	d.Set("type", dataset.Type)
	d.Set("name", dataset.Name)
	d.Set("pool", dataset.Pool)
	d.Set("encrypted", dataset.Encrypted)
	d.Set("encryption_root", dataset.EncryptionRoot)
	d.Set("key_loaded", dataset.KeyLoaded)
	d.Set("mountpoint", dataset.Mountpoint)
	d.Set("deduplication", dataset.Deduplication.String())
	d.Set("aclmode", dataset.Aclmode.String())
	d.Set("acltype", dataset.Acltype.String())
	d.Set("xattr", dataset.Xattr.String())
	d.Set("atime", dataset.Atime.String())
	d.Set("casesensitivity", dataset.Casesensitivity.String())
	d.Set("exec", dataset.Exec.String())
	d.Set("sync", dataset.Sync.String())
	d.Set("compression", dataset.Compression.String())
	d.Set("compressratio", dataset.Compressratio.String())
	d.Set("origin", dataset.Origin.String())
	d.Set("quota", dataset.Quota.String())
	d.Set("refquota", dataset.Refquota.String())
	d.Set("reservation", dataset.Reservation.String())
	d.Set("refreservation", dataset.Refreservation.String())
	d.Set("copies", dataset.Copies.String())
	d.Set("snapdir", dataset.Snapdir.String())
	d.Set("readonly", dataset.Readonly.String())
	d.Set("recordsize", dataset.Recordsize.String())
	d.Set("key_format", dataset.KeyFormat.String())
	d.Set("encryption_algorithm", dataset.EncryptionAlgorithm.String())
	d.Set("used", dataset.Used.String())
	d.Set("available", dataset.Available.String())
	d.Set("special_small_block_size", dataset.SpecialSmallBlockSize.String())
	d.Set("pbkdf2iters", dataset.Pbkdf2Iters.String())
	d.Set("locked", dataset.Locked)
	d.SetId(d.Id())

	return nil
}

func datasetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return datasetRead(ctx, d, m)
}

func datasetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return datasetRead(ctx, d, m)
}

func datasetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func datasetImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	_, err := m.(v2.Client).PoolDataset().Get(ctx, d.Id())
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
