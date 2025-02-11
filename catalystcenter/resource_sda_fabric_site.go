package catalystcenter

import (
	"context"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricSite() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete Site from SDA Fabric

- Add Site in SDA Fabric
`,

		CreateContext: resourceSdaFabricSiteCreate,
		ReadContext:   resourceSdaFabricSiteRead,
		UpdateContext: resourceSdaFabricSiteUpdate,
		DeleteContext: resourceSdaFabricSiteDelete,
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

						"description": &schema.Schema{
							Description: `Fabric Site info successfully retrieved from sda fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_domain_type": &schema.Schema{
							Description: `Fabric Domain Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_name": &schema.Schema{
							Description: `Fabric Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_type": &schema.Schema{
							Description: `Fabric Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
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

						"fabric_name": &schema.Schema{
							Description: `Warning - Starting DNA Center 2.2.3.5 release, this field has been deprecated. SD-Access Fabric does not need it anymore.  It will be removed in future DNA Center releases.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fabric_type": &schema.Schema{
							Description: `Type of SD-Access Fabric. Allowed values are "FABRIC_SITE" or "FABRIC_ZONE".  Default value is "FABRIC_SITE".
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Existing site name hierarchy available at global level. For Example "Global/Chicago/Building21/Floor1"
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

func resourceSdaFabricSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricSiteAddSiteInSdaFabricV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)
	queryParamImport := catalystcentersdkgo.GetSiteFromSdaFabricV1QueryParams{}
	queryParamImport.SiteNameHierarchy = vvSiteNameHierarchy
	item2, _, err := client.Sda.GetSiteFromSdaFabricV1(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_name_hierarchy"] = item2.SiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricSiteRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddSiteInSdaFabricV1(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddSiteInSdaFabricV1", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddSiteInSdaFabricV1", err))
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
				"Failure when executing AddSiteInSdaFabricV1", err))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetSiteFromSdaFabricV1QueryParams{}
	queryParamValidate.SiteNameHierarchy = vvSiteNameHierarchy
	item3, _, err := client.Sda.GetSiteFromSdaFabricV1(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddSiteInSdaFabricV1", err,
			"Failure at AddSiteInSdaFabricV1, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["site_name_hierarchy"] = item3.SiteNameHierarchy

	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricSiteRead(ctx, d, m)
}

func resourceSdaFabricSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteFromSdaFabricV1")
		queryParams1 := catalystcentersdkgo.GetSiteFromSdaFabricV1QueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetSiteFromSdaFabricV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetSiteFromSdaFabricV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteFromSdaFabricV1 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSdaFabricSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricSiteRead(ctx, d, m)
}

func resourceSdaFabricSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := catalystcentersdkgo.DeleteSiteFromSdaFabricV1QueryParams{}

	vvSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	queryParamDelete.SiteNameHierarchy = vvSiteNameHierarchy

	response1, restyResp1, err := client.Sda.DeleteSiteFromSdaFabricV1(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSiteFromSdaFabricV1", err, restyResp1.String(),
				"Failure at DeleteSiteFromSdaFabricV1, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSiteFromSdaFabricV1", err,
			"Failure at DeleteSiteFromSdaFabricV1, unexpected response", ""))
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
				"Failure when executing DeleteSiteFromSdaFabricV1", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricSiteAddSiteInSdaFabricV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaAddSiteInSdaFabricV1 {
	request := catalystcentersdkgo.RequestSdaAddSiteInSdaFabricV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_name")))) {
		request.FabricName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_type")))) {
		request.FabricType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
