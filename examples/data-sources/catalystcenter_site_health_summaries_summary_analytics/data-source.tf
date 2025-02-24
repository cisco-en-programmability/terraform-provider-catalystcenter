
data "catalystcenter_site_health_summaries_summary_analytics" "example" {
    provider = catalystcenter
    attribute = "string"
    end_time = 1609459200
    id = "string"
    site_hierarchy = "string"
    site_hierarchy_id = "string"
    site_type = "string"
    start_time = 1609459200
    view = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_site_health_summaries_summary_analytics_example" {
    value = data.catalystcenter_site_health_summaries_summary_analytics.example.item
}
