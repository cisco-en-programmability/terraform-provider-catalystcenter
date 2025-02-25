
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_image_distribution" "example" {
  provider = catalystcenter

  parameters {
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a5"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a4"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a6"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a7"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a8"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
  }
}

output "catalystcenter_image_distribution_example" {
  value = catalystcenter_image_distribution.example
}

data "catalystcenter_task" "example" {
  depends_on = [catalystcenter_image_distribution.example]
  provider   = catalystcenter
  task_id    = catalystcenter_image_distribution.example.item.0.task_id
}

output "catalystcenter_task_example" {
  value = data.catalystcenter_task.example.item
}
