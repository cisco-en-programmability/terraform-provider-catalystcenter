package catalystcenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceCustomPrompt() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on System Settings.

- Save custom prompt added by user in Catalyst Center. API will always override the existing prompts. User should
provide all the custom prompt in case of any update
`,

		CreateContext: resourceNetworkDeviceCustomPromptCreate,
		ReadContext:   resourceNetworkDeviceCustomPromptRead,
		UpdateContext: resourceNetworkDeviceCustomPromptUpdate,
		DeleteContext: resourceNetworkDeviceCustomPromptDelete,
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

						"custom_password_prompt": &schema.Schema{
							Description: `Password for Custom Prompt
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_username_prompt": &schema.Schema{
							Description: `Username for Custom Prompt
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_password_prompt": &schema.Schema{
							Description: `Default Password for Custom Prompt
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_username_prompt": &schema.Schema{
							Description: `Default Username for Custom Prompt
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

						"password_prompt": &schema.Schema{
							Description: `Password for Custom Prompt
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"username_prompt": &schema.Schema{
							Description: `Username for Custom Prompt
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

func resourceNetworkDeviceCustomPromptCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkDeviceCustomPromptCustomPromptPostAPI(ctx, "parameters.0", d)
	vPasswordPrompt := resourceItem["password_prompt"]

	vvPasswordPrompt := interfaceToString(vPasswordPrompt)

	vUsernamePrompt := resourceItem["username_prompt"]

	vvUsernamePrompt := interfaceToString(vUsernamePrompt)

	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.SystemSettings.CustomPromptPostAPI(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CustomPromptPOSTAPI", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CustomPromptPOSTAPI", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CustomPromptPOSTAPI", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
				"Failure when executing CustomPromptPOSTAPI", err1))
			return diags
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["password_prompt"] = vvPasswordPrompt
	resourceMap["username_prompt"] = vvUsernamePrompt
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkDeviceCustomPromptRead(ctx, d, m)
}

func resourceNetworkDeviceCustomPromptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	// resourceID := d.Id()
	// resourceMap := separateResourceID(resourceID)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CustomPromptSupportGETAPI")

		response1, restyResp1, err := client.SystemSettings.CustomPromptSupportGetAPI()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsCustomPromptSupportGetAPIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CustomPromptSupportGETAPI response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceNetworkDeviceCustomPromptUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing NetworkDeviceCustomPromptUpdate", err, "Update method is not supported",
		"Failure at NetworkDeviceCustomPromptUpdate, unexpected response", ""))

	return diags
}

func resourceNetworkDeviceCustomPromptDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing NetworkDeviceCustomPromptDelete", err, "Delete method is not supported",
		"Failure at NetworkDeviceCustomPromptDelete, unexpected response", ""))

	return diags
}

func expandRequestNetworkDeviceCustomPromptCustomPromptPostAPI(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSystemSettingsCustomPromptPostAPI {
	request := catalystcentersdkgo.RequestSystemSettingsCustomPromptPostAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username_prompt")))) {
		request.UsernamePrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_prompt")))) {
		request.PasswordPrompt = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
