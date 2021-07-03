package sharing

import (
	"context"
	"strconv"
	"time"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/sharing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"
)

func DataNFSResource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataNFSRead,
		Schema: map[string]*schema.Schema{
			"truenas_sharing_nfs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"paths": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"aliases": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"networks": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"hosts": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"security": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alldirs": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ro": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"quiet": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"maproot_user": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maproot_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mapall_user": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mapall_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenNFSData(list []*sharing.NFS) []interface{} {
	if list == nil {
		return make([]interface{}, 0)
	}

	l := make([]interface{}, len(list), len(list))
	for i, nfs := range list {
		n := make(map[string]interface{})
		n["id"] = strconv.Itoa(nfs.Id)
		n["paths"] = nfs.Paths
		n["enabled"] = nfs.Enabled
		n["paths"] = nfs.Paths
		n["aliases"] = nfs.Aliases
		n["comment"] = nfs.Comment
		n["networks"] = nfs.Networks
		n["hosts"] = nfs.Hosts
		n["alldirs"] = nfs.Alldirs
		n["ro"] = nfs.Ro
		n["quiet"] = nfs.Quiet
		n["maproot_user"] = nfs.MaprootUser
		n["maproot_group"] = nfs.MaprootGroup
		n["mapall_user"] = nfs.MapallUser
		n["mapall_group"] = nfs.MapallGroup
		n["security"] = nfs.Security
		n["enabled"] = nfs.Enabled
		l[i] = n
	}
	return l
}

func dataNFSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	list, err := m.(v2.Client).SharingNFS().List(ctx)
	if err != nil {
		return diag.Errorf("failed to get nfs list: %s", err)
	}

	err = d.Set("truenas_sharing_nfs", flattenNFSData(list))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(time.Now().String())

	return nil
}
