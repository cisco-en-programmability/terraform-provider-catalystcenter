
resource "catalystcenter_sites_dns_settings" "example" {
  provider = catalystcenter

  parameters {

    dns {

      dns_servers = ["string"]
      domain_name = "string"
    }
    id = "string"
  }
}

output "catalystcenter_sites_dns_settings_example" {
  value = catalystcenter_sites_dns_settings.example
}