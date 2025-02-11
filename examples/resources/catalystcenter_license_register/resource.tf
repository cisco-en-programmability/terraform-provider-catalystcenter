
resource "catalystcenter_license_register" "example" {
    provider = meraki
    parameters = [{
      
      smart_account_id = "string"
      virtual_account_id = "string"
    }]
}

output "catalystcenter_license_register_example" {
    value = catalystcenter_license_register.example
}