package catalystcenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSitesTelemetrySettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Network Settings.

- Sets telemetry settings for the given site; *null* values indicate that the setting will be inherited from the parent
site.
`,

		CreateContext: resourceSitesTelemetrySettingsCreate,
		ReadContext:   resourceSitesTelemetrySettingsRead,
		UpdateContext: resourceSitesTelemetrySettingsUpdate,
		DeleteContext: resourceSitesTelemetrySettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"application_visibility": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"collector": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"collector_type": &schema.Schema{
													Description: `Collector Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"port": &schema.Schema{
													Description: `Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"enable_on_wired_access_devices": &schema.Schema{
										Description: `Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmp_traps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_trap_servers": &schema.Schema{
										Description: `External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"use_builtin_trap_server": &schema.Schema{
										Description: `Enable this server as a destination server for SNMP traps and messages from your network
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"syslogs": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_syslog_servers": &schema.Schema{
										Description: `External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"use_builtin_syslog_server": &schema.Schema{
										Description: `Enable this server as a destination server for syslog messages.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"wired_data_collection": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_wired_data_collectio": &schema.Schema{
										Description: `Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"wireless_telemetry": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_wireless_telemetry": &schema.Schema{
										Description: `Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application_visibility": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"collector": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"collector_type": &schema.Schema{
													Description: `Collector Type`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"port": &schema.Schema{
													Description: `Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"enable_on_wired_access_devices": &schema.Schema{
										Description: `Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Site Id, retrievable from the *id* attribute in */dna/intent/api/v1/sites*
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"snmp_traps": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_trap_servers": &schema.Schema{
										Description: `External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"use_builtin_trap_server": &schema.Schema{
										Description: `Enable this server as a destination server for SNMP traps and messages from your network
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"syslogs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_syslog_servers": &schema.Schema{
										Description: `External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"use_builtin_syslog_server": &schema.Schema{
										Description: `Enable this server as a destination server for syslog messages.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"wired_data_collection": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_wired_data_collectio": &schema.Schema{
										Description: `Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"wireless_telemetry": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_wireless_telemetry": &schema.Schema{
										Description: `Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
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

func resourceSitesTelemetrySettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSitesTelemetrySettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTelemetrySettingsForASite")
		vvID := vID
		queryParams1 := catalystcentersdkgo.RetrieveTelemetrySettingsForASiteV1QueryParams{}

		response1, restyResp1, err := client.NetworkSettings.RetrieveTelemetrySettingsForASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveTelemetrySettingsForASiteV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTelemetrySettingsForASite response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSitesTelemetrySettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.SetTelemetrySettingsForASite(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SetTelemetrySettingsForASite", err, restyResp1.String(),
					"Failure at SetTelemetrySettingsForASite, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetTelemetrySettingsForASite", err,
				"Failure at SetTelemetrySettingsForASite, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing SetTelemetrySettingsForASite", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing SetTelemetrySettingsForASite", err1))
				return diags
			}
		}

	}

	return resourceSitesTelemetrySettingsRead(ctx, d, m)
}

func resourceSitesTelemetrySettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SitesTelemetrySettings", err, "Delete method is not supported",
		"Failure at SitesTelemetrySettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1 {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wired_data_collection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wired_data_collection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wired_data_collection")))) {
		request.WiredDataCollection = expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1WiredDataCollection(ctx, key+".wired_data_collection.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wireless_telemetry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wireless_telemetry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wireless_telemetry")))) {
		request.WirelessTelemetry = expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1WirelessTelemetry(ctx, key+".wireless_telemetry.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_traps")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_traps")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_traps")))) {
		request.SNMPTraps = expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1SNMPTraps(ctx, key+".snmp_traps.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".syslogs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".syslogs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".syslogs")))) {
		request.Syslogs = expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1Syslogs(ctx, key+".syslogs.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_visibility")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_visibility")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_visibility")))) {
		request.ApplicationVisibility = expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1ApplicationVisibility(ctx, key+".application_visibility.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1WiredDataCollection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1WiredDataCollection {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1WiredDataCollection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_wired_data_collectio")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_wired_data_collectio")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_wired_data_collectio")))) {
		request.EnableWiredDataCollectio = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1WirelessTelemetry(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1WirelessTelemetry {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1WirelessTelemetry{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_wireless_telemetry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_wireless_telemetry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_wireless_telemetry")))) {
		request.EnableWirelessTelemetry = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1SNMPTraps(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1SNMPTraps {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1SNMPTraps{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_builtin_trap_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_builtin_trap_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_builtin_trap_server")))) {
		request.UseBuiltinTrapServer = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_trap_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_trap_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_trap_servers")))) {
		request.ExternalTrapServers = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1Syslogs(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1Syslogs {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1Syslogs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_builtin_syslog_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_builtin_syslog_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_builtin_syslog_server")))) {
		request.UseBuiltinSyslogServer = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_syslog_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_syslog_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_syslog_servers")))) {
		request.ExternalSyslogServers = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1ApplicationVisibility(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibility {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibility{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".collector")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".collector")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".collector")))) {
		request.Collector = expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1ApplicationVisibilityCollector(ctx, key+".collector.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_on_wired_access_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_on_wired_access_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_on_wired_access_devices")))) {
		request.EnableOnWiredAccessDevices = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSitesTelemetrySettingsSetTelemetrySettingsForASiteV1ApplicationVisibilityCollector(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibilityCollector {
	request := catalystcentersdkgo.RequestNetworkSettingsSetTelemetrySettingsForASiteV1ApplicationVisibilityCollector{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".collector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".collector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".collector_type")))) {
		request.CollectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
