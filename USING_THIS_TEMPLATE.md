# Using This Template

## Commitizen

[Commitizen](https://commitizen-tools.github.io/commitizen/)

After you have created a copy of this template and have installed Commitizen. It is recommended to use [pipx](https://github.com/pypa/pipx) to install Commitizen.

`pipx install commitizen`

Then run

`commitizen init`

Answer the prompts and it should generate a Commitizen config file. When the promt asks you if you want to install the pre-commit hooks select no. For projects that have Go and Terratests I recommend `.cz.toml`.

## Pre-Commit

[Pre-Commit](https://pre-commit.com/)

It is recommended to use pipx to install Pre-commit.

`pipx install pre-commit`

Then run

`pre-commit autoupdate`

To update your pre-commit hooks to the most recent version.

Then run

`pre-commit install --hook-type commit-msg`

This will check that all commit messages are Conventional Commit/Commitizen friendly. This ensures that you can use Commitizen to both maintain semantic versioning and to generate a CHANGELOG.md from those commit messages; `cz bump --changelog`.

## terraform_docs

Install [terraform-docs](https://terraform-docs.io/).

Go: `go install github.com/terraform-docs/terraform-docs@v0.16.0`

Homebrew: `brew install terraform-docs`

## tfsec

Install [tfsec](https://github.com/aquasecurity/tfsec) or comment out the line in the `.pre-commit-config.yaml`.

Go: `go install github.com/aquasecurity/tfsec/cmd/tfsec@latest`

Hombrew: `brew install tfsec`

## tfswitch or tfenv

I recommend using [tfswitch]() or [tfenv]() to manage your Terraform versions. It is important that all developers interacting with the same state file are using the same version of Terraform.

create a `.terraform-version` file in the root of your project with the Terraform version that the tfstate file uses:

### .terraform-version

```
1.1.9
```

## Terratest

This module uses [Terratest](https://terratest.gruntwork.io/) to test the Terraform module. Change dirctory to `./tests/src` and run `go mod download` to install the test dependencies. To run the tests from the `./tests/src` directory run `go test`.

## Read the Docs

Read the Terratest, Commitizen, and Pre-commit docs. Write good tests. Use the pre-commit hooks. Be safe, be happy.

**delete after reading**
