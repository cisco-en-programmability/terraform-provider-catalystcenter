
resource "catalystcenter_app_policy_queuing_profile" "example" {
  provider = catalystcenter
 
  parameters {

    clause {

      instance_id = 1
      interface_speed_bandwidth_clauses {

        instance_id     = 1
        interface_speed = "string"
        tc_bandwidth_settings {

          bandwidth_percentage = 1
          instance_id          = 1
          traffic_class        = "string"
        }
      }
      is_common_between_all_interface_speeds = "false"
      tc_dscp_settings {

        dscp          = "string"
        instance_id   = 1
        traffic_class = "string"
      }
      type = "string"
    }
    description = "string"
    id          = "string"
    name        = "string"
  }
}

output "catalystcenter_app_policy_queuing_profile_example" {
  value = catalystcenter_app_policy_queuing_profile.example
}