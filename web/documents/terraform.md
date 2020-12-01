# Terraform

Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. Terraform can manage existing and popular service providers as well as custom in-house solutions.

## Handling Local Name Conflicts

Whenever possible, we recommend using a provider's preferred local name, which is usually the same as the "type" portion of its source address.

However, it's sometimes necessary to use two providers with the same preferred local name in the same module, usually when the providers are named after a generic infrastructure type. Terraform requires unique local names for each provider in a module, so you'll need to use a non-preferred name for at least one of them.

When this happens, we recommend combining each provider's namespace with its type name to produce compound local names with a dash:

```hcl
terraform {
  required_providers {
    # In the rare situation of using two providers that
    # have the same type name -- "http" in this example --
    # use a compound local name to distinguish them.
    hashicorp-http = {
      source  = "hashicorp/http"
      version = "~> 2.0"
    }
    mycorp-http = {
      source  = "mycorp/http"
      version = "~> 1.0"
    }
  }
}

# References to these providers elsewhere in the
# module will use these compound local names.
provider "mycorp_http" {
  # ...
}

data "http" "example" {
  provider = hashicorp_http
  #...
}
```