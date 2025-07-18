---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_event_email_config_update Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs update operation on Event Management.
  Update Email Destination
---

# catalystcenter_event_email_config_update (Resource)

It performs update operation on Event Management.

- Update Email Destination


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_event_email_config_update" "example" {
  provider = catalystcenter
  parameters = [{

    email_config_id = "string"
    from_email      = "string"
    primary_smt_p_config = [{

      host_name = "string"
      password  = "******"
      port      = "string"
      smtp_type = "string"
      user_name = "string"
    }]
    secondary_smt_p_config = [{

      host_name = "string"
      password  = "******"
      port      = "string"
      smtp_type = "string"
      user_name = "string"
    }]
    subject  = "string"
    to_email = "string"
  }]
}

output "catalystcenter_event_email_config_update_example" {
  value = catalystcenter_event_email_config_update.example
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

- `email_config_id` (String) Required only for update email configuration
- `from_email` (String) From Email
- `primary_smt_p_config` (Block List) (see [below for nested schema](#nestedblock--parameters--primary_smt_p_config))
- `secondary_smt_p_config` (Block List) (see [below for nested schema](#nestedblock--parameters--secondary_smt_p_config))
- `subject` (String) Subject
- `to_email` (String) To Email

<a id="nestedblock--parameters--primary_smt_p_config"></a>
### Nested Schema for `parameters.primary_smt_p_config`

Optional:

- `host_name` (String) Host Name
- `password` (String, Sensitive) Password
- `port` (String) Port
- `smtp_type` (String) smtpType
- `user_name` (String) User Name


<a id="nestedblock--parameters--secondary_smt_p_config"></a>
### Nested Schema for `parameters.secondary_smt_p_config`

Optional:

- `host_name` (String) Host Name
- `password` (String, Sensitive) Password
- `port` (String) Port
- `smtp_type` (String) smtpType
- `user_name` (String) User Name



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `status_uri` (String)
