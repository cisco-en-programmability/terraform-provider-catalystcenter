package catalystcenter

import (
	"context"

	"time"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- API to create a network for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and EndPoint AAA, and/or DNS center server
settings.
`,

		CreateContext: resourceNetworkCreateCreate,
		ReadContext:   resourceNetworkCreateRead,
		DeleteContext: resourceNetworkCreateDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site id to which site details to associate with the network settings.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"client_and_endpoint_aaa": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address": &schema.Schema{
													Description: `IP address for ISE serve (eg: 1.1.1.4)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"network": &schema.Schema{
													Description: `IP address for AAA or ISE server (eg: 2.2.2.1)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol for AAA or ISE serve (eg: RADIUS)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"servers": &schema.Schema{
													Description: `Server type AAA or ISE server (eg: AAA)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"shared_secret": &schema.Schema{
													Description: `Shared secret for ISE server
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"dhcp_server": &schema.Schema{
										Description: `DHCP Server IP (eg: 1.1.1.1)
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"dns_server": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"domain_name": &schema.Schema{
													Description: `Domain Name of DHCP (eg; cisco)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"primary_ip_address": &schema.Schema{
													Description: `Primary IP Address for DHCP (eg: 2.2.2.2)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"secondary_ip_address": &schema.Schema{
													Description: `Secondary IP Address for DHCP (eg: 3.3.3.3)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"message_of_theday": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"banner_message": &schema.Schema{
													Description: `Massage for Banner message (eg; Good day)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"retain_existing_banner": &schema.Schema{
													Description: `Retain existing Banner Message (eg: "true" or "false")
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"netflowcollector": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address": &schema.Schema{
													Description: `IP Address for NetFlow collector (eg: 3.3.3.1)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"port": &schema.Schema{
													Description: `Port for NetFlow Collector (eg; 443)
`,
													Type:     schema.TypeFloat,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"network_aaa": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address": &schema.Schema{
													Description: `IP address for AAA and ISE server (eg: 1.1.1.1)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"network": &schema.Schema{
													Description: `IP Address for AAA or ISE server (eg: 2.2.2.2)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol for AAA or ISE serve (eg: RADIUS)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"servers": &schema.Schema{
													Description: `Server type for AAA Network (eg: AAA)
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"shared_secret": &schema.Schema{
													Description: `Shared secret for ISE Server
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"ntp_server": &schema.Schema{
										Description: `IP address for NTP server (eg: 1.1.1.2)
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"snmp_server": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"configure_dnac_ip": &schema.Schema{
													Description: `Configuration CATALYST IP for SNMP Server (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													ForceNew:     true,
													Computed:     true,
												},
												"ip_addresses": &schema.Schema{
													Description: `IP Address for SNMP Server (eg: 4.4.4.1)
`,
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"syslog_server": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"configure_dnac_ip": &schema.Schema{
													Description: `Configuration CATALYST IP for syslog server (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													ForceNew:     true,
													Computed:     true,
												},
												"ip_addresses": &schema.Schema{
													Description: `IP Address for syslog server (eg: 4.4.4.4)
`,
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"timezone": &schema.Schema{
										Description: `Input for time zone (eg: Africa/Abidjan)
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceNetworkCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vSiteID := resourceItem["site_id"]

	vPersistbapioutput := resourceItem["persistbapioutput"]

	vvSiteID := vSiteID.(string)
	request1 := expandRequestNetworkCreateCreateNetwork(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.CreateNetworkHeaderParams{}

	headerParams1.Persistbapioutput = vPersistbapioutput.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.NetworkSettings.CreateNetwork(vvSiteID, request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetwork", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateNetwork", err,
				"Failure at CreateNetwork execution", bapiError))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenNetworkSettingsCreateNetworkItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreateNetwork response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceNetworkCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkCreateCreateNetwork(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetwork {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetwork{}
	request.Settings = expandRequestNetworkCreateCreateNetworkSettings(ctx, key, d)
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettings {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server")))) {
		request.DhcpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server")))) {
		request.DNSServer = expandRequestNetworkCreateCreateNetworkSettingsDNSServer(ctx, key+".dns_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".syslog_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".syslog_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".syslog_server")))) {
		request.SyslogServer = expandRequestNetworkCreateCreateNetworkSettingsSyslogServer(ctx, key+".syslog_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_server")))) {
		request.SNMPServer = expandRequestNetworkCreateCreateNetworkSettingsSNMPServer(ctx, key+".snmp_server.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netflowcollector")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netflowcollector")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netflowcollector")))) {
		request.Netflowcollector = expandRequestNetworkCreateCreateNetworkSettingsNetflowcollector(ctx, key+".netflowcollector.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ntp_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ntp_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ntp_server")))) {
		request.NtpServer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timezone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timezone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timezone")))) {
		request.Timezone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".message_of_theday")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".message_of_theday")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".message_of_theday")))) {
		request.MessageOfTheday = expandRequestNetworkCreateCreateNetworkSettingsMessageOfTheday(ctx, key+".message_of_theday.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_aaa")))) {
		request.NetworkAAA = expandRequestNetworkCreateCreateNetworkSettingsNetworkAAA(ctx, key+".network_aaa.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_and_endpoint_aaa")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_and_endpoint_aaa")))) {
		request.ClientAndEndpointAAA = expandRequestNetworkCreateCreateNetworkSettingsClientAndEndpointAAA(ctx, key+".client_and_endpoint_aaa.0", d)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsDNSServer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsDNSServer {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsDNSServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain_name")))) {
		request.DomainName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_ip_address")))) {
		request.PrimaryIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_ip_address")))) {
		request.SecondaryIPAddress = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsSyslogServer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSyslogServer {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSyslogServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addresses")))) {
		request.IPAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_dnac_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_dnac_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_dnac_ip")))) {
		request.ConfigureDnacIP = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsSNMPServer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSNMPServer {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsSNMPServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addresses")))) {
		request.IPAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_dnac_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_dnac_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_dnac_ip")))) {
		request.ConfigureDnacIP = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsNetflowcollector(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetflowcollector {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetflowcollector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsMessageOfTheday(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_message")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_message")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_message")))) {
		request.BannerMessage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retain_existing_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retain_existing_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retain_existing_banner")))) {
		request.RetainExistingBanner = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsNetworkAAA(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetworkAAA {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsNetworkAAA{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network")))) {
		request.Network = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkCreateCreateNetworkSettingsClientAndEndpointAAA(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network")))) {
		request.Network = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	return &request
}

func flattenNetworkSettingsCreateNetworkItem(item *catalystcentersdkgo.ResponseNetworkSettingsCreateNetwork) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
