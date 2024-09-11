---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_app_policy_intent_create Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Application Policy.
  Create/Update/Delete application policy
---

# catalystcenter_app_policy_intent_create (Resource)

It performs create operation on Application Policy.

- Create/Update/Delete application policy



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_app_policy_intent_create" "example" {
  provider = catalystcenter
  parameters {

    create_list {

      advanced_policy_scope {

        advanced_policy_scope_element {

          group_id = ["string"]
          ssid     = ["string"]
        }
        name = "string"
      }
      consumer {

        scalable_group {

          id_ref = "string"
        }
      }
      contract {

        id_ref = "string"
      }
      delete_policy_status = "string"
      exclusive_contract {

        clause {

          device_removal_behavior = "string"
          host_tracking_enabled   = "false"
          relevance_level         = "string"
          type                    = "string"
        }
      }
      name         = "string"
      policy_scope = "string"
      priority     = "string"
      producer {

        scalable_group {

          id_ref = "string"
        }
      }
    }
    delete_list = ["string"]
    update_list {

      advanced_policy_scope {

        advanced_policy_scope_element {

          group_id = ["string"]
          id       = "string"
          ssid     = ["string"]
        }
        id   = "string"
        name = "string"
      }
      consumer {

        id = "string"
        scalable_group {

          id_ref = "string"
        }
      }
      contract {

        id_ref = "string"
      }
      delete_policy_status = "string"
      exclusive_contract {

        clause {

          device_removal_behavior = "string"
          host_tracking_enabled   = "false"
          id                      = "string"
          relevance_level         = "string"
          type                    = "string"
        }
        id = "string"
      }
      id           = "string"
      name         = "string"
      policy_scope = "string"
      priority     = "string"
      producer {

        id = "string"
        scalable_group {

          id_ref = "string"
        }
      }
    }
  }
}

output "catalystcenter_app_policy_intent_create_example" {
  value = catalystcenter_app_policy_intent_create.example
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

Optional:

- `create_list` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list))
- `delete_list` (List of String) Delete list of Group Based Policy ids
- `update_list` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list))

<a id="nestedblock--parameters--create_list"></a>
### Nested Schema for `parameters.create_list`

Optional:

- `advanced_policy_scope` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--advanced_policy_scope))
- `consumer` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--consumer))
- `contract` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--contract))
- `delete_policy_status` (String) NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
- `exclusive_contract` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--exclusive_contract))
- `name` (String) Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
- `policy_scope` (String) Policy name
- `priority` (String) Set to 4095 while producer refer to application Scalable group otherwise 100
- `producer` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--producer))

<a id="nestedblock--parameters--create_list--advanced_policy_scope"></a>
### Nested Schema for `parameters.create_list.advanced_policy_scope`

Optional:

- `advanced_policy_scope_element` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--advanced_policy_scope--advanced_policy_scope_element))
- `name` (String) Policy name

<a id="nestedblock--parameters--create_list--advanced_policy_scope--advanced_policy_scope_element"></a>
### Nested Schema for `parameters.create_list.advanced_policy_scope.advanced_policy_scope_element`

Optional:

- `group_id` (List of String) The site(s) ID where the Application QoS Policy will be deployed.
- `ssid` (List of String) Ssid



<a id="nestedblock--parameters--create_list--consumer"></a>
### Nested Schema for `parameters.create_list.consumer`

Optional:

- `scalable_group` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--consumer--scalable_group))

<a id="nestedblock--parameters--create_list--consumer--scalable_group"></a>
### Nested Schema for `parameters.create_list.consumer.scalable_group`

Optional:

- `id_ref` (String) Id ref to application Scalable group



<a id="nestedblock--parameters--create_list--contract"></a>
### Nested Schema for `parameters.create_list.contract`

Optional:

- `id_ref` (String) Id ref to Queueing profile


<a id="nestedblock--parameters--create_list--exclusive_contract"></a>
### Nested Schema for `parameters.create_list.exclusive_contract`

Optional:

- `clause` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--exclusive_contract--clause))

<a id="nestedblock--parameters--create_list--exclusive_contract--clause"></a>
### Nested Schema for `parameters.create_list.exclusive_contract.clause`

Optional:

- `device_removal_behavior` (String) Device eemoval behavior
- `host_tracking_enabled` (String) Is host tracking enabled
- `relevance_level` (String) Relevance level
- `type` (String) Type



<a id="nestedblock--parameters--create_list--producer"></a>
### Nested Schema for `parameters.create_list.producer`

Optional:

- `scalable_group` (Block List) (see [below for nested schema](#nestedblock--parameters--create_list--producer--scalable_group))

<a id="nestedblock--parameters--create_list--producer--scalable_group"></a>
### Nested Schema for `parameters.create_list.producer.scalable_group`

Optional:

- `id_ref` (String) Id ref to application-set or application Scalable group




<a id="nestedblock--parameters--update_list"></a>
### Nested Schema for `parameters.update_list`

Optional:

- `advanced_policy_scope` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--advanced_policy_scope))
- `consumer` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--consumer))
- `contract` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--contract))
- `delete_policy_status` (String) NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
- `exclusive_contract` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--exclusive_contract))
- `id` (String) Id of Group based policy
- `name` (String) Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
- `policy_scope` (String) Policy name
- `priority` (String) Set to 4095 while producer refer to application Scalable group otherwise 100
- `producer` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--producer))

<a id="nestedblock--parameters--update_list--advanced_policy_scope"></a>
### Nested Schema for `parameters.update_list.advanced_policy_scope`

Optional:

- `advanced_policy_scope_element` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--advanced_policy_scope--advanced_policy_scope_element))
- `id` (String) Id of Advance policy scope
- `name` (String) Policy name

<a id="nestedblock--parameters--update_list--advanced_policy_scope--advanced_policy_scope_element"></a>
### Nested Schema for `parameters.update_list.advanced_policy_scope.advanced_policy_scope_element`

Optional:

- `group_id` (List of String) The site(s) ID where the Application QoS Policy will be deployed.
- `id` (String) Id of Advance policy scope element
- `ssid` (List of String) Ssid



<a id="nestedblock--parameters--update_list--consumer"></a>
### Nested Schema for `parameters.update_list.consumer`

Optional:

- `id` (String) Id of Consumer
- `scalable_group` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--consumer--scalable_group))

<a id="nestedblock--parameters--update_list--consumer--scalable_group"></a>
### Nested Schema for `parameters.update_list.consumer.scalable_group`

Optional:

- `id_ref` (String) Id ref to application Scalable group



<a id="nestedblock--parameters--update_list--contract"></a>
### Nested Schema for `parameters.update_list.contract`

Optional:

- `id_ref` (String) Id ref to Queueing profile


<a id="nestedblock--parameters--update_list--exclusive_contract"></a>
### Nested Schema for `parameters.update_list.exclusive_contract`

Optional:

- `clause` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--exclusive_contract--clause))
- `id` (String) Id of Exclusive contract

<a id="nestedblock--parameters--update_list--exclusive_contract--clause"></a>
### Nested Schema for `parameters.update_list.exclusive_contract.clause`

Optional:

- `device_removal_behavior` (String) Device removal behavior
- `host_tracking_enabled` (String) Host tracking enabled
- `id` (String) Id of Business relevance or Application policy knobs clause
- `relevance_level` (String) Relevance level
- `type` (String) Type



<a id="nestedblock--parameters--update_list--producer"></a>
### Nested Schema for `parameters.update_list.producer`

Optional:

- `id` (String) Id of Producer
- `scalable_group` (Block List) (see [below for nested schema](#nestedblock--parameters--update_list--producer--scalable_group))

<a id="nestedblock--parameters--update_list--producer--scalable_group"></a>
### Nested Schema for `parameters.update_list.producer.scalable_group`

Optional:

- `id_ref` (String) Id ref to application-set or application Scalable group





<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)