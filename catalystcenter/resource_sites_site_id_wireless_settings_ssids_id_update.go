package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSitesSiteIDWirelessSettingsSSIDsIDUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action allows to either update SSID at global 'siteId' or override SSID at given non-global 'siteId'
`,

		CreateContext: resourceSitesSiteIDWirelessSettingsSSIDsIDUpdateCreate,
		ReadContext:   resourceSitesSiteIDWirelessSettingsSSIDsIDUpdateRead,
		DeleteContext: resourceSitesSiteIDWirelessSettingsSSIDsIDUpdateDelete,
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

						"task_id": &schema.Schema{
							Description: `Task ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task URL
`,
							Type:     schema.TypeString,
							Computed: true,
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
						"id": &schema.Schema{
							Description: `id path parameter. SSID ID
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site UUID
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"aaa_override": &schema.Schema{
							Description: `Activate the AAA Override feature when set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"acct_servers": &schema.Schema{
							Description: `List of Accounting server IpAddresses
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"acl_name": &schema.Schema{
							Description: `Pre-Auth Access Control List (ACL) Name
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"auth_server": &schema.Schema{
							Description: `For Guest SSIDs ('wlanType' is 'Guest' and 'l3AuthType' is 'web_auth'), the Authentication Server('authServer') is mandatory. Otherwise, it defaults to 'auth_external'.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"auth_servers": &schema.Schema{
							Description: `List of Authentication/Authorization server IpAddresses
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"auth_type": &schema.Schema{
							Description: `L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"basic_service_set_client_idle_timeout": &schema.Schema{
							Description: `This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out. Default is Basic ServiceSet ClientIdle Timeout if exists else 300. If it needs to be disabled , pass 0 as its value else valid range is [15 to 100000].
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"basic_service_set_max_idle_enable": &schema.Schema{
							Description: `Activate the maximum idle feature for the Basic Service Set
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"cckm_tsf_tolerance": &schema.Schema{
							Description: `The default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"client_exclusion_enable": &schema.Schema{
							Description: `Activate the feature that allows for the exclusion of clients
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"client_exclusion_timeout": &schema.Schema{
							Description: `This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"client_rate_limit": &schema.Schema{
							Description: `This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"coverage_hole_detection_enable": &schema.Schema{
							Description: `Activate Coverage Hole Detection feature when set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"directed_multicast_service_enable": &schema.Schema{
							Description: `The Directed Multicast Service feature becomes operational when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"egress_qos": &schema.Schema{
							Description: `Egress QOS
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"external_auth_ip_address": &schema.Schema{
							Description: `External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"fast_transition": &schema.Schema{
							Description: `Fast Transition
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"fast_transition_over_the_distributed_system_enable": &schema.Schema{
							Description: `Enable Fast Transition over the Distributed System when set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"ghz24_policy": &schema.Schema{
							Description: `2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"ghz6_policy_client_steering": &schema.Schema{
							Description: `True if 6 GHz Policy Client Steering is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"ingress_qos": &schema.Schema{
							Description: `Ingress QOS
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"is_ap_beacon_protection_enabled": &schema.Schema{
							Description: `When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key8021x": &schema.Schema{
							Description: `When set to true, the 802.1X authentication key is in use
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key8021x_plus_ft": &schema.Schema{
							Description: `When set to true, the 802.1X-Plus-FT authentication key is in use
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key8021x_sha256": &schema.Schema{
							Description: `When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_easy_psk": &schema.Schema{
							Description: `When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_owe": &schema.Schema{
							Description: `When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_psk": &schema.Schema{
							Description: `When set to true, the Pre-shared Key (PSK) authentication feature is enabled
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_psk_plus_ft": &schema.Schema{
							Description: `When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_psk_sha256": &schema.Schema{
							Description: `The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_sae": &schema.Schema{
							Description: `When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_sae_ext": &schema.Schema{
							Description: `When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_sae_ext_plus_ft": &schema.Schema{
							Description: `When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_sae_plus_ft": &schema.Schema{
							Description: `Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_suite_b1921x": &schema.Schema{
							Description: `When set to true, the SuiteB192-1x authentication key feature is enabled.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_auth_key_suite_b1x": &schema.Schema{
							Description: `When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_broadcast_ssi_d": &schema.Schema{
							Description: `When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_cckm_enabled": &schema.Schema{
							Description: `True if CCKM is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_enabled": &schema.Schema{
							Description: `Set SSID's admin status as 'Enabled' when set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_fast_lane_enabled": &schema.Schema{
							Description: `True if FastLane is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_hex": &schema.Schema{
							Description: `True if passphrase is in Hex format, else False.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_mac_filtering_enabled": &schema.Schema{
							Description: `When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_posturing_enabled": &schema.Schema{
							Description: `Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_radius_profiling_enabled": &schema.Schema{
							Description: `'true' if Radius profiling needs to be enabled, defaults to 'false' if not specified. At least one AAA/PSN server is required to enable Radius Profiling.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_random_mac_filter_enabled": &schema.Schema{
							Description: `Deny clients using randomized MAC addresses when set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"l3_auth_type": &schema.Schema{
							Description: `L3 Authentication Type. When 'wlanType' is 'Enterprise', ‘l3AuthType' is optional and defaults to 'open' if not specified. If 'wlanType' is 'Guest' then 'l3AuthType' is mandatory.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"management_frame_protection_clientprotection": &schema.Schema{
							Description: `Management Frame Protection Client
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"multi_psk_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"passphrase": &schema.Schema{
										Description: `Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"passphrase_type": &schema.Schema{
										Description: `Passphrase Type(default:ASCII)
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"nas_options": &schema.Schema{
							Description: `Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"neighbor_list_enable": &schema.Schema{
							Description: `The Neighbor List feature is enabled when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"open_ssid": &schema.Schema{
							Description: `Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"passphrase": &schema.Schema{
							Description: `Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"policy_profile_name": &schema.Schema{
							Description: `Policy Profile Name.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Description: `WLAN Profile Name, if not passed autogenerated profile name will be assigned.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"protected_management_frame": &schema.Schema{
							Description: `(REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"rsn_cipher_suite_ccmp128": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"rsn_cipher_suite_ccmp256": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"rsn_cipher_suite_gcmp128": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"rsn_cipher_suite_gcmp256": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"session_time_out": &schema.Schema{
							Description: `This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"session_time_out_enable": &schema.Schema{
							Description: `Turn on the feature that imposes a time limit on user sessions
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"sleeping_client_enable": &schema.Schema{
							Description: `When set to true, this will activate the timeout settings that apply to clients in sleep mode
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"sleeping_client_timeout": &schema.Schema{
							Description: `This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"ssid": &schema.Schema{
							Description: `Name of the SSID
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"ssid_radio_type": &schema.Schema{
							Description: `Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"web_passthrough": &schema.Schema{
							Description: `When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"wlan_band_select_enable": &schema.Schema{
							Description: `Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"wlan_type": &schema.Schema{
							Description: `Wlan Type
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
	}
}

func resourceSitesSiteIDWirelessSettingsSSIDsIDUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vSiteID := resourceItem["site_id"]
	vID := resourceItem["id"]

	vvSiteID := vSiteID.(string)
	vvID := vID.(string)
	request1 := expandRequestSitesSiteIDWirelessSettingsSSIDsIDUpdateUpdateOrOverridessid(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Wireless.UpdateOrOverridessid(vvSiteID, vvID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UpdateOrOverrideSSID", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil || strings.Contains(restyResp3.String(), "<!doctype html>") {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing UpdateOrOverrideSSID", err1))
			return diags
		}
	}

	vItem1 := flattenWirelessUpdateOrOverridessidItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateOrOverridessid response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceSitesSiteIDWirelessSettingsSSIDsIDUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSitesSiteIDWirelessSettingsSSIDsIDUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSitesSiteIDWirelessSettingsSSIDsIDUpdateUpdateOrOverridessid(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessUpdateOrOverridessid {
	request := catalystcentersdkgo.RequestWirelessUpdateOrOverridessid{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_fast_lane_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_fast_lane_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_fast_lane_enabled")))) {
		request.IsFastLaneEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_mac_filtering_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_mac_filtering_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_mac_filtering_enabled")))) {
		request.IsMacFilteringEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_radio_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_radio_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_radio_type")))) {
		request.SSIDRadioType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_broadcast_ssi_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_broadcast_ssi_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_broadcast_ssi_d")))) {
		request.IsBroadcastSSID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition")))) {
		request.FastTransition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_time_out_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_time_out_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_time_out_enable")))) {
		request.SessionTimeOutEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_time_out")))) {
		request.SessionTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_exclusion_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_exclusion_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_exclusion_enable")))) {
		request.ClientExclusionEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_exclusion_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) {
		request.ClientExclusionTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".basic_service_set_max_idle_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".basic_service_set_max_idle_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".basic_service_set_max_idle_enable")))) {
		request.BasicServiceSetMaxIDleEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".basic_service_set_client_idle_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) {
		request.BasicServiceSetClientIDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".directed_multicast_service_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".directed_multicast_service_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".directed_multicast_service_enable")))) {
		request.DirectedMulticastServiceEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbor_list_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbor_list_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbor_list_enable")))) {
		request.NeighborListEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_frame_protection_clientprotection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_frame_protection_clientprotection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_frame_protection_clientprotection")))) {
		request.ManagementFrameProtectionClientprotection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nas_options")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nas_options")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nas_options")))) {
		request.NasOptions = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
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
		request.MultipSKSettings = expandRequestSitesSiteIDWirelessSettingsSSIDsIDUpdateUpdateOrOverridessidMultipSKSettingsArray(ctx, key+".multi_psk_settings", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_rate_limit")))) {
		request.ClientRateLimit = interfaceToIntPtr(v)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rsn_cipher_suite_ccmp128")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rsn_cipher_suite_ccmp128")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rsn_cipher_suite_ccmp128")))) {
		request.RsnCipherSuiteCcmp128 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz6_policy_client_steering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz6_policy_client_steering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz6_policy_client_steering")))) {
		request.Ghz6PolicyClientSteering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key8021x")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key8021x")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key8021x")))) {
		request.IsAuthKey8021X = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key8021x_plus_ft")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key8021x_plus_ft")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key8021x_plus_ft")))) {
		request.IsAuthKey8021XPlusFT = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key8021x_sha256")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key8021x_sha256")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key8021x_sha256")))) {
		request.IsAuthKey8021XSHA256 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_sae")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_sae")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_sae")))) {
		request.IsAuthKeySae = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_sae_plus_ft")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_sae_plus_ft")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_sae_plus_ft")))) {
		request.IsAuthKeySaePlusFT = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_psk")))) {
		request.IsAuthKeyPSK = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_psk_plus_ft")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_psk_plus_ft")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_psk_plus_ft")))) {
		request.IsAuthKeyPSKPlusFT = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_owe")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_owe")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_owe")))) {
		request.IsAuthKeyOWE = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_easy_psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_easy_psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_easy_psk")))) {
		request.IsAuthKeyEasyPSK = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_psk_sha256")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_psk_sha256")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_psk_sha256")))) {
		request.IsAuthKeyPSKSHA256 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".open_ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".open_ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".open_ssid")))) {
		request.OpenSSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_band_select_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_band_select_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_band_select_enable")))) {
		request.WLANBandSelectEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_enabled")))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_servers")))) {
		request.AuthServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".acct_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".acct_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".acct_servers")))) {
		request.AcctServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".egress_qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".egress_qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".egress_qos")))) {
		request.EgressQos = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ingress_qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ingress_qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ingress_qos")))) {
		request.IngressQos = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_type")))) {
		request.WLANType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".l3_auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".l3_auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".l3_auth_type")))) {
		request.L3AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_server")))) {
		request.AuthServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_auth_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_auth_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_auth_ip_address")))) {
		request.ExternalAuthIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".web_passthrough")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".web_passthrough")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".web_passthrough")))) {
		request.WebPassthrough = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sleeping_client_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sleeping_client_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sleeping_client_enable")))) {
		request.SleepingClientEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sleeping_client_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sleeping_client_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sleeping_client_timeout")))) {
		request.SleepingClientTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".acl_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".acl_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".acl_name")))) {
		request.ACLName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_posturing_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_posturing_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_posturing_enabled")))) {
		request.IsPosturingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_suite_b1x")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_suite_b1x")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_suite_b1x")))) {
		request.IsAuthKeySuiteB1X = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_suite_b1921x")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_suite_b1921x")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_suite_b1921x")))) {
		request.IsAuthKeySuiteB1921X = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_sae_ext")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_sae_ext")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_sae_ext")))) {
		request.IsAuthKeySaeExt = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auth_key_sae_ext_plus_ft")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auth_key_sae_ext_plus_ft")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auth_key_sae_ext_plus_ft")))) {
		request.IsAuthKeySaeExtPlusFT = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_ap_beacon_protection_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_ap_beacon_protection_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_ap_beacon_protection_enabled")))) {
		request.IsApBeaconProtectionEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz24_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz24_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz24_policy")))) {
		request.Ghz24Policy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cckm_tsf_tolerance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cckm_tsf_tolerance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cckm_tsf_tolerance")))) {
		request.CckmTsfTolerance = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_cckm_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_cckm_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_cckm_enabled")))) {
		request.IsCckmEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_hex")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_hex")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_hex")))) {
		request.IsHex = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_random_mac_filter_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_random_mac_filter_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_random_mac_filter_enabled")))) {
		request.IsRandomMacFilterEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition_over_the_distributed_system_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition_over_the_distributed_system_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition_over_the_distributed_system_enable")))) {
		request.FastTransitionOverTheDistributedSystemEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_radius_profiling_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_radius_profiling_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_radius_profiling_enabled")))) {
		request.IsRadiusProfilingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_profile_name")))) {
		request.PolicyProfileName = interfaceToString(v)
	}
	return &request
}

func expandRequestSitesSiteIDWirelessSettingsSSIDsIDUpdateUpdateOrOverridessidMultipSKSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestWirelessUpdateOrOverridessidMultipSKSettings {
	request := []catalystcentersdkgo.RequestWirelessUpdateOrOverridessidMultipSKSettings{}
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
		i := expandRequestSitesSiteIDWirelessSettingsSSIDsIDUpdateUpdateOrOverridessidMultipSKSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSitesSiteIDWirelessSettingsSSIDsIDUpdateUpdateOrOverridessidMultipSKSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessUpdateOrOverridessidMultipSKSettings {
	request := catalystcentersdkgo.RequestWirelessUpdateOrOverridessidMultipSKSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase_type")))) {
		request.PassphraseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	return &request
}

func flattenWirelessUpdateOrOverridessidItem(item *catalystcentersdkgo.ResponseWirelessUpdateOrOverridessidResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
