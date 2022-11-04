# Broken Crossplane / Terraform Things

- Provider doesn't accept credential strings (only paths to strings)

- Most resource fields are auto-generated, don't fully populate resource GET response = no config drift detection, broken reconciliation

- Missing sensitive input type for things which may be sensitive (e.g. cloud provider credential url string:///some_secret_string)

- Provider IDs are randomly generated strings, making them difficult to reference...
