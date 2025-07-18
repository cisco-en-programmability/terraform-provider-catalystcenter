package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source retrieves the latest system events
`,

		ReadContext: dataSourceSystemHealthRead,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: `domain query parameter. Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Specifies the maximum number of system health events to return per page. Must be an integer between 1 and 50, inclusive
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point for the list of system health events to return. Must be an integer greater than or equal to 0
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"subdomain": &schema.Schema{
				Description: `subdomain query parameter. Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary": &schema.Schema{
				Description: `summary query parameter. Fetch the latest high severity event
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cimcaddress": &schema.Schema{
							Description: `List of configured cimc addresse(s)
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"health_events": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Details of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"domain": &schema.Schema{
										Description: `Domain of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"hostname": &schema.Schema{
										Description: `Hostname of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance": &schema.Schema{
										Description: `Instance of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"severity": &schema.Schema{
										Description: `Severity of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `State of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Event status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sub_domain": &schema.Schema{
										Description: `Sub domain of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"timestamp": &schema.Schema{
										Description: `Time of the event occurance
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"host_name": &schema.Schema{
							Description: `Cluster name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `API version
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSummary, okSummary := d.GetOk("summary")
	vDomain, okDomain := d.GetOk("domain")
	vSubdomain, okSubdomain := d.GetOk("subdomain")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SystemHealthAPI")
		queryParams1 := catalystcentersdkgo.SystemHealthAPIQueryParams{}

		if okSummary {
			queryParams1.Summary = vSummary.(bool)
		}
		if okDomain {
			queryParams1.Domain = vDomain.(string)
		}
		if okSubdomain {
			queryParams1.Subdomain = vSubdomain.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.HealthAndPerformance.SystemHealthAPI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SystemHealthAPI", err,
				"Failure at SystemHealthAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SystemHealthAPI", err,
				"Failure at SystemHealthAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemHealthAPIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemHealthAPI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemHealthAPIItem(item *catalystcentersdkgo.ResponseHealthAndPerformanceSystemHealthAPI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["health_events"] = flattenHealthAndPerformanceSystemHealthAPIItemHealthEvents(item.HealthEvents)
	respItem["version"] = item.Version
	respItem["host_name"] = item.HostName
	respItem["cimcaddress"] = item.Cimcaddress
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceSystemHealthAPIItemHealthEvents(items *[]catalystcentersdkgo.ResponseHealthAndPerformanceSystemHealthAPIHealthEvents) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["severity"] = item.Severity
		respItem["hostname"] = item.Hostname
		respItem["instance"] = item.Instance
		respItem["sub_domain"] = item.SubDomain
		respItem["domain"] = item.Domain
		respItem["description"] = item.Description
		respItem["state"] = item.State
		respItem["timestamp"] = item.Timestamp
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
