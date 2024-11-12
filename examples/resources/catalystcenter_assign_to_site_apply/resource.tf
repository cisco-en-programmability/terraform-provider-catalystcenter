
resource "catalystcenter_assign_to_site_apply" "example" {
  provider = meraki
  parameters {

    device_ids = ["string"]
    site_id    = "string"
  }
}

output "catalystcenter_assign_to_site_apply_example" {
  value = catalystcenter_assign_to_site_apply.example
}