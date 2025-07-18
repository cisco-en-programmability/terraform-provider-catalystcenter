
resource "catalystcenter_site_kpi_summaries_trend_analytics" "example" {
  provider = catalystcenter
 
  parameters {

    attributes = ["string"]
    end_time   = 1
    filters {

      key      = "string"
      operator = "string"
      value    = "string"
    }
    page {

      limit           = 1
      offset          = 1
      timestamp_order = "string"
    }
    start_time     = 1
    trend_interval = "string"
  }
}

output "catalystcenter_site_kpi_summaries_trend_analytics_example" {
  value = catalystcenter_site_kpi_summaries_trend_analytics.example
}