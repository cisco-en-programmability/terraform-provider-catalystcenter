package catalystcenter

import (
	"context"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBusinessSdaHostonboardingSSIDIPpool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Fabric Wireless.

- Add SSID to IP Pool Mapping

- Update SSID to IP Pool Mapping
`,

		CreateContext: resourceBusinessSdaHostonboardingSSIDIPpoolCreate,
		ReadContext:   resourceBusinessSdaHostonboardingSSIDIPpoolRead,
		UpdateContext: resourceBusinessSdaHostonboardingSSIDIPpoolUpdate,
		DeleteContext: resourceBusinessSdaHostonboardingSSIDIPpoolDelete,
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

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"scalable_group_name": &schema.Schema{
										Description: `Scalable Group Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"vlan_name": &schema.Schema{
							Description: `VLAN Name`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssid_names": &schema.Schema{
							Description: `List of SSIDs
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vlan_name": &schema.Schema{
							Description: `VLAN Name
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

func resourceBusinessSdaHostonboardingSSIDIPpoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestBusinessSdaHostonboardingSSIDIPpoolAddSSIDToIPPoolMappingV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vVLANName := resourceItem["vlan_name"]
	vvVLANName := interfaceToString(vVLANName)
	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)
	queryParamImport := catalystcentersdkgo.GetSSIDToIPPoolMappingV1QueryParams{}
	queryParamImport.VLANName = vvVLANName
	queryParamImport.SiteNameHierarchy = vvSiteNameHierarchy
	item2, _, err := client.FabricWireless.GetSSIDToIPPoolMappingV1(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["vlan_name"] = item2.VLANName
		resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.FabricWireless.AddSSIDToIPPoolMappingV1(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddSSIDToIPPoolMappingV1", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddSSIDToIPPoolMappingV1", err))
		return diags
	}
	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddSSIDToIPPoolMappingV1", err))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetSSIDToIPPoolMappingV1QueryParams{}
	queryParamValidate.VLANName = vvVLANName
	queryParamValidate.SiteNameHierarchy = vvSiteNameHierarchy
	item3, _, err := client.FabricWireless.GetSSIDToIPPoolMappingV1(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddSSIDToIPPoolMappingV1", err,
			"Failure at AddSSIDToIPPoolMappingV1, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["vlan_name"] = item3.VLANName
	resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy

	d.SetId(joinResourceID(resourceMap))
	return resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx, d, m)
}

func resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vVLANName := resourceMap["vlan_name"]

	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSSIDToIPPoolMappingV1")
		queryParams1 := catalystcentersdkgo.GetSSIDToIPPoolMappingV1QueryParams{}

		queryParams1.VLANName = vVLANName

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		// has_unknown_response: None

		response1, restyResp1, err := client.FabricWireless.GetSSIDToIPPoolMappingV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessGetSSIDToIPPoolMappingV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDToIPPoolMappingV1 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceBusinessSdaHostonboardingSSIDIPpoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestBusinessSdaHostonboardingSSIDIPpoolUpdateSSIDToIPPoolMappingV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.FabricWireless.UpdateSSIDToIPPoolMappingV1(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSSIDToIPPoolMappingV1", err, restyResp1.String(),
					"Failure at UpdateSSIDToIPPoolMappingV1, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSSIDToIPPoolMappingV1", err,
				"Failure at UpdateSSIDToIPPoolMappingV1, unexpected response", ""))
			return diags
		}

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
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
			for response2.Status == "IN_PROGRESS" {
				time.Sleep(10 * time.Second)
				response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
				if err != nil || response2 == nil {
					if restyResp2 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if response2.Status == "FAILURE" {
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing UpdateSSIDToIPPoolMappingV1", err))
				return diags
			}
		}

	}

	return resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx, d, m)
}

func resourceBusinessSdaHostonboardingSSIDIPpoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete BusinessSdaHostonboardingSSIDIPpool on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestBusinessSdaHostonboardingSSIDIPpoolAddSSIDToIPPoolMappingV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestFabricWirelessAddSSIDToIPPoolMappingV1 {
	request := catalystcentersdkgo.RequestFabricWirelessAddSSIDToIPPoolMappingV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_name")))) {
		request.ScalableGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_names")))) {
		request.SSIDNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestBusinessSdaHostonboardingSSIDIPpoolUpdateSSIDToIPPoolMappingV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestFabricWirelessUpdateSSIDToIPPoolMappingV1 {
	request := catalystcentersdkgo.RequestFabricWirelessUpdateSSIDToIPPoolMappingV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_name")))) {
		request.ScalableGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_names")))) {
		request.SSIDNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
