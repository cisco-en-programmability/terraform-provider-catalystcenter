
resource "catalystcenter_qos_policy_setting" "example" {
    provider = catalystcenter

    parameters {

      deploy_by_default_on_wired_devices = "false"
    }
}

output "catalystcenter_qos_policy_setting_example" {
    value = catalystcenter_qos_policy_setting.example
}
