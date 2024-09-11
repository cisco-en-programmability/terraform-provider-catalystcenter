package catalystcenter

import (
	"context"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceMapsImportStart() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sites.

- Initiates a map archive import of a tar.gz file.  The archive must consist of one xmlDir/MapsImportExport.xml map
descriptor file, and 1 or more images for the map areas nested under /images folder.
`,

		CreateContext: resourceMapsImportStartCreate,
		ReadContext:   resourceMapsImportStartRead,
		DeleteContext: resourceMapsImportStartDelete,
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
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceMapsImportStartCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	response1, err := client.Sites.ImportMapArchiveStartImport()
	if err = d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportMapArchiveStartImport response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceMapsImportStartRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceMapsImportStartDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
