package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexibleReportContent() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- This is used to download the flexible report. The API returns report content. Save the response to a file by
converting the response data as a blob and setting the file format available from content-disposition response header.
`,

		ReadContext: dataSourceFlexibleReportContentRead,
		Schema: map[string]*schema.Schema{
			"execution_id": &schema.Schema{
				Description: `executionId path parameter. Id of execution
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"report_id": &schema.Schema{
				Description: `reportId path parameter. Id of the report
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFlexibleReportContentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID := d.Get("report_id")
	vExecutionID := d.Get("execution_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DownloadFlexibleReportV1")
		vvReportID := vReportID.(string)
		vvExecutionID := vExecutionID.(string)

		// has_unknown_response: True

		response1, err := client.Reports.DownloadFlexibleReportV1(vvReportID, vvExecutionID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 DownloadFlexibleReportV1", err,
				"Failure at DownloadFlexibleReportV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DownloadFlexibleReportV1 response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
