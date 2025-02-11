
resource "catalystcenter_site_kpi_summaries_summary_analytics" "example" {
    provider = catalystcenter
    item {

      attributes {



      }
    }
    parameters {

      attributes = ["string"]
      end_time = 1
      filters {

        key = "string"
        operator = "string"
        value = "string"
      }
      start_time = 1
    }
}

output "catalystcenter_site_kpi_summaries_summary_analytics_example" {
    value = catalystcenter_site_kpi_summaries_summary_analytics.example
}
