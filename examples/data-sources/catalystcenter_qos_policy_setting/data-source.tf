
data "catalystcenter_qos_policy_setting" "example" {
  provider = catalystcenter
}

output "catalystcenter_qos_policy_setting_example" {
  value = data.catalystcenter_qos_policy_setting.example.item
}
