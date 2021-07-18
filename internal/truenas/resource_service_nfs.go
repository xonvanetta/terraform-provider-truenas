package truenas

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/service"
)

func resourceServiceNFS() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceServiceNFSRead,
		CreateContext: resourceServiceNFSCreate,
		UpdateContext: resourceServiceNFSUpdate,
		DeleteContext: resourceServiceNFSDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_on_boot": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"servers": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"udp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_non_root": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"v4": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"v4_v3_owner": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"v4_krb": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bindip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"statd_lockd_log": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"v4_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"v4_krb_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"userd_manage_gids": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceServiceNFSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	service, err := m.(v2.Client).Service().Get(ctx, "nfs")
	if err != nil {
		return diag.Errorf("failed to get service: %s", err)
	}

	d.SetId(strconv.Itoa(service.ID))
	d.Set("start_on_boot", service.Enable)
	d.Set("state", service.State)

	nfs, err := m.(v2.Client).ServiceNFS().Get(ctx)
	if err != nil {
		return diag.Errorf("failed to get nfs service: %s", err)
	}

	d.Set("servers", nfs.Servers)
	d.Set("udp", nfs.UDP)
	d.Set("allow_non_root", nfs.AllowNonroot)
	d.Set("v4", nfs.V4)
	d.Set("v4_v3_owner", nfs.V4V3Owner)
	d.Set("v4_krb", nfs.V4Krb)
	d.Set("bindip", nfs.Bindip)
	d.Set("statd_lockd_log", nfs.RpcstatdPort)
	d.Set("v4_domain", nfs.RpclockdPort)
	d.Set("v4_krb_enabled", nfs.MountdLog)
	d.Set("userd_manage_gids", nfs.StatdLockdLog)

	return nil
}

func resourceServiceNFSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := "nfs"
	err := m.(v2.Client).Service().Start(ctx, name)
	if err != nil {
		return diag.Errorf("failed to start %s service: %s", name, err)
	}

	err = m.(v2.Client).Service().Update(ctx, name, d.Get("start_on_boot").(bool))
	if err != nil {
		return diag.Errorf("failed to update %s service: %s", name, err)
	}

	nfs := &service.NFS{
		Servers:         d.Get("servers").(int),
		UDP:             d.Get("udp").(bool),
		AllowNonroot:    d.Get("allow_non_root").(bool),
		V4:              d.Get("v4").(bool),
		V4V3Owner:       d.Get("v4_v3_owner").(bool),
		V4Krb:           d.Get("v4_krb").(bool),
		StatdLockdLog:   d.Get("statd_lockd_log").(bool),
		V4Domain:        d.Get("v4_domain").(string),
		V4KrbEnabled:    d.Get("v4_krb_enabled").(bool),
		UserdManageGids: d.Get("userd_manage_gids").(bool),
	}

	err = m.(v2.Client).ServiceNFS().Update(ctx, nfs)
	if err != nil {
		return diag.Errorf("failed to update %s service settings: %s", name, err)
	}

	return resourceServiceNFSRead(ctx, d, m)
}

func resourceServiceNFSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := "nfs"

	if d.HasChange("start_on_boot") {
		err := m.(v2.Client).Service().Update(ctx, name, d.Get("start_on_boot").(bool))
		if err != nil {
			return diag.Errorf("failed to update %s service: %s", name, err)
		}
	}

	nfs := &service.NFS{
		Servers:         d.Get("servers").(int),
		UDP:             d.Get("udp").(bool),
		AllowNonroot:    d.Get("allow_non_root").(bool),
		V4:              d.Get("v4").(bool),
		V4V3Owner:       d.Get("v4_v3_owner").(bool),
		V4Krb:           d.Get("v4_krb").(bool),
		StatdLockdLog:   d.Get("statd_lockd_log").(bool),
		V4Domain:        d.Get("v4_domain").(string),
		V4KrbEnabled:    d.Get("v4_krb_enabled").(bool),
		UserdManageGids: d.Get("userd_manage_gids").(bool),
	}

	err := m.(v2.Client).ServiceNFS().Update(ctx, nfs)
	if err != nil {
		return diag.Errorf("failed to update %s service settings: %s", name, err)
	}

	return resourceServiceNFSRead(ctx, d, m)
}

func resourceServiceNFSDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := "nfs"

	err := m.(v2.Client).Service().Stop(ctx, name)
	if err != nil {
		return diag.Errorf("failed to stop %s service: %s", name, err)
	}

	err = m.(v2.Client).Service().Update(ctx, name, false)
	if err != nil {
		return diag.Errorf("failed to disable start on boot %s service: %s", name, err)
	}

	return nil
}
