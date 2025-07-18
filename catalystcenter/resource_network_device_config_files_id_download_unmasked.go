package catalystcenter

import (
	"context"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkDeviceConfigFilesIDDownloadUnmasked() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Archive.

- Download the unmasked (raw) device configuration by providing the file **id** and a **password**. The response will be
a password-protected zip file containing the unmasked configuration. Password must contain a minimum of 8 characters,
one lowercase letter, one uppercase letter, one number, one special character (**-=[];,./~!@#$%^&*()_+{}|:?**). It may
not contain white space or the characters **<>**.
`,

		CreateContext: resourceNetworkDeviceConfigFilesIDDownloadUnmaskedCreate,
		ReadContext:   resourceNetworkDeviceConfigFilesIDDownloadUnmaskedRead,
		DeleteContext: resourceNetworkDeviceConfigFilesIDDownloadUnmaskedDelete,
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
						"password": &schema.Schema{
							Description: `Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.
`,
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceConfigFilesIDDownloadUnmaskedCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)
	request1 := expandRequestNetworkDeviceConfigFilesIDDownloadUnmaskedDownloadUnmaskedrawDeviceConfigurationAsZIP(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.ConfigurationArchive.DownloadUnmaskedrawDeviceConfigurationAsZIP(vvID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

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
func resourceNetworkDeviceConfigFilesIDDownloadUnmaskedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDeviceConfigFilesIDDownloadUnmaskedDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDeviceConfigFilesIDDownloadUnmaskedDownloadUnmaskedrawDeviceConfigurationAsZIP(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP {
	request := catalystcentersdkgo.RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	return &request
}
