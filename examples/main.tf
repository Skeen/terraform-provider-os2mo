provider "hashicups" {
  username = "education"
  password = "test1234"
}


module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}

data "hashicups_order" "order" {
  id = 2
}

output "order" {
  value = data.hashicups_order.order
}
