package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceMapsImportPerform() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sites.

- For a previously initatied import, approves the import to be performed, accepting that data loss may occur.  A Map
import will fully replace existing Maps data for the site(s) defined in the archive. The Map Archive Import Status API
/maps/import/${contextUuid}/status should always be checked to validate the pre-import validation output prior to
performing the import.
`,

		CreateContext: resourceMapsImportPerformCreate,
		ReadContext:   resourceMapsImportPerformRead,
		DeleteContext: resourceMapsImportPerformDelete,
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
						"import_context_uuid": &schema.Schema{
							Description: `importContextUuid path parameter. The unique import context UUID given by a previous call of Start Import API
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

func resourceMapsImportPerformCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vImportContextUUID := resourceItem["import_context_uuid"]

	vvImportContextUUID := vImportContextUUID.(string)

	_, response1, err := client.Sites.ImportMapArchivePerformImport(vvImportContextUUID)

	if err != nil || response1 == nil {
		diags = append(diags, diagError(
			"Failure when executing ImportMapArchivePerformImport", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//Analizar verificacion.

	vItem1 := response1.String()
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportMapArchivePerformImport response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceMapsImportPerformRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceMapsImportPerformDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
