variable "mysql_config" {
  description = "The Config for the MySQL DB"

  type = object({
    address = string
    port    = number
  })
  default = {
    address = "mock-mysql-address"
    port    = 12345
  }
}

variable "db_remote_state_bucket" {
  description = "The name of the S3 bucket for the database's remote state"
  type        = string
}

variable "db_remote_state_key" {
  description = "The name of the environment we're deploying to"
  type        = string
}

variable "environment" {
  description = "The name of the environment we're deploying to"
  type        = string
  default     = "develop"
}

variable "server_text" {
  description = "The text the web server should return"
  default     = "Hello, World"
  type        = string
}