
resource "catalystcenter_network_bugs_trigger_scan" "example" {
  provider            = catalystcenter
  failed_devices_only = "false"
}

output "catalystcenter_network_bugs_trigger_scan_example" {
  value = catalystcenter_network_bugs_trigger_scan.example
}