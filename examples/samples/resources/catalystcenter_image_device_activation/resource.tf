
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_image_device_activation" "example" {
  provider = catalystcenter

  parameters {
    //client_type=                    "string"
    //client_url=                     "string"
    //schedule_validate=              "false"
    payload {
      activate_lower_image_version = "true"
      device_upgrade_mode          = "string"
      device_uuid                  = "3eb928b8-2414-4121-ac35-1247e5d666a4"
      distribute_if_needed         = "true"
      image_uuid_list = [
        "6af2b040-a312-4f57-8c8e-21f5e3e07598"
      ]
      smu_image_uuid_list = [
        "6af2b040-a312-4f57-8c8e-21f5e3e07597"
      ]
    }
  }
}

output "catalystcenter_image_device_activation_example" {
  value = catalystcenter_image_device_activation.example
}

data "catalystcenter_task" "example" {
  depends_on = [catalystcenter_image_device_activation.example]
  provider   = catalystcenter
  task_id    = catalystcenter_image_device_activation.example.item.0.task_id
}

output "catalystcenter_task_example" {
  value = data.catalystcenter_task.example.item
}
