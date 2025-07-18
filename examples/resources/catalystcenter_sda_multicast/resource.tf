
resource "catalystcenter_sda_multicast" "example" {
  provider = catalystcenter
 
  parameters {

    multicast_method = "string"
    multicast_type   = "string"
    multicast_vn_info {

      external_rp_ip_address = "string"
      internal_rp_ip_address = ["string"]
      ip_pool_name           = "string"
      ssm_info {

        ssm_group_range   = "string"
        ssm_wildcard_mask = "string"
      }
      virtual_network_name = "string"
    }
    site_name_hierarchy = "string"
  }
}

output "catalystcenter_sda_multicast_example" {
  value = catalystcenter_sda_multicast.example
}