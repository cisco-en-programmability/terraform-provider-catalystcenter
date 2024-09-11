# Configure provider with your Cisco Catalyst Center SDK credentials
provider "catalystcenter" {
  # Cisco Catalyst Center user name
  username = "admin"
  # it can be set using the environment variable CATALYST_BASE_URL

  # Cisco Catalyst Center password
  password = "admin123"
  # it can be set using the environment variable CATALYST_USERNAME

  # Cisco Catalyst Center base URL, FQDN or IP
  base_url = "https://172.168.196.2"
  # it can be set using the environment variable CATALYST_PASSWORD

  # Boolean to enable debugging
  debug = "false"
  # it can be set using the environment variable CATALYST_DEBUG

  # Boolean to enable or disable SSL certificate verification
  ssl_verify = "false"
  # it can be set using the environment variable CATALYST_SSL_VERIFY
}