---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_health_score_definitions Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages read and update operations on Devices.
  Update health threshold, include status of overall health status.
  And also to synchronize with global profile issue thresholds of the definition for given id. For detailed information
  about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
  programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-
  issueAndHealthDefinitions-1.0.0-resolved.yaml
---

# catalystcenter_health_score_definitions (Resource)

It manages read and update operations on Devices.

- Update health threshold, include status of overall health status.
And also to synchronize with global profile issue thresholds of the definition for given id. For detailed information
about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
issueAndHealthDefinitions-1.0.0-resolved.yaml

## Example Usage

```terraform
resource "catalystcenter_health_score_definitions" "example" {
  provider = catalystcenter

  parameters {

    id                             = "string"
    include_for_overall_health     = "false"
    synchronize_to_issue_threshold = "false"
    threshold_value                = 1.0
  }
}

output "catalystcenter_health_score_definitions_example" {
  value = catalystcenter_health_score_definitions.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `id` (String) id path parameter. Health score definition id.

Optional:

- `include_for_overall_health` (String) Include For Overall Health
- `synchronize_to_issue_threshold` (String) Synchronize To Issue Threshold
- `threshold_value` (Number) Thresehold Value


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `definition_status` (String)
- `description` (String)
- `device_family` (String)
- `display_name` (String)
- `id` (String)
- `include_for_overall_health` (String)
- `last_modified` (String)
- `name` (String)
- `synchronize_to_issue_threshold` (String)
- `threshold_value` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_health_score_definitions.example "id:=string"
```