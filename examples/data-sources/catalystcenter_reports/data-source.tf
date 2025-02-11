
data "catalystcenter_reports" "example" {
    provider = catalystcenter
    view_group_id = "string"
    view_id = "string"
}

output "catalystcenter_reports_example" {
    value = data.catalystcenter_reports.example.items
}

data "catalystcenter_reports" "example" {
    provider = catalystcenter
    report_id = "string"
}

output "catalystcenter_reports_example" {
    value = data.catalystcenter_reports.example.item
}
