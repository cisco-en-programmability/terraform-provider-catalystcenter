terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.32-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_reserve_ip_subpool" "example" {
  provider = catalystcenter
  parameters {
    ipv4_prefix      = "true"
    ipv4_global_pool = "172.30.200.0/24"

    //id                 = "string"
    ipv4_dhcp_servers  = ["100.0.2.2"]
    ipv4_dns_servers   = ["100.0.2.1"]
    ipv4_gate_way      = "172.30.200.1"
    ipv4_prefix_length = 27
    ipv4_subnet        = "172.30.200.0"
    # ipv4_total_host    = 2
    # ipv6_address_space = "false"
    # ipv6_dhcp_servers = ["abc.def.ghi.jkl2"]
    # ipv6_dns_servers   = ["abc.def.ghi.jkl"]
    # ipv6_gate_way      = "abc.def.ghi.jkl"
    # ipv6_global_pool   = "10.0.0.1"
    # ipv6_prefix        = "false"
    # ipv6_prefix_length = 10
    # ipv6_subnet        = "2"
    # ipv6_total_host    = 2
    name    = "devhf_SVG_medtek_vlan855"
    site_id = "e1c84cd3-81ce-4094-840c-6d485feec6d6"
    /*slaac_support      = "false"*/
    type = "Generic"
  }
}

output "catalystcenter_reserve_ip_subpool_example" {
  value = catalystcenter_reserve_ip_subpool.example
}
