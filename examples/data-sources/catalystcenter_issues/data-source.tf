
data "catalystcenter_issues" "example" {
    provider = catalystcenter
    ai_driven = "string"
    device_id = "string"
    end_time = 1609459200
    issue_status = "string"
    mac_address = "string"
    priority = "string"
    site_id = "string"
    start_time = 1609459200
}

output "catalystcenter_issues_example" {
    value = data.catalystcenter_issues.example.items
}
