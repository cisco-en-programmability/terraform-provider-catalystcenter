
resource "catalystcenter_security_advisories_trigger_scan" "example" {
    provider = meraki
    failed_devices_only = "false"
}

output "catalystcenter_security_advisories_trigger_scan_example" {
    value = catalystcenter_security_advisories_trigger_scan.example
}