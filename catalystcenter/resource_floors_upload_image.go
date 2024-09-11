package catalystcenter

import (
	"context"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceFloorsUploadImage() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Uploads floor image.
`,

		CreateContext: resourceFloorsUploadImageCreate,
		ReadContext:   resourceFloorsUploadImageRead,
		DeleteContext: resourceFloorsUploadImageDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: `id path parameter. Floor Id
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceFloorsUploadImageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)

	response1, err := client.SiteDesign.UploadsFloorImage(vvID)
	if err = d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UploadsFloorImage response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceFloorsUploadImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFloorsUploadImageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
