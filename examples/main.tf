terraform {
  required_providers {
    os2mo = {
      versions = ["0.3.0"]
      source = "hashicorp.com/edu/os2mo"
    }
  }
}

provider "os2mo" {
  username = "education"
  password = "test123"
}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}

data "os2mo_order" "order" {
  id = 1
}

output "order" {
  value = data.os2mo_order.order
}

resource "os2mo_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 2
  }
  items {
    coffee {
      id = 2
    }
    quantity = 3
  }
}

output "edu_order" {
  value = os2mo_order.edu
}


data "os2mo_order" "first" {
  id = 1
}

output "first_order" {
  value = data.os2mo_order.first
}
