
data "catalystcenter_security_advisories_results_advisories" "example" {
    provider = catalystcenter
    cvss_base_score = "string"
    device_count = 1.0
    id = "string"
    limit = 1
    offset = 1
    order = "string"
    security_impact_rating = "string"
    sort_by = "string"
}

output "catalystcenter_security_advisories_results_advisories_example" {
    value = data.catalystcenter_security_advisories_results_advisories.example.items
}
