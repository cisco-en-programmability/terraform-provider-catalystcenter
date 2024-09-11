
resource "catalystcenter_site_health_summaries_summary_analytics" "example" {
  provider = catalystcenter
  parameters {

    attributes = ["string"]
    end_time   = 1
    start_time = 1
    views      = ["string"]
  }
}

output "catalystcenter_site_health_summaries_summary_analytics_example" {
  value = catalystcenter_site_health_summaries_summary_analytics.example
}