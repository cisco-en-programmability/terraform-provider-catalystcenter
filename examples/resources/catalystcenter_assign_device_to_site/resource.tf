
resource "catalystcenter_assign_device_to_site" "example" {
  provider = catalystcenter
  site_id  = "string"
  parameters {

    device {

      ip = "string"
    }
  }
}

output "catalystcenter_assign_device_to_site_example" {
  value = catalystcenter_assign_device_to_site.example
}