package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceCredential() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get device credential details. This data source has been deprecated and will not be available in a Cisco DNA
Center release after August 1st 2024 23:59:59 GMT. Please refer new Intent API : Get All Global Credentials V2
`,

		ReadContext: dataSourceDeviceCredentialRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site id to retrieve the credential details associated with the site.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"enable_password": &schema.Schema{
										Description: `Enable Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},

									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"http_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"secure": &schema.Schema{
										Description: `Secure`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"http_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"secure": &schema.Schema{
										Description: `Secure`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmp_v2_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"read_community": &schema.Schema{
										Description: `Read Community`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmp_v2_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"write_community": &schema.Schema{
										Description: `Write Community`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmp_v3": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_password": &schema.Schema{
										Description: `Auth Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"privacy_password": &schema.Schema{
										Description: `Privacy Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"privacy_type": &schema.Schema{
										Description: `Privacy Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snmp_mode": &schema.Schema{
										Description: `Snmp Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceDeviceCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCredentialDetailsV1")
		queryParams1 := catalystcentersdkgo.GetDeviceCredentialDetailsV1QueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.GetDeviceCredentialDetailsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceCredentialDetailsV1", err,
				"Failure at GetDeviceCredentialDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetDeviceCredentialDetailsV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCredentialDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["snmp_v3"] = flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemSNMPV3(item.SNMPV3)
	respItem["http_read"] = flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemHTTPRead(item.HTTPRead)
	respItem["http_write"] = flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemHTTPWrite(item.HTTPWrite)
	respItem["snmp_v2_write"] = flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemSNMPV2Write(item.SNMPV2Write)
	respItem["snmp_v2_read"] = flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemSNMPV2Read(item.SNMPV2Read)
	respItem["cli"] = flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemCli(item.Cli)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemSNMPV3(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV3) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["username"] = item.Username
		respItem["auth_password"] = item.AuthPassword
		respItem["auth_type"] = item.AuthType
		respItem["privacy_password"] = item.PrivacyPassword
		respItem["privacy_type"] = item.PrivacyType
		respItem["snmp_mode"] = item.SNMPMode
		respItem["comments"] = item.Comments
		respItem["description"] = item.Description
		respItem["credential_type"] = item.CredentialType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemHTTPRead(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1HTTPRead) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["secure"] = item.Secure
		respItem["username"] = item.Username
		respItem["password"] = item.Password
		respItem["port"] = item.Port
		respItem["comments"] = item.Comments
		respItem["description"] = item.Description
		respItem["credential_type"] = item.CredentialType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemHTTPWrite(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1HTTPWrite) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["secure"] = item.Secure
		respItem["username"] = item.Username
		respItem["password"] = item.Password
		respItem["port"] = item.Port
		respItem["comments"] = item.Comments
		respItem["description"] = item.Description
		respItem["credential_type"] = item.CredentialType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemSNMPV2Write(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV2Write) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["write_community"] = item.WriteCommunity
		respItem["comments"] = item.Comments
		respItem["description"] = item.Description
		respItem["credential_type"] = item.CredentialType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemSNMPV2Read(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1SNMPV2Read) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["read_community"] = item.ReadCommunity
		respItem["comments"] = item.Comments
		respItem["description"] = item.Description
		respItem["credential_type"] = item.CredentialType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetDeviceCredentialDetailsV1ItemCli(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialDetailsV1Cli) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["username"] = item.Username
		respItem["enable_password"] = item.EnablePassword
		respItem["password"] = item.Password
		respItem["comments"] = item.Comments
		respItem["description"] = item.Description
		respItem["credential_type"] = item.CredentialType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
