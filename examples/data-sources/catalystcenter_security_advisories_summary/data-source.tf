
data "catalystcenter_security_advisories_summary" "example" {
    provider = catalystcenter
}

output "catalystcenter_security_advisories_summary_example" {
    value = data.catalystcenter_security_advisories_summary.example.item
}
