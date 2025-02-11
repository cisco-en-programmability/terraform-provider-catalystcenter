
resource "catalystcenter_wireless_profiles_id_site_tags_bulk" "example" {
    provider = meraki
    id = "string"
    parameters = [{
      
      items = []
    }]
}

output "catalystcenter_wireless_profiles_id_site_tags_bulk_example" {
    value = catalystcenter_wireless_profiles_id_site_tags_bulk.example
}