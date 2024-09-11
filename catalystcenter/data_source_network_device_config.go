package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Archive.

- Returns the historical device configurations (running configuration , startup configuration , vlan if applicable) by
specified criteria
`,

		ReadContext: dataSourceNetworkDeviceConfigRead,
		Schema: map[string]*schema.Schema{
			"created_by": &schema.Schema{
				Description: `createdBy query parameter. Comma separated values for createdBy SCHEDULED, USER, CONFIG_CHANGE_EVENT, SCHEDULED_FIRST_TIME, DR_CALL_BACK, PRE_DEPLOY
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_time": &schema.Schema{
				Description: `createdTime query parameter. Supported with logical filters GT,GTE,LT,LTE & BT : time in milliseconds (epoc format)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. comma separated device id for example cf35b0a1-407f-412f-b2f4-f0c3156695f9,aaa38191-0c22-4158-befd-779a09d7cec1 . if device id is not provided it will fetch for all devices
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_type": &schema.Schema{
				Description: `fileType query parameter. Config File Type can be RUNNINGCONFIG or STARTUPCONFIG
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_id": &schema.Schema{
							Description: `UUID of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_name": &schema.Schema{
							Description: `Hostname of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Description: `IP address of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"versions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_change_type": &schema.Schema{
										Description: `Source of configuration change (OUT_OF_BAND - Change was made outside Catalyst Center, IN_BAND - Change was made from Catalyst Center, NOT_APPLICABLE - System Triggered)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"created_by": &schema.Schema{
										Description: `Reason for archive collection (CONFIG_CHANGE_EVENT - Syslog event based colletion, SCHEDULED - Weekly scheduled collection, SCHEDULED_FIRST_TIME-First Time Managed collection,DR_CALL_BACK- Post Disaster Recovery collection, USER- On Demand Trigger, PRE_DEPLOY- Predeploy collection) 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"created_time": &schema.Schema{
										Description: `Source of configuration change (OUT_OF_BAND - Change was made outside Catalyst Center, IN_BAND - Change was made from Catalyst Center, NOT_APPLICABLE - System Triggered)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"files": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"download_path": &schema.Schema{
													Description: `File download path (deprecated).
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"file_id": &schema.Schema{
													Description: `Unique file ID.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"file_type": &schema.Schema{
													Description: `Type of configuration file (RUNNINGCONFIG, STARTUPCONFIG, or VLAN).
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"id": &schema.Schema{
										Description: `Unique version ID.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_updated_time": &schema.Schema{
										Description: `Latest time stamp when the collected configuration is verified to be running on the device. LastUpdatedTime and createdTime will differ when verified configs are the same.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"startup_running_status": &schema.Schema{
										Description: `Sync status of running and startup configurations (IN_SYNC - if both startup and running config are same excluding dynamic configurations, OUT_OF_SYNC - otherwise). 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"syslog_config_event_dto": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config_method": &schema.Schema{
													Description: `Connection mode used to do the changes pn the device (if available in Syslog).
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"device_uuid": &schema.Schema{
													Description: `UUID of the device as recieved in syslog. 
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"login_ip_address": &schema.Schema{
													Description: `IP address from which the configuration changes were made (if available in Syslog).
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"out_of_band": &schema.Schema{
													Description: `True if configuration changes were made from outside of the Catalist Center. False otherwise.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"process_name": &schema.Schema{
													Description: `Name of the process that made configuration change (only available when configuration got changed by a program such as YANG suite )
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"syslog_time": &schema.Schema{
													Description: `Time of configuration change as recorded in the syslog.
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"terminal_name": &schema.Schema{
													Description: `Name of the terminal used to make changes on the device (if available in Syslog).
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"user_name": &schema.Schema{
													Description: `Name of the user who made the configuration change (if available in Syslog).
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"tags": &schema.Schema{
										Description: `Labels added to a configuration version.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func dataSourceNetworkDeviceConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vFileType, okFileType := d.GetOk("file_type")
	vCreatedTime, okCreatedTime := d.GetOk("created_time")
	vCreatedBy, okCreatedBy := d.GetOk("created_by")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConfigurationArchiveDetails")
		queryParams1 := catalystcentersdkgo.GetConfigurationArchiveDetailsQueryParams{}

		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okFileType {
			queryParams1.FileType = vFileType.(string)
		}
		if okCreatedTime {
			queryParams1.CreatedTime = vCreatedTime.(string)
		}
		if okCreatedBy {
			queryParams1.CreatedBy = vCreatedBy.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.ConfigurationArchive.GetConfigurationArchiveDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConfigurationArchiveDetails", err,
				"Failure at GetConfigurationArchiveDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationArchiveGetConfigurationArchiveDetailsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConfigurationArchiveDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationArchiveGetConfigurationArchiveDetailsItems(items *catalystcentersdkgo.ResponseConfigurationArchiveGetConfigurationArchiveDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_address"] = item.IPAddress
		respItem["device_id"] = item.DeviceID
		respItem["versions"] = flattenConfigurationArchiveGetConfigurationArchiveDetailsItemsVersions(item.Versions)
		respItem["device_name"] = item.DeviceName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationArchiveGetConfigurationArchiveDetailsItemsVersions(items *[]catalystcentersdkgo.ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["files"] = flattenConfigurationArchiveGetConfigurationArchiveDetailsItemsVersionsFiles(item.Files)
		respItem["created_by"] = item.CreatedBy
		respItem["config_change_type"] = item.ConfigChangeType
		respItem["syslog_config_event_dto"] = flattenConfigurationArchiveGetConfigurationArchiveDetailsItemsVersionsSyslogConfigEventDto(item.SyslogConfigEventDto)
		respItem["created_time"] = item.CreatedTime
		respItem["startup_running_status"] = item.StartupRunningStatus
		respItem["id"] = item.ID
		respItem["tags"] = item.Tags
		respItem["last_updated_time"] = item.LastUpdatedTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationArchiveGetConfigurationArchiveDetailsItemsVersionsFiles(items *[]catalystcentersdkgo.ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersionsFiles) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["file_type"] = item.FileType
		respItem["file_id"] = item.FileID
		respItem["download_path"] = item.DownloadPath
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationArchiveGetConfigurationArchiveDetailsItemsVersionsSyslogConfigEventDto(items *[]catalystcentersdkgo.ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersionsSyslogConfigEventDto) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["user_name"] = item.UserName
		respItem["device_uuid"] = item.DeviceUUID
		respItem["out_of_band"] = boolPtrToString(item.OutOfBand)
		respItem["config_method"] = item.ConfigMethod
		respItem["terminal_name"] = item.TerminalName
		respItem["login_ip_address"] = item.LoginIPAddress
		respItem["process_name"] = item.ProcessName
		respItem["syslog_time"] = item.SyslogTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
