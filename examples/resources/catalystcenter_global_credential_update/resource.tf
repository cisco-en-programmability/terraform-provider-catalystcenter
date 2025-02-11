
resource "catalystcenter_global_credential_update" "example" {
    provider = meraki
    global_credential_id = "string"
    parameters = [{
      
      site_uuids = ["string"]
    }]
}

output "catalystcenter_global_credential_update_example" {
    value = catalystcenter_global_credential_update.example
}