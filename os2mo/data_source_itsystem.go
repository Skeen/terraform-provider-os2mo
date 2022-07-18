package os2mo

import (
	"context"

	"github.com/Khan/genqlient/graphql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcesItsystems() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcesReadItsystems,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcesReadItsystems(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resp, err := getITSystems(meta.(graphql.Client))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("uuid", resp.Itsystems[0].Uuid.String())
	d.Set("system_type", resp.Itsystems[0].System_type)
	d.Set("type", resp.Itsystems[0].Type)
	d.Set("user_key", resp.Itsystems[0].User_key)
	d.Set("name", resp.Itsystems[0].Name)
	/*
		exchange := make([]map[string]interface{}, 1)
		e := make(map[string]interface{})
		e["type"] = exchangeSettings.Type
		e["durable"] = exchangeSettings.Durable
		e["auto_delete"] = exchangeSettings.AutoDelete
		e["arguments"] = exchangeSettings.Arguments
		exchange[0] = e
	*/
	d.SetId(resp.Itsystems[0].Uuid.String())

	var diags diag.Diagnostics
	return diags
}
