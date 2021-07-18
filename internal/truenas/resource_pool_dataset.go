package truenas

import (
	"context"

	"code.cloudfoundry.org/bytefmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/pool"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
)

func resourcePoolDataset() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourcePoolDatasetRead,
		CreateContext: resourcePoolDatasetCreate,
		UpdateContext: resourcePoolDatasetUpdate,
		DeleteContext: resourcePoolDatasetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePoolDatasetImport,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:         schema.TypeString,
				Default:      "FILESYSTEM",
				ValidateFunc: validation.StringInSlice([]string{"FILESYSTEM", "VOLUME"}, false),
				Required:     true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volsize": {
				Type:        schema.TypeString,
				Description: "Size in E,P,T,G,M,K,B bytes",
				Optional:    true,
			},
			"volblocksize": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"512", "1K", "2K", "4K", "8K", "16K", "32K", "64K", "128K"}, true),
				Optional:     true,
				Default:      "16K",
			},
			"sparse": {
				Type:        schema.TypeBool,
				Description: "used by VOLUME to use Thin provisioning, ONLY USED WITH CREATING THE VOLUME!",
				Optional:    true,
			},
			"force_size": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sync": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"STANDARD", "ALWAYS", "DISABLED"}, false),
				Optional:     true,
			},
			"compression": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"OFF", "LZ4", "GZIP", "GZIP-1", "GZIP-9", "ZSTD", "ZSTD-FAST", "ZLE", "LZJB", "ZSTD-1", "ZSTD-2", "ZSTD-3", "ZSTD-4", "ZSTD-5", "ZSTD-6", "ZSTD-7", "ZSTD-8", "ZSTD-9", "ZSTD-10", "ZSTD-11", "ZSTD-12", "ZSTD-13", "ZSTD-14", "ZSTD-15", "ZSTD-16", "ZSTD-17", "ZSTD-18", "ZSTD-19", "ZSTD-FAST-1", "ZSTD-FAST-2", "ZSTD-FAST-3", "ZSTD-FAST-4", "ZSTD-FAST-5", "ZSTD-FAST-6", "ZSTD-FAST-7", "ZSTD-FAST-8", "ZSTD-FAST-9", "ZSTD-FAST-10", "ZSTD-FAST-20", "ZSTD-FAST-30", "ZSTD-FAST-40", "ZSTD-FAST-50", "ZSTD-FAST-60", "ZSTD-FAST-70", "ZSTD-FAST-80", "ZSTD-FAST-90", "ZSTD-FAST-100", "ZSTD-FAST-500", "ZSTD-FAST-1000"}, false),
				Optional:     true,
			},
			"atime": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ON", "OFF"}, false),
				Optional:     true,
			},
			"exec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ON", "OFF"}, false),
				Optional:     true,
			},
			"deduplication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ON", "VERIFY", "OFF"}, false),
				Optional:     true,
			},
			"readonly": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ON", "OFF"}, false),
				Optional:     true,
			},
			"recordsize": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"512", "1K", "2K", "4K", "8K", "16K", "32K", "64K", "128K", "256K", "512K", "1024K"}, false),
				Optional:     true,
			},
			"casesensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"SENSITIVE", "INSENSITIVE", "MIXED"}, false),
				Optional:     true,
			},
			"aclmode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"PASSTHROUGH", "RESTRICTED"}, false),
				Optional:     true,
			},
			"acltype": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"NOACL", "NFS4ACL", "POSIXACL"}, false),
				Optional:     true,
			},
			"share_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"GENERIC", "SMB"}, false),
				Optional:     true,
			},
			"xattr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ON", "SA"}, false),
				Optional:     true,
			},
			"inherit_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourcePoolDatasetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	dataset, err := m.(v2.Client).PoolDataset().Get(ctx, d.Id())
	if err != nil {
		return diag.Errorf("failed to get dataset: %s", err)
	}

	d.Set("id", dataset.Name)
	d.Set("type", dataset.Type)
	d.Set("name", dataset.Name)
	d.Set("force_size", dataset.ForceSize.String())
	d.Set("comments", dataset.Comments.String())
	d.Set("sync", dataset.Sync.String())
	d.Set("compression", dataset.Compression.String())
	d.Set("atime", dataset.Atime.String())
	d.Set("exec", dataset.Exec.String())
	d.Set("deduplication", dataset.Deduplication.String())
	d.Set("readonly", dataset.Readonly.String())
	d.Set("recordsize", dataset.Recordsize.String())
	d.Set("aclmode", dataset.Aclmode.String())
	d.Set("acltype", dataset.Acltype.String())
	d.Set("xattr", dataset.Xattr.String())

	d.SetId(dataset.Name)

	if dataset.Type == "VOLUME" {
		d.Set("volblocksize", dataset.Volblocksize.String())
		d.Set("volsize", dataset.Volsize.String())
	}

	return nil
}

func resourcePoolDatasetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	dataset := &pool.Dataset{
		Type:          d.Get("type").(string),
		Name:          d.Get("name").(string),
		ForceSize:     pool.NewValue(d.Get("force_size")),
		Comments:      pool.NewValue(d.Get("comments").(string)),
		Sync:          pool.NewValue(d.Get("sync").(string)),
		Compression:   pool.NewValue(d.Get("compression").(string)),
		Atime:         pool.NewValue(d.Get("atime").(string)),
		Exec:          pool.NewValue(d.Get("exec").(string)),
		Deduplication: pool.NewValue(d.Get("deduplication").(string)),
		Readonly:      pool.NewValue(d.Get("readonly").(string)),
		Recordsize:    pool.NewValue(d.Get("recordsize").(string)),
		Aclmode:       pool.NewValue(d.Get("aclmode").(string)),
		Acltype:       pool.NewValue(d.Get("acltype").(string)),
		Xattr:         pool.NewValue(d.Get("xattr").(string)),
	}

	if dataset.Type == "VOLUME" {
		dataset.Volblocksize = pool.NewValue(d.Get("volblocksize"))
		dataset.Sparse = pool.NewValue(d.Get("sparse"))

		volsize, err := bytefmt.ToBytes(d.Get("volsize").(string))
		if err != nil {
			return diag.Errorf("failed to convert string to bytes: %s", err)
		}

		if volsize == 0 {
			return diag.Errorf("volsize must more than 0 in size on volume")
		}
		dataset.Volsize = pool.NewValue(volsize)
	}

	err := m.(v2.Client).PoolDataset().Create(ctx, dataset)
	if err != nil {
		return diag.Errorf("failed to create dataset: %s", err)
	}
	d.SetId(dataset.Name)

	return resourcePoolDatasetRead(ctx, d, m)
}

func resourcePoolDatasetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	dataset := &pool.Dataset{
		ForceSize:     pool.NewValue(d.Get("force_size")),
		Comments:      pool.NewValue(d.Get("comments").(string)),
		Sync:          pool.NewValue(d.Get("sync").(string)),
		Compression:   pool.NewValue(d.Get("compression").(string)),
		Atime:         pool.NewValue(d.Get("atime").(string)),
		Exec:          pool.NewValue(d.Get("exec").(string)),
		Deduplication: pool.NewValue(d.Get("deduplication").(string)),
		Readonly:      pool.NewValue(d.Get("readonly").(string)),
		Recordsize:    pool.NewValue(d.Get("recordsize").(string)),
		Aclmode:       pool.NewValue(d.Get("aclmode").(string)),
		Acltype:       pool.NewValue(d.Get("acltype").(string)),
		Xattr:         pool.NewValue(d.Get("xattr").(string)),
	}

	if d.Get("type") == "VOLUME" {
		volsize, err := bytefmt.ToBytes(d.Get("volsize").(string))
		if err != nil {
			return diag.Errorf("failed to convert string to bytes: %s", err)
		}

		if volsize == 0 {
			return diag.Errorf("volsize must more than 0 in size on volume")
		}
		dataset.Volsize = pool.NewValue(volsize)
	}

	return resourcePoolDatasetRead(ctx, d, m)
}

func resourcePoolDatasetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := m.(v2.Client).PoolDataset().Delete(ctx, d.Id())
	if err != nil {
		return diag.Errorf("failed to delete dataset: %s", err)
	}

	return nil
}

func resourcePoolDatasetImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	_, err := m.(v2.Client).PoolDataset().Get(ctx, d.Id())
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
