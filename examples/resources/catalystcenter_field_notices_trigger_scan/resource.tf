
resource "catalystcenter_field_notices_trigger_scan" "example" {
  provider            = catalystcenter
  failed_devices_only = "false"
}

output "catalystcenter_field_notices_trigger_scan_example" {
  value = catalystcenter_field_notices_trigger_scan.example
}