terraform {
  required_providers {
    os2mo = {
      versions = ["0.3.0"]
      source = "hashicorp.com/edu/os2mo"
    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

data "os2mo_coffees" "all" {}

# Returns all coffees
output "all_coffees" {
  value = data.os2mo_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in data.os2mo_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}
