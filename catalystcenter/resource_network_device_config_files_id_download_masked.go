package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkDeviceConfigFilesIDDownloadMasked() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Archive.

- Download the masked (sanitized) device configuration by providing the file **id**.
`,

		CreateContext: resourceNetworkDeviceConfigFilesIDDownloadMaskedCreate,
		ReadContext:   resourceNetworkDeviceConfigFilesIDDownloadMaskedRead,
		DeleteContext: resourceNetworkDeviceConfigFilesIDDownloadMaskedDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
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
							Description: `id path parameter. The value of **id** can be obtained from the response of API **/dna/intent/api/v1/networkDeviceConfigFiles**
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

func resourceNetworkDeviceConfigFilesIDDownloadMaskedCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.ConfigurationArchive.DownloadMaskedDeviceConfiguration(vvID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceNetworkDeviceConfigFilesIDDownloadMaskedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDeviceConfigFilesIDDownloadMaskedDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
