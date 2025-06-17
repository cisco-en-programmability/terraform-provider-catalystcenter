
resource "catalystcenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider = catalystcenter
 
  parameters {

    scalable_group_name = "string"
    site_name_hierarchy = "string"
    ssid_names          = ["string"]
    vlan_name           = "string"
  }
}

output "catalystcenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = catalystcenter_business_sda_hostonboarding_ssid_ippool.example
}