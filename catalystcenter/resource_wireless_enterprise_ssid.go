package catalystcenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessEnterpriseSSID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- Creates enterprise SSID

- Update enterprise SSID

- Deletes given enterprise SSID
`,

		CreateContext: resourceWirelessEnterpriseSSIDCreate,
		ReadContext:   resourceWirelessEnterpriseSSIDRead,
		UpdateContext: resourceWirelessEnterpriseSSIDUpdate,
		DeleteContext: resourceWirelessEnterpriseSSIDDelete,
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

						"group_uuid": &schema.Schema{
							Description: `Group Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"inherited_group_name": &schema.Schema{
							Description: `Inherited Group Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"inherited_group_uuid": &schema.Schema{
							Description: `Inherited Group Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_override": &schema.Schema{
										Description: `Aaa Override
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_server": &schema.Schema{
										Description: `Auth Server
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"basic_service_set_client_idle_timeout": &schema.Schema{
										Description: `Basic Service Set ClientIdle Timeout
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"client_exclusion_timeout": &schema.Schema{
										Description: `Client Exclusion Timeout
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"client_rate_limit": &schema.Schema{
										Description: `Client Rate Limit. (in bits per second)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"coverage_hole_detection_enable": &schema.Schema{
										Description: `Coverage Hole Detection Enable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_basic_service_set_max_idle": &schema.Schema{
										Description: `Enable Basic Service Set Max Idle
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_broadcast_ssi_d": &schema.Schema{
										Description: `Enable Broadcast SSID
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_client_exclusion": &schema.Schema{
										Description: `Enable Client Exclusion
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_directed_multicast_service": &schema.Schema{
										Description: `Enable Directed MulticastService
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_fast_lane": &schema.Schema{
										Description: `Enable Fast Lane
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_mac_filtering": &schema.Schema{
										Description: `Enable MAC Filtering
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_neighbor_list": &schema.Schema{
										Description: `Enable NeighborList
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_session_time_out": &schema.Schema{
										Description: `Enable Session Time Out
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"fast_transition": &schema.Schema{
										Description: `Fast Transition
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_enabled": &schema.Schema{
										Description: `Is Enabled
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_fabric": &schema.Schema{
										Description: `Is Fabric
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mfp_client_protection": &schema.Schema{
										Description: `Mfp Client Protection
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"multi_psk_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"passphrase": &schema.Schema{
													Description: `Passphrase`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"passphrase_type": &schema.Schema{
													Description: `Passphrase Type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"priority": &schema.Schema{
													Description: `Priority
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"nas_options": &schema.Schema{
										Description: `Nas Options`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"passphrase": &schema.Schema{
										Description: `Passphrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protected_management_frame": &schema.Schema{
										Description: `Protected Management Frame`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"radio_policy": &schema.Schema{
										Description: `Radio Policy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_level": &schema.Schema{
										Description: `Security Level
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"session_time_out": &schema.Schema{
										Description: `sessionTimeOut
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"traffic_type": &schema.Schema{
										Description: `Traffic Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_type": &schema.Schema{
										Description: `Wlan Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `Version
`,
							Type:     schema.TypeInt,
							Computed: true,
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

						"aaa_override": &schema.Schema{
							Description: `Aaa Override
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"auth_key_mgmt": &schema.Schema{
							Description: `Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"basic_service_set_client_idle_timeout": &schema.Schema{
							Description: `Basic Service Set Client Idle Timeout
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"client_exclusion_timeout": &schema.Schema{
							Description: `Client Exclusion Timeout
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"client_rate_limit": &schema.Schema{
							Description: `Client Rate Limit (in bits per second)
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"coverage_hole_detection_enable": &schema.Schema{
							Description: `Coverage Hole Detection Enable
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_basic_service_set_max_idle": &schema.Schema{
							Description: `Enable Basic Service Set Max Idle
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_broadcast_ssi_d": &schema.Schema{
							Description: `Enable Broadcase SSID
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_client_exclusion": &schema.Schema{
							Description: `Enable Client Exclusion
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_directed_multicast_service": &schema.Schema{
							Description: `Enable Directed Multicast Service
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_fast_lane": &schema.Schema{
							Description: `Enable FastLane
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_mac_filtering": &schema.Schema{
							Description: `Enable MAC Filtering
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_neighbor_list": &schema.Schema{
							Description: `Enable Neighbor List
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_session_time_out": &schema.Schema{
							Description: `Enable Session Timeout
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"fast_transition": &schema.Schema{
							Description: `Fast Transition
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ghz24_policy": &schema.Schema{
							Description: `Ghz24 Policy
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ghz6_policy_client_steering": &schema.Schema{
							Description: `Ghz6 Policy Client Steering
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"mfp_client_protection": &schema.Schema{
							Description: `Management Frame Protection Client
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"multi_psk_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"passphrase": &schema.Schema{
										Description: `Passphrase
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"passphrase_type": &schema.Schema{
										Description: `Passphrase Type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `SSID NAME
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nas_options": &schema.Schema{
							Description: `Nas Options
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"passphrase": &schema.Schema{
							Description: `Passphrase
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policy_profile_name": &schema.Schema{
							Description: `Policy Profile Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Description: `Profile Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protected_management_frame": &schema.Schema{
							Description: `(Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_policy": &schema.Schema{
							Description: `Radio Policy Enum
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rsn_cipher_suite_ccmp256": &schema.Schema{
							Description: `Rsn Cipher Suite Ccmp256
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"rsn_cipher_suite_gcmp128": &schema.Schema{
							Description: `Rsn Cipher Suite Gcmp 128
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"rsn_cipher_suite_gcmp256": &schema.Schema{
							Description: `Rsn Cipher Suite Gcmp256
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"security_level": &schema.Schema{
							Description: `Security Level
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"session_time_out": &schema.Schema{
							Description: `Session Time Out
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ssid_name": &schema.Schema{
							Description: `ssidName path parameter. Enter the SSID name to be deleted
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"traffic_type": &schema.Schema{
							Description: `Traffic Type Enum (voicedata or data )
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessEnterpriseSSIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSID(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSSIDName := resourceItem["name"]
	vvSSIDName := interfaceToString(vSSIDName)

	queryParams1 := catalystcentersdkgo.GetEnterpriseSSIDQueryParams{}
	queryParams1.SSIDName = vvSSIDName
	getResponse2, err := searchWirelessGetEnterpriseSSID(m, queryParams1)
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvSSIDName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessEnterpriseSSIDRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Wireless.CreateEnterpriseSSID(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEnterpriseSSID", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEnterpriseSSID", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
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
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing CreateEnterpriseSSID", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvSSIDName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessEnterpriseSSIDRead(ctx, d, m)
}

func resourceWirelessEnterpriseSSIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSSIDName, okSSIDName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEnterpriseSSID")
		queryParams1 := catalystcentersdkgo.GetEnterpriseSSIDQueryParams{}

		if okSSIDName {
			queryParams1.SSIDName = vSSIDName
		}

		response1, restyResp1, _ := client.Wireless.GetEnterpriseSSID(&queryParams1)

		/*		if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetEnterpriseSSID", err,
					"Failure at GetEnterpriseSSID, unexpected response", ""))
				return diags
			}*/
		if response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetEnterpriseSSIDItems(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEnterpriseSSID search response",
				err))
			return diags
		}
		response1NoPointer := *response1
		responseItemSSID := *response1NoPointer[0].SSIDDetails
		responseItemSSIDItem := responseItemSSID[0]
		request := expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSID(ctx, "parameters.0", d)
		vParameters := map[string]interface{}{
			"basic_service_set_client_idle_timeout": request.BasicServiceSetClientIDleTimeout,
			"client_exclusion_timeout":              request.ClientExclusionTimeout,
			"enable_basic_service_set_max_idle":     boolPtrToString(request.EnableBasicServiceSetMaxIDle),
			"enable_broadcast_ssi_d":                boolPtrToString(responseItemSSIDItem.EnableBroadcastSSID),
			"enable_client_exclusion":               boolPtrToString(request.EnableClientExclusion),
			"enable_directed_multicast_service":     boolPtrToString(request.EnableDirectedMulticastService),
			"enable_fast_lane":                      boolPtrToString(responseItemSSIDItem.EnableFastLane),
			"enable_mac_filtering":                  boolPtrToString(responseItemSSIDItem.EnableMacFiltering),
			"enable_neighbor_list":                  boolPtrToString(request.EnableNeighborList),
			"enable_session_time_out":               boolPtrToString(request.EnableSessionTimeOut),
			"fast_transition":                       responseItemSSIDItem.FastTransition,
			"mfp_client_protection":                 responseItemSSIDItem.MfpClientProtection,
			"name":                                  responseItemSSIDItem.Name,
			"nas_options":                           responseItemSSIDItem.NasOptions,
			"passphrase":                            responseItemSSIDItem.Passphrase,
			"radio_policy":                          responseItemSSIDItem.RadioPolicy,
			"security_level":                        responseItemSSIDItem.SecurityLevel,
			"session_time_out":                      request.SessionTimeOut,
			"traffic_type":                          responseItemSSIDItem.TrafficType,
			"protected_management_frame":            responseItemSSIDItem.ProtectedManagementFrame,
			"ssid_name":                             request.Name,
		}
		// vItem1 := flattenWirelessGetEnterpriseSSIDItems(response1)
		if err := d.Set("parameters", []map[string]interface{}{vParameters}); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEnterpriseSSID search response",
				err))
			return diags
		}
	}
	return diags
}

func resourceWirelessEnterpriseSSIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSSIDName := resourceMap["name"]

	queryParams1 := catalystcentersdkgo.GetEnterpriseSSIDQueryParams{}
	queryParams1.SSIDName = vSSIDName
	item, err := searchWirelessGetEnterpriseSSID(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetEnterpriseSSID", err,
			"Failure at GetEnterpriseSSID, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		// NOTE: After testing, we consider that this trigger is only applicable for WPA2_PERSONAL security, so we limit it to the trigger being executed only with that type of security.
		if d.HasChange("parameters.0.passphrase") {
			request2 := expandRequestWirelessPskOverridePSKOverride(ctx, "parameters", d)
			if request2 != nil {
				log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request2))
			}
			if request2 != nil {
				newRequest2 := *request2
				newRequest2[0].SSID = vSSIDName
				request2 = &newRequest2
			}
			response2, restyResp2, err := client.Wireless.PSKOverride(request2)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] resty response for update operation => %v", restyResp2.String())
					diags = append(diags, diagErrorWithAltAndResponse(
						"Failure when executing PSKOverride", err, restyResp2.String(),
						"Failure at PSKOverride, unexpected response", ""))
					return diags
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing PSKOverride", err,
					"Failure at PSKOverride, unexpected response", ""))
				return diags
			}
			executionId := response2.ExecutionID
			log.Printf("[DEBUG] ExecutionID => %s", executionId)
			time.Sleep(5 * time.Second)
			response3, restyResp3, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp3 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp3.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
			for statusIsPending(response3.Status) {
				time.Sleep(20 * time.Second)
				response3, restyResp3, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
				if err != nil || response3 == nil {
					if restyResp3 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp3.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if statusIsFailure(response3.Status) {
				log.Printf("[DEBUG] Error %s", response3.BapiError)
				diags = append(diags, diagError(
					"Failure when executing PSKOverride", err))
				return diags
			}
		}
		log.Printf("[DEBUG] Name used for update operation %v", queryParams1)
		request1 := expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil {
			request1.Name = vSSIDName
		}
		response1, restyResp1, err := client.Wireless.UpdateEnterpriseSSID(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEnterpriseSSID", err, restyResp1.String(),
					"Failure at UpdateEnterpriseSSID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEnterpriseSSID", err,
				"Failure at UpdateEnterpriseSSID, unexpected response", ""))
			return diags
		}
		executionId := response1.ExecutionID
		log.Printf("[DEBUG] ExecutionID => %s", executionId)
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
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
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing UpdateEnterpriseSSID", err))
			return diags
		}
	}

	return resourceWirelessEnterpriseSSIDRead(ctx, d, m)
}

func resourceWirelessEnterpriseSSIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSSIDName := resourceMap["name"]

	queryParams1 := catalystcentersdkgo.GetEnterpriseSSIDQueryParams{}
	queryParams1.SSIDName = vSSIDName
	var vvSSIDName string
	item, err := searchWirelessGetEnterpriseSSID(m, queryParams1)
	if err != nil || item == nil {
		return diags
	}

	vvSSIDName = queryParams1.SSIDName
	response1, restyResp1, err := client.Wireless.DeleteEnterpriseSSID(vvSSIDName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteEnterpriseSSID", err, restyResp1.String(),
				"Failure at DeleteEnterpriseSSID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteEnterpriseSSID", err,
			"Failure at DeleteEnterpriseSSID, unexpected response", ""))
		return diags
	}

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
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
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteEnterpriseSSID", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestWirelessPskOverridePSKOverride(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessPSKOverride {
	request := catalystcentersdkgo.RequestWirelessPSKOverride{}
	if v := expandRequestWirelessPskOverridePSKOverrideItemArray(ctx, key, d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessPskOverridePSKOverrideItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemWirelessPSKOverride {
	request := []catalystcentersdkgo.RequestItemWirelessPSKOverride{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestWirelessPskOverridePSKOverrideItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessPskOverridePSKOverrideItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemWirelessPSKOverride {
	request := catalystcentersdkgo.RequestItemWirelessPSKOverride{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_name")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.PassPhrase = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSID(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessCreateEnterpriseSSID {
	request := catalystcentersdkgo.RequestWirelessCreateEnterpriseSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_level")))) {
		request.SecurityLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fast_lane")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fast_lane")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fast_lane")))) {
		request.EnableFastLane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_mac_filtering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_mac_filtering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_mac_filtering")))) {
		request.EnableMacFiltering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_policy")))) {
		request.RadioPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_broadcast_ssi_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) {
		request.EnableBroadcastSSID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition")))) {
		request.FastTransition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_session_time_out")))) {
		request.EnableSessionTimeOut = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_time_out")))) {
		request.SessionTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_client_exclusion")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_client_exclusion")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_client_exclusion")))) {
		request.EnableClientExclusion = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_exclusion_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) {
		request.ClientExclusionTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_basic_service_set_max_idle")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) {
		request.EnableBasicServiceSetMaxIDle = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".basic_service_set_client_idle_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) {
		request.BasicServiceSetClientIDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_directed_multicast_service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) {
		request.EnableDirectedMulticastService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_neighbor_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_neighbor_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_neighbor_list")))) {
		request.EnableNeighborList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mfp_client_protection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mfp_client_protection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mfp_client_protection")))) {
		request.MfpClientProtection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nas_options")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nas_options")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nas_options")))) {
		request.NasOptions = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_profile_name")))) {
		request.PolicyProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_override")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_override")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_override")))) {
		request.AAAOverride = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_enable")))) {
		request.CoverageHoleDetectionEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protected_management_frame")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protected_management_frame")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protected_management_frame")))) {
		request.ProtectedManagementFrame = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multi_psk_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multi_psk_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multi_psk_settings")))) {
		request.MultipSKSettings = expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSIDMultipSKSettingsArray(ctx, key+".multi_psk_settings", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_rate_limit")))) {
		request.ClientRateLimit = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_key_mgmt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_key_mgmt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_key_mgmt")))) {
		request.AuthKeyMgmt = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_gcmp256")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp256")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp256")))) {
		request.RsnCipherSuiteGcmp256 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_ccmp256")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_ccmp256")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_ccmp256")))) {
		request.RsnCipherSuiteCcmp256 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_gcmp128")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp128")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp128")))) {
		request.RsnCipherSuiteGcmp128 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz6_policy_client_steering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz6_policy_client_steering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz6_policy_client_steering")))) {
		request.Ghz6PolicyClientSteering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz24_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz24_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz24_policy")))) {
		request.Ghz24Policy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSIDMultipSKSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestWirelessCreateEnterpriseSSIDMultipSKSettings {
	request := []catalystcentersdkgo.RequestWirelessCreateEnterpriseSSIDMultipSKSettings{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSIDMultipSKSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSIDMultipSKSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessCreateEnterpriseSSIDMultipSKSettings {
	request := catalystcentersdkgo.RequestWirelessCreateEnterpriseSSIDMultipSKSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase_type")))) {
		request.PassphraseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSID(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessUpdateEnterpriseSSID {
	request := catalystcentersdkgo.RequestWirelessUpdateEnterpriseSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_level")))) {
		request.SecurityLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fast_lane")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fast_lane")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fast_lane")))) {
		request.EnableFastLane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_mac_filtering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_mac_filtering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_mac_filtering")))) {
		request.EnableMacFiltering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_policy")))) {
		request.RadioPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_broadcast_ssi_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) {
		request.EnableBroadcastSSID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition")))) {
		request.FastTransition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_session_time_out")))) {
		request.EnableSessionTimeOut = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_time_out")))) {
		request.SessionTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_client_exclusion")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_client_exclusion")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_client_exclusion")))) {
		request.EnableClientExclusion = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_exclusion_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) {
		request.ClientExclusionTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_basic_service_set_max_idle")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) {
		request.EnableBasicServiceSetMaxIDle = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".basic_service_set_client_idle_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) {
		request.BasicServiceSetClientIDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_directed_multicast_service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) {
		request.EnableDirectedMulticastService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_neighbor_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_neighbor_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_neighbor_list")))) {
		request.EnableNeighborList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mfp_client_protection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mfp_client_protection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mfp_client_protection")))) {
		request.MfpClientProtection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nas_options")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nas_options")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nas_options")))) {
		request.NasOptions = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_profile_name")))) {
		request.PolicyProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_override")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_override")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_override")))) {
		request.AAAOverride = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_enable")))) {
		request.CoverageHoleDetectionEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protected_management_frame")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protected_management_frame")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protected_management_frame")))) {
		request.ProtectedManagementFrame = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multi_psk_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multi_psk_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multi_psk_settings")))) {
		request.MultipSKSettings = expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSIDMultipSKSettingsArray(ctx, key+".multi_psk_settings", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_rate_limit")))) {
		request.ClientRateLimit = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_key_mgmt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_key_mgmt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_key_mgmt")))) {
		request.AuthKeyMgmt = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_gcmp256")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp256")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp256")))) {
		request.RsnCipherSuiteGcmp256 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_ccmp256")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_ccmp256")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_ccmp256")))) {
		request.RsnCipherSuiteCcmp256 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_gcmp128")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp128")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_gcmp128")))) {
		request.RsnCipherSuiteGcmp128 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz6_policy_client_steering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz6_policy_client_steering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz6_policy_client_steering")))) {
		request.Ghz6PolicyClientSteering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz24_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz24_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz24_policy")))) {
		request.Ghz24Policy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSIDMultipSKSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestWirelessUpdateEnterpriseSSIDMultipSKSettings {
	request := []catalystcentersdkgo.RequestWirelessUpdateEnterpriseSSIDMultipSKSettings{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSIDMultipSKSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSIDMultipSKSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessUpdateEnterpriseSSIDMultipSKSettings {
	request := catalystcentersdkgo.RequestWirelessUpdateEnterpriseSSIDMultipSKSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase_type")))) {
		request.PassphraseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetEnterpriseSSID(m interface{}, queryParams catalystcentersdkgo.GetEnterpriseSSIDQueryParams) (*catalystcentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetails, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetails
	var ite *catalystcentersdkgo.ResponseWirelessGetEnterpriseSSID
	ite, _, err = client.Wireless.GetEnterpriseSSID(&queryParams)
	if err != nil {
		return foundItem, err
	}

	if ite == nil {
		return foundItem, err
	}

	items := ite

	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		itemsCopy2 := *item.SSIDDetails
		for _, item := range itemsCopy2 {
			if item.Name == queryParams.SSIDName {

				return &item, err
			}
		}
	}
	return foundItem, err
}
