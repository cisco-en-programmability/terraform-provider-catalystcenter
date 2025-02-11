
resource "catalystcenter_site_kpi_summaries_top_n_analytics" "example" {
    provider = catalystcenter

    parameters {

      end_time = 1
      filters {

        key = "string"
        operator = "string"
        value = "string"
      }
      group_by = ["string"]
      start_time = 1
      top_n = 1
    }
}

output "catalystcenter_site_kpi_summaries_top_n_analytics_example" {
    value = catalystcenter_site_kpi_summaries_top_n_analytics.example
}
