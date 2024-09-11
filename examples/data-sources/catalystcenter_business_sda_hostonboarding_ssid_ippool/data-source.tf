
data "catalystcenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider            = catalystcenter
  site_name_hierarchy = "string"
  vlan_name           = "string"
}

output "catalystcenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = data.catalystcenter_business_sda_hostonboarding_ssid_ippool.example.item
}
