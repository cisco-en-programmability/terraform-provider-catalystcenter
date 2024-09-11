
resource "catalystcenter_compliance" "example" {
  provider = catalystcenter
  parameters {

    categories   = ["string"]
    device_uuids = ["string"]
    trigger_full = "false"
  }
}

output "catalystcenter_compliance_example" {
  value = catalystcenter_compliance.example
}