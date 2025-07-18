package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReportsExecutionsDownload() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Returns report content. Save the response to a file by converting the response data as a blob and setting the file
format available from content-disposition response header.
`,

		ReadContext: dataSourceReportsExecutionsDownloadRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"execution_id": &schema.Schema{
				Description: `executionId path parameter. executionId of report execution
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"report_id": &schema.Schema{
				Description: `reportId path parameter. reportId of report
`,
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceReportsExecutionsDownloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID := d.Get("report_id")
	vExecutionID := d.Get("execution_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DownloadReportContent")
		vvReportID := vReportID.(string)
		vvExecutionID := vExecutionID.(string)

		// has_unknown_response: None

		response1, _, err := client.Reports.DownloadReportContent(vvReportID, vvExecutionID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing DownloadReportContent", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing DownloadReportContent", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		vvDirpath := d.Get("dirpath").(string)
		err = response1.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}
