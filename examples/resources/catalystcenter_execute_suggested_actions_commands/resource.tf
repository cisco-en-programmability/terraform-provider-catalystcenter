
resource "catalystcenter_execute_suggested_actions_commands" "example" {
  provider = catalystcenter
  parameters = [{

    entity_type  = "string"
    entity_value = "string"
  }]
}

output "catalystcenter_execute_suggested_actions_commands_example" {
  value = catalystcenter_execute_suggested_actions_commands.example
}