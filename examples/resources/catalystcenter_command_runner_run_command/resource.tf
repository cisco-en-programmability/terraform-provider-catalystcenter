
resource "catalystcenter_command_runner_run_command" "example" {
  provider = catalystcenter
  parameters = [{

    commands     = ["string"]
    description  = "string"
    device_uuids = ["string"]
    name         = "string"
    timeout      = 1
  }]
}

output "catalystcenter_command_runner_run_command_example" {
  value = catalystcenter_command_runner_run_command.example
}