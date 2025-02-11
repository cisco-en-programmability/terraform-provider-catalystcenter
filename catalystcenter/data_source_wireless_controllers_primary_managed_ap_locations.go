package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersPrimaryManagedApLocations() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieves all the details of Primary Managed AP locations associated with the specific Wireless Controller.
`,

		ReadContext: dataSourceWirelessControllersPrimaryManagedApLocationsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"managed_ap_locations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"site_id": &schema.Schema{
										Description: `The site id of the managed ap location.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"site_name_hierarchy": &schema.Schema{
										Description: `The site name hierarchy of the managed ap location.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessControllersPrimaryManagedApLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPrimaryManagedApLocationsForSpecificWirelessControllerV1")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := catalystcentersdkgo.GetPrimaryManagedApLocationsForSpecificWirelessControllerV1QueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.Wireless.GetPrimaryManagedApLocationsForSpecificWirelessControllerV1(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPrimaryManagedApLocationsForSpecificWirelessControllerV1", err,
				"Failure at GetPrimaryManagedApLocationsForSpecificWirelessControllerV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPrimaryManagedApLocationsForSpecificWirelessControllerV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1Items(items *[]catalystcentersdkgo.ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["managed_ap_locations"] = flattenWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1ItemsManagedApLocations(item.ManagedApLocations)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1ItemsManagedApLocations(items *[]catalystcentersdkgo.ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerV1ResponseManagedApLocations) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_id"] = item.SiteID
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItems = append(respItems, respItem)
	}
	return respItems
}
