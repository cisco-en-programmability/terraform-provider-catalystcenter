---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_devices_device_controllability_settings Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages read and update operations on Site Design.
  Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some
  device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs
  to manage devices. Changes are made on network devices during discovery, when adding a device to Inventory, or when
  assigning a device to a site. If changes are made to any settings that are under the scope of this process, these
  changes are applied to the network devices during the Provision and Update Telemetry Settings operations, even if Device
  Controllability is disabled. The following device settings will be enabled as part of Device Controllability when
  devices are discovered.
  SNMP Credentials.
    NETCONF Credentials.
  Subsequent to discovery, devices will be added to Inventory. The following device settings will be enabled when devices
  are added to inventory.
  Cisco TrustSec (CTS) Credentials.
  The following device settings will be enabled when devices are assigned to a site. Some of these settings can be defined
  at a site level under Design > Network Settings > Telemetry & Wireless.
  Wired Endpoint Data Collection Enablement.
    Controller Certificates.
    SNMP Trap Server Definitions.
    Syslog Server Definitions.
    Application Visibility.
    Application QoS Policy.
    Wireless Service Assurance (WSA).
    Wireless Telemetry.
    DTLS Ciphersuite.
    AP Impersonation.
  If Device Controllability is disabled, Catalyst Center does not configure any of the preceding credentials or settings
  on devices during discovery, at runtime, or during site assignment. However, the telemetry settings and related
  configuration are pushed when the device is provisioned or when the update Telemetry Settings action is performed.
  Catalyst Center identifies and automatically corrects the following telemetry configuration issues on the device.
  SWIM certificate issue.
    IOS WLC NA certificate issue.
    PKCS12 certificate issue.
    IOS telemetry configuration issue.
  The autocorrect telemetry config feature is supported only when Device Controllability is enabled.
---

# catalystcenter_network_devices_device_controllability_settings (Resource)

It manages read and update operations on Site Design.

- Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some
device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs
to manage devices. Changes are made on network devices during discovery, when adding a device to Inventory, or when
assigning a device to a site. If changes are made to any settings that are under the scope of this process, these
changes are applied to the network devices during the Provision and Update Telemetry Settings operations, even if Device
Controllability is disabled. The following device settings will be enabled as part of Device Controllability when
devices are discovered.


  SNMP Credentials.
  NETCONF Credentials.

Subsequent to discovery, devices will be added to Inventory. The following device settings will be enabled when devices
are added to inventory.


  Cisco TrustSec (CTS) Credentials.

The following device settings will be enabled when devices are assigned to a site. Some of these settings can be defined
at a site level under Design > Network Settings > Telemetry & Wireless.


  Wired Endpoint Data Collection Enablement.
  Controller Certificates.
  SNMP Trap Server Definitions.
  Syslog Server Definitions.
  Application Visibility.
  Application QoS Policy.
  Wireless Service Assurance (WSA).
  Wireless Telemetry.
  DTLS Ciphersuite.
  AP Impersonation.

If Device Controllability is disabled, Catalyst Center does not configure any of the preceding credentials or settings
on devices during discovery, at runtime, or during site assignment. However, the telemetry settings and related
configuration are pushed when the device is provisioned or when the update Telemetry Settings action is performed.
Catalyst Center identifies and automatically corrects the following telemetry configuration issues on the device.


  SWIM certificate issue.
  IOS WLC NA certificate issue.
  PKCS12 certificate issue.
  IOS telemetry configuration issue.

The autocorrect telemetry config feature is supported only when Device Controllability is enabled.

## Example Usage

```terraform
resource "catalystcenter_network_devices_device_controllability_settings" "example" {
  provider = catalystcenter
 
  parameters {

    autocorrect_telemetry_config = "false"
    device_controllability       = "false"
  }
}

output "catalystcenter_network_devices_device_controllability_settings_example" {
  value = catalystcenter_network_devices_device_controllability_settings.example
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

Optional:

- `autocorrect_telemetry_config` (String) If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
- `device_controllability` (String) If it is true, device controllability is enabled. If it is false, device controllability is disabled.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `autocorrect_telemetry_config` (String)
- `device_controllability` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_network_devices_device_controllability_settings.example "id:=string"
```
