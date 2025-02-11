package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIssueDefinitionsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Issues.

`,

		CreateContext: resourceSystemIssueDefinitionsCountCreate,
		ReadContext:   resourceSystemIssueDefinitionsCountRead,
		UpdateContext: resourceSystemIssueDefinitionsCountUpdate,
		DeleteContext: resourceSystemIssueDefinitionsCountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
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
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemIssueDefinitionsCountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["xca_lle_rid"] = interfaceToString(resourceItem["xca_lle_rid"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSystemIssueDefinitionsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vXCaLLERID := resourceMap["xca_lle_rid"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1")

		headerParams1 := catalystcentersdkgo.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1HeaderParams{}
		queryParams1 := catalystcentersdkgo.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1QueryParams{}

		headerParams1.XCaLLERID = vXCaLLERID

		// has_unknown_response: None

		response1, restyResp1, err := client.Issues.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersV1 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSystemIssueDefinitionsCountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSystemIssueDefinitionsCountRead(ctx, d, m)
}

func resourceSystemIssueDefinitionsCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SystemIssueDefinitionsCount on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
