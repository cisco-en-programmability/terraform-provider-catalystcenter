terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_discovery" "example" {
  provider = catalystcenter

  parameters {

    #attribute_info = ["string"]
    cdp_level = 3
    # device_ids                = " "
    discovery_condition  = "Complete"
    discovery_status     = "Inactive"
    discovery_type       = "CDP"
    enable_password_list = []
    global_credential_id_list = [
      "52b611b4-5ee4-4581-8a01-a2062cd2d3c1",
      "ba85d1d0-616f-429c-9ed0-95369d13afa1",
      "627905c5-3120-490f-b373-c832cad07a84",
      "7bb143dd-89d6-4ae7-a5d2-afe50942fa03"
    ]
    /*
      http_read_credential {

        comments = "string"
        credential_type = "string"
        description = "string"
        id = "string"
        instance_tenant_id = "string"
        instance_uuid = "string"
        password = "******"
        port = 1
        secure = "false"
        username = "string"
      }
      http_write_credential {

        comments = "string"
        credential_type = "string"
        description = "string"
        id = "string"
        instance_tenant_id = "string"
        instance_uuid = "string"
        password = "******"
        port = 1
        secure = "false"
        username = "string"
      }
      */
    ip_address_list = "10.121.1.1"
    # ip_filter_list  = [""]
    is_auto_cdp = "true"
    lldp_level  = 1
    name        = "DMZ Lab 4"
    #netconf_port = "string"
    num_devices = 1
    #parent_discovery_id = "string"
    # password_list = []
    #preferred_mgmt_ipmethod = "string"
    protocol_order = "ssh"
    retry          = 1
    retry_count    = 3
    #snmp_auth_passphrase = "string"
    #snmp_auth_protocol = "string"
    #snmp_mode = "string"
    #snmp_priv_passphrase = "string"
    #snmp_priv_protocol = "string"
    #snmp_ro_community = "string"
    #snmp_ro_community_desc = "string"
    #snmp_rw_community = "string"
    #snmp_rw_community_desc = "string"
    #snmp_ro_community = "string"
    #snmp_ro_community_desc = "string"
    #snmp_rw_community = "string"
    #snmp_rw_community_desc = "string"
    #snmp_user_name = "string"
    #snmp_version = "string"
    time_out       = 30
    update_mgmt_ip = "false"
    #user_name_list = ["string"]

  }
}

output "catalystcenter_discovery_example" {
  value     = catalystcenter_discovery.example
  sensitive = true
}
