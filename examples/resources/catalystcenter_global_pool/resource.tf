
resource "catalystcenter_global_pool" "example" {
    provider = catalystcenter

    parameters {

      id = "string"
      settings {

        ippool {

          ip_address_space = "string"
          dhcp_server_ips = ["string"]
          dns_server_ips = ["string"]
          gateway = "string"
          id = "string"
          ip_pool_cidr = "string"
          ip_pool_name = "string"
          type = "string"
        }
      }
    }
}

output "catalystcenter_global_pool_example" {
    value = catalystcenter_global_pool.example
}
