package catalystcenter

import (
	"context"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceTagsInterfacesMembersAssociationsQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Tag.

- Fetches the tags associated with the given interface *ids*. Interfaces that don't have any tags associated will not be
included in the response. A tag is a user-defined or system-defined construct to group resources. When an interface is
tagged, it is called a member of the tag. *ids* can be fetched via */dna/intent/api/v1/interface* API.
`,

		CreateContext: resourceTagsInterfacesMembersAssociationsQueryCreate,
		ReadContext:   resourceTagsInterfacesMembersAssociationsQueryRead,
		DeleteContext: resourceTagsInterfacesMembersAssociationsQueryDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ids": &schema.Schema{
							Description: `List of member ids (network device or interface), maximum 500 ids can be passed.
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id of the member (network device or interface)
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Tag id
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Tag name
`,
													Type:     schema.TypeString,
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
				},
			},
		},
	}
}

func resourceTagsInterfacesMembersAssociationsQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestTagsInterfacesMembersAssociationsQueryQueryTheTagsAssociatedWithInterfacesV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Tag.QueryTheTagsAssociatedWithInterfacesV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing QueryTheTagsAssociatedWithInterfacesV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenTagQueryTheTagsAssociatedWithInterfacesV1Items(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting QueryTheTagsAssociatedWithInterfacesV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceTagsInterfacesMembersAssociationsQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceTagsInterfacesMembersAssociationsQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestTagsInterfacesMembersAssociationsQueryQueryTheTagsAssociatedWithInterfacesV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestTagQueryTheTagsAssociatedWithInterfacesV1 {
	request := catalystcentersdkgo.RequestTagQueryTheTagsAssociatedWithInterfacesV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ids")))) {
		request.IDs = interfaceToSliceString(v)
	}
	return &request
}

func flattenTagQueryTheTagsAssociatedWithInterfacesV1Items(items *[]catalystcentersdkgo.ResponseTagQueryTheTagsAssociatedWithInterfacesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["tags"] = flattenTagQueryTheTagsAssociatedWithInterfacesV1ItemsTags(item.Tags)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagQueryTheTagsAssociatedWithInterfacesV1ItemsTags(items *[]catalystcentersdkgo.ResponseTagQueryTheTagsAssociatedWithInterfacesV1ResponseTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
