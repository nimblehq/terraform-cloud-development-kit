## Install

First, you need to setup the necessary environment variables on Terraform cloud.

Then install `pipenv` as package management, for installing the dependencies:

```sh
pipenv install
```

Run `synth` command to verify if the source code is ready.

```sh
cdktf synth
```

Start deploy the infrastructure

```sh
cdktf deploy --auto-approve
```

Clean up command if no longer in use.

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
