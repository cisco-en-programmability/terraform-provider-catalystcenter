package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsersExternalServers() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get external users authentication servers.
`,

		ReadContext: dataSourceUsersExternalServersRead,
		Schema: map[string]*schema.Schema{
			"invoke_source": &schema.Schema{
				Description: `invokeSource query parameter. The source that invokes this API. The value of this query parameter must be set to "external".
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_servers": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_attribute": &schema.Schema{
										Description: `Aaa Attribute`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"accounting_port": &schema.Schema{
										Description: `RADIUS server accounting requests port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"authentication_port": &schema.Schema{
										Description: `RADIUS server authorization requests port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"retries": &schema.Schema{
										Description: `Retries`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"role": &schema.Schema{
										Description: `Role of AAA server, primary or secondary server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"server_id": &schema.Schema{
										Description: `Server Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"server_ip": &schema.Schema{
										Description: `Server Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"socket_timeout": &schema.Schema{
										Description: `Timeout in seconds
`,
										Type:     schema.TypeInt,
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

func dataSourceUsersExternalServersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vInvokeSource := d.Get("invoke_source")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetExternalAuthenticationServersAPIV1")
		queryParams1 := catalystcentersdkgo.GetExternalAuthenticationServersAPIV1QueryParams{}

		queryParams1.InvokeSource = vInvokeSource.(string)

		response1, restyResp1, err := client.UserandRoles.GetExternalAuthenticationServersAPIV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetExternalAuthenticationServersAPIV1", err,
				"Failure at GetExternalAuthenticationServersAPIV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetExternalAuthenticationServersAPIV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalAuthenticationServersAPIV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetExternalAuthenticationServersAPIV1Item(item *catalystcentersdkgo.ResponseUserandRolesGetExternalAuthenticationServersAPIV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_servers"] = flattenUserandRolesGetExternalAuthenticationServersAPIV1ItemAAAServers(item.AAAServers)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetExternalAuthenticationServersAPIV1ItemAAAServers(items *[]catalystcentersdkgo.ResponseUserandRolesGetExternalAuthenticationServersAPIV1ResponseAAAServers) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["accounting_port"] = item.AccountingPort
		respItem["retries"] = item.Retries
		respItem["protocol"] = item.Protocol
		respItem["socket_timeout"] = item.SocketTimeout
		respItem["server_ip"] = item.ServerIP
		respItem["shared_secret"] = item.SharedSecret
		respItem["server_id"] = item.ServerID
		respItem["authentication_port"] = item.AuthenticationPort
		respItem["aaa_attribute"] = item.AAAAttribute
		respItem["role"] = item.Role
		respItems = append(respItems, respItem)
	}
	return respItems
}
