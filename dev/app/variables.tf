variable "mysql_config" {
  description = "The Config for the MySQL DB"

  type = object({
    address = string
    port = number
  })
  default = {
    address = "mock-mysql-address"
    port = 12345
  }
}