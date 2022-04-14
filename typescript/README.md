## Install

- This project is configured to store the state on Terraform Cloud - under Nimble organization,
  https://github.com/nimblehq/terraform-cloud-development-kit/blob/14eb4997c5a904f4aacb0b50ede39c4c5eb75459/typescript/main.ts#L12-L18

  Make sure the variables that hold AWS access key are available on Terraform Cloud. The access key can be generated on AWS Console. ([more info](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html#Using_CreateAccessKey))
  - AWS_ACCESS_KEY_ID
  - AWS_SECRET_ACCESS_KEY  

- Install the project on your local:
  ```bash
  $ npm install
  ```

- Run `synth` command to verify if the source code is ready.

  ```sh
  cdktf synth
  ```

- Start deploy the infrastructure

  ```sh
  cdktf deploy
  ```

- Deploy the infrastructure.

  ```sh
  cdktf destroy
  ```

## Docs

- https://learn.hashicorp.com/tutorials/terraform/cdktf-build
- https://www.terraform.io/cdktf
