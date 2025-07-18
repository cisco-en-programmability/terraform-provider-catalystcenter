---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_icap_settings_deploy_id_delete_deploy Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Sensors.
  Remove the ICAP configuration from the device by id without preview-deploy. The path parameter id can be retrieved
  from the GET /dna/intent/api/v1/icapSettings API. The response body contains a task object with a taskId and a URL.
  Use the URL to check the task status. ICAP FULL, ONBOARDING, OTA, and SPECTRUM configurations have a durationInMins
  field. A disable task is scheduled to remove the configuration from the device. Removing the ICAP intent should be done
  after the pre-scheduled disable task has been deployed. For detailed information about the usage of the API, please
  refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
  specs/blob/main/Assurance/CECatCenterOrg-ICAPAPIs-1.0.0-resolved.yaml
---

# catalystcenter_icap_settings_deploy_id_delete_deploy (Resource)

It performs create operation on Sensors.

- Remove the ICAP configuration from the device by *id* without preview-deploy. The path parameter *id* can be retrieved
from the **GET /dna/intent/api/v1/icapSettings** API. The response body contains a task object with a taskId and a URL.
Use the URL to check the task status. ICAP FULL, ONBOARDING, OTA, and SPECTRUM configurations have a durationInMins
field. A disable task is scheduled to remove the configuration from the device. Removing the ICAP intent should be done
after the pre-scheduled disable task has been deployed. For detailed information about the usage of the API, please
refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_icap_settings_deploy_id_delete_deploy" "example" {
  provider = catalystcenter
  id       = "string"
  parameters = [{

    object = "string"
  }]
}

output "catalystcenter_icap_settings_deploy_id_delete_deploy_example" {
  value = catalystcenter_icap_settings_deploy_id_delete_deploy.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `id` (String) id path parameter. A unique ID of the deployed ICAP object, which can be obtained from **GET /dna/intent/api/v1/icapSettings**

Optional:

- `object` (String) object


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
