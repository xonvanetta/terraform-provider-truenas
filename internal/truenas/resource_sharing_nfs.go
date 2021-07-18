package truenas

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/sharing"
)

func resourceSharingNFS() *schema.Resource {
	return &schema.Resource{
		ReadContext:   sharingNFSRead,
		CreateContext: sharingNFSCreate,
		UpdateContext: sharingNFSUpdate,
		DeleteContext: sharingNFSDelete,
		Importer: &schema.ResourceImporter{
			StateContext: sharingNFSImport,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"paths": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"aliases": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"networks": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hosts": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alldirs": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ro": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"quiet": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"maproot_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maproot_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapall_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapall_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func nfsFromSchema(d *schema.ResourceData) *sharing.NFS {
	return &sharing.NFS{
		Paths:        toStrings(d.Get("paths").(*schema.Set).List()),
		Aliases:      toStrings(d.Get("aliases").(*schema.Set).List()),
		Comment:      d.Get("comment").(string),
		Networks:     toStrings(d.Get("networks").(*schema.Set).List()),
		Hosts:        toStrings(d.Get("hosts").(*schema.Set).List()),
		Alldirs:      d.Get("alldirs").(bool),
		Ro:           d.Get("ro").(bool),
		Quiet:        d.Get("quiet").(bool),
		MaprootUser:  d.Get("maproot_user").(string),
		MaprootGroup: d.Get("maproot_group").(string),
		MapallUser:   d.Get("mapall_user").(string),
		MapallGroup:  d.Get("mapall_group").(string),
		Security:     toStrings(d.Get("security").(*schema.Set).List()),
		Enabled:      d.Get("enabled").(bool),
	}
}

func sharingNFSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	nfs := nfsFromSchema(d)

	err := m.(v2.Client).SharingNFS().Create(ctx, nfs)
	if err != nil {
		return diag.Errorf("failed to create nfs: %s, %v", err, nfs)
	}

	d.SetId(strconv.Itoa(nfs.Id))
	return sharingNFSRead(ctx, d, m)
}

func sharingNFSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Errorf("failed to format id to int: %s", err)
	}
	nfs, err := m.(v2.Client).SharingNFS().Get(ctx, id)
	if err != nil {
		return diag.Errorf("failed to get nfs: %s", err)
	}

	d.Set("id", strconv.Itoa(nfs.Id))
	d.Set("paths", nfs.Paths)
	d.Set("enabled", nfs.Enabled)
	d.Set("paths", nfs.Paths)
	d.Set("aliases", nfs.Aliases)
	d.Set("comment", nfs.Comment)
	d.Set("networks", nfs.Networks)
	d.Set("hosts", nfs.Hosts)
	d.Set("alldirs", nfs.Alldirs)
	d.Set("ro", nfs.Ro)
	d.Set("quiet", nfs.Quiet)
	d.Set("maproot_user", nfs.MaprootUser)
	d.Set("maproot_group", nfs.MaprootGroup)
	d.Set("mapall_user", nfs.MapallUser)
	d.Set("mapall_group", nfs.MapallGroup)
	d.Set("security", nfs.Security)
	d.Set("enabled", nfs.Enabled)

	d.SetId(d.Id())
	return nil
}

func sharingNFSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	nfs := nfsFromSchema(d)

	var err error
	nfs.Id, err = strconv.Atoi(d.Id())
	if err != nil {
		return diag.Errorf("failed to format id to int: %s", err)
	}

	err = m.(v2.Client).SharingNFS().Update(ctx, nfs)
	if err != nil {
		return diag.Errorf("failed to update nfs: %s", err)
	}

	return sharingNFSRead(ctx, d, m)
}

func sharingNFSDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Errorf("failed to format id to int: %s", err)
	}

	err = m.(v2.Client).SharingNFS().Delete(ctx, id)
	if err != nil {
		return diag.Errorf("failed to delete nfs: %s", err)
	}

	return nil
}

func sharingNFSImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed to format id to int: %w", err)
	}

	_, err = m.(v2.Client).SharingNFS().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
