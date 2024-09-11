
resource "catalystcenter_site_health_summaries_summary_analytics_create" "example" {
  provider          = catalystcenter
  id                = "string"
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  site_type         = "string"
  parameters {

    attributes = ["string"]
    end_time   = 1
    start_time = 1
    views      = ["string"]
  }
}

output "catalystcenter_site_health_summaries_summary_analytics_create_example" {
  value = catalystcenter_site_health_summaries_summary_analytics_create.example
}