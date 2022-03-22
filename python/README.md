## Install

First, you need to setup the necessary environment variables on Terraform cloud.

Then install `pipenv` as package management.

```sh
pipenv install
```

```sh
cdktf deploy --auto-approve
```

```sh
cdktf destroy --auto-approve
```

## Docs

```sh
pipenv shell

import cdktf_cdktf_provider_aws

help(cdktf_cdktf_provider_aws.AwsProvider)

help(cdktf_cdktf_provider_aws.ecs)
```

Or

- https://constructs.dev/packages/@cdktf/provider-aws/v/5.0.48?submodule=ecs&lang=python
