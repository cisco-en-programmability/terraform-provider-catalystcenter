
data "catalystcenter_eox_status_summary" "example" {
    provider = catalystcenter
}

output "catalystcenter_eox_status_summary_example" {
    value = data.catalystcenter_eox_status_summary.example.item
}
