
resource "catalystcenter_sda_fabric_site" "example" {
  provider = catalystcenter

  parameters {

    fabric_name         = "string"
    fabric_type         = "string"
    site_name_hierarchy = "string"
  }
}

output "catalystcenter_sda_fabric_site_example" {
  value = catalystcenter_sda_fabric_site.example
}