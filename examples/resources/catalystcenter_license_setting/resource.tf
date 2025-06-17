
resource "catalystcenter_license_setting" "example" {
  provider = catalystcenter
 
  parameters {

    auto_registration_virtual_account_id = "string"
    default_smart_account_id             = "string"
  }
}

output "catalystcenter_license_setting_example" {
  value = catalystcenter_license_setting.example
}