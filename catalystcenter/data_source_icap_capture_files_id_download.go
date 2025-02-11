package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapCaptureFilesIDDownload() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Downloads a specific ICAP packet capture file. For detailed information about the usage of the API, please refer to
the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapCaptureFilesIDDownloadRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"object": &schema.Schema{
							Description: `object`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapCaptureFilesIDDownloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DownloadsASpecificICapPacketCaptureFileV1")
		vvID := vID.(string)

		headerParams1 := catalystcentersdkgo.DownloadsASpecificICapPacketCaptureFileV1HeaderParams{}

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Sensors.DownloadsASpecificICapPacketCaptureFileV1(vvID, &headerParams1)

		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 DownloadsASpecificICapPacketCaptureFileV1", err,
				"Failure at DownloadsASpecificICapPacketCaptureFileV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(response1))

		vvDirpath := d.Get("dirpath").(string)
		err = response1.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}
