# Terraform Provider Tsuru

Initial code to integrate tsuru with terraform.io.

Terraform configuration example:
```
provider "tsuru" {
    target = "https://tsuru.deployment.com/1.0",
    token = "tsuru-token"
}

resource "tsuru_app" "terraform-provider-tsuru-test-app" {
    name = "terraform-provider-tsuru-test-app",
    platform = "static",
    team = "nice-team",
    plan = "small",
    router = "galeb_dev",
    pool = "dev",
    description = "Terraform tsuru provider testing app",
}
```
