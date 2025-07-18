package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAuthenticationPolicyServers() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on System Settings.

- API to get Authentication and Policy Servers
`,

		ReadContext: dataSourceAuthenticationPolicyServersRead,
		Schema: map[string]*schema.Schema{
			"is_ise_enabled": &schema.Schema{
				Description: `isIseEnabled query parameter. Valid values are : true, false
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"role": &schema.Schema{
				Description: `role query parameter. Authentication and Policy Server Role (Example: primary, secondary)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Description: `state query parameter. Valid values are: ACTIVE, DELETED, FAILED, INACTIVE, INPROGRESS, RBAC-FAILURE, RBAC-SUCCESS
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accounting_port": &schema.Schema{
							Description: `Accounting port of RADIUS server (readonly). The range is from 1 to 65535. Default is 1813. E.g. 1813
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"authentication_port": &schema.Schema{
							Description: `Authentication port of RADIUS server (readonly). The range is from 1 to 65535. Default is 1812. E.g. 1812
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"cisco_ise_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description about the Cisco ISE server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"external_cisco_ise_ip_addr_dtos": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"external_cisco_ise_ip_addresses": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"external_ip_address": &schema.Schema{
																Description: `External IP Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"failure_reason": &schema.Schema{
										Description: `Reason for integration failure between CATALYST and the ISE server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fqdn": &schema.Schema{
										Description: `Fully-qualified domain name of the Cisco ISE server (readonly). E.g. xi-62.my.com
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Internal record identifier
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_address": &schema.Schema{
										Description: `IP Address of the Cisco ISE Server (readonly)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": &schema.Schema{
										Description: `Password of the Cisco ISE server. For security reasons the value will always be empty
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},

									"role": &schema.Schema{
										Description: `Role of the Cisco ISE server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sshkey": &schema.Schema{
										Description: `SSH key of the Cisco ISE server. For security reasons the value will always be empty
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"subscriber_name": &schema.Schema{
										Description: `Subscriber name of the Cisco ISE server (readonly). E.g. pxgrid_client_1662589467
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"trust_state": &schema.Schema{
										Description: `Trust State between CATALYST and the ISE server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `Type (Example: ISE)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"user_name": &schema.Schema{
										Description: `User name of the Cisco ISE server
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"encryption_key": &schema.Schema{
							Description: `Encryption key used to encrypt shared secret
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"encryption_scheme": &schema.Schema{
							Description: `Type of encryption scheme for additional security (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Internal record identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Description: `IP address of authentication and policy server (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_ise_enabled": &schema.Schema{
							Description: `Value true for Cisco ISE Server (readonly). Default value is false
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ise_enabled": &schema.Schema{
							Description: `Value true for Cisco ISE Server (readonly). Default value is false
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"message_key": &schema.Schema{
							Description: `Encryption key used to encrypt shared secret (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"multi_dnac_enabled": &schema.Schema{
							Description: `Internal use only
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"port": &schema.Schema{
							Description: `Port of TACACS server (readonly). The range is from 1 to 65535. Default is 49.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"protocol": &schema.Schema{
							Description: `Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pxgrid_enabled": &schema.Schema{
							Description: `Value true for enable, false for disable. Default value is true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rbac_uuid": &schema.Schema{
							Description: `Internal use only
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"retries": &schema.Schema{
							Description: `Number of communication retries between devices and authentication and policy server. The range is from 1 to 3. Default is 3
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"role": &schema.Schema{
							Description: `Role of authentication and policy server (readonly). E.g. primary, secondary
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"shared_secret": &schema.Schema{
							Description: `Shared secret between devices and authentication and policy server (readonly)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"state": &schema.Schema{
							Description: `State of authentication and policy server
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"timeout_seconds": &schema.Schema{
							Description: `Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"use_dnac_cert_for_pxgrid": &schema.Schema{
							Description: `Value true to use CATALYST certificate for Pxgrid. Default value is false
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAuthenticationPolicyServersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vIsIseEnabled, okIsIseEnabled := d.GetOk("is_ise_enabled")
	vState, okState := d.GetOk("state")
	vRole, okRole := d.GetOk("role")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAuthenticationAndPolicyServers")
		queryParams1 := catalystcentersdkgo.GetAuthenticationAndPolicyServersQueryParams{}

		if okIsIseEnabled {
			queryParams1.IsIseEnabled = vIsIseEnabled.(bool)
		}
		if okState {
			queryParams1.State = vState.(string)
		}
		if okRole {
			queryParams1.Role = vRole.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.SystemSettings.GetAuthenticationAndPolicyServers(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAuthenticationAndPolicyServers", err,
				"Failure at GetAuthenticationAndPolicyServers, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAuthenticationAndPolicyServers", err,
				"Failure at GetAuthenticationAndPolicyServers, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSystemSettingsGetAuthenticationAndPolicyServersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthenticationAndPolicyServers response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSystemSettingsGetAuthenticationAndPolicyServersItems(items *[]catalystcentersdkgo.ResponseSystemSettingsGetAuthenticationAndPolicyServersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_address"] = item.IPAddress
		respItem["shared_secret"] = item.SharedSecret
		respItem["protocol"] = item.Protocol
		respItem["role"] = item.Role
		respItem["port"] = item.Port
		respItem["authentication_port"] = item.AuthenticationPort
		respItem["accounting_port"] = item.AccountingPort
		respItem["retries"] = item.Retries
		respItem["timeout_seconds"] = item.TimeoutSeconds
		respItem["is_ise_enabled"] = boolPtrToString(item.IsIseEnabled)
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["state"] = item.State
		respItem["cisco_ise_dtos"] = flattenSystemSettingsGetAuthenticationAndPolicyServersItemsCiscoIseDtos(item.CiscoIseDtos)
		respItem["encryption_scheme"] = item.EncryptionScheme
		respItem["message_key"] = item.MessageKey
		respItem["encryption_key"] = item.EncryptionKey
		respItem["use_dnac_cert_for_pxgrid"] = boolPtrToString(item.UseDnacCertForPxgrid)
		respItem["ise_enabled"] = boolPtrToString(item.IseEnabled)
		respItem["pxgrid_enabled"] = boolPtrToString(item.PxgridEnabled)
		respItem["rbac_uuid"] = item.RbacUUID
		respItem["multi_dnac_enabled"] = boolPtrToString(item.MultiDnacEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSystemSettingsGetAuthenticationAndPolicyServersItemsCiscoIseDtos(items *[]catalystcentersdkgo.ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtos) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["subscriber_name"] = item.SubscriberName
		respItem["description"] = item.Description
		respItem["password"] = item.Password
		respItem["user_name"] = item.UserName
		respItem["fqdn"] = item.Fqdn
		respItem["ip_address"] = item.IPAddress
		respItem["trust_state"] = item.TrustState
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["sshkey"] = item.SSHkey
		respItem["type"] = item.Type
		respItem["failure_reason"] = item.FailureReason
		respItem["role"] = item.Role
		respItem["external_cisco_ise_ip_addr_dtos"] = flattenSystemSettingsGetAuthenticationAndPolicyServersItemsCiscoIseDtosExternalCiscoIseIPAddrDtos(item.ExternalCiscoIseIPAddrDtos)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSystemSettingsGetAuthenticationAndPolicyServersItemsCiscoIseDtosExternalCiscoIseIPAddrDtos(item *catalystcentersdkgo.ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtosExternalCiscoIseIPAddrDtos) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["external_cisco_ise_ip_addresses"] = flattenSystemSettingsGetAuthenticationAndPolicyServersItemsCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses(item.ExternalCiscoIseIPAddresses)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSystemSettingsGetAuthenticationAndPolicyServersItemsCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses(items *[]catalystcentersdkgo.ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["external_ip_address"] = item.ExternalIPAddress
		respItems = append(respItems, respItem)
	}
	return respItems
}
