# STACKIT Secrets-Manager CLI

This project provides the command line interface to the STACKIT Secrets-Manager.

## Installation

The recommended way for installing the STACKIT Secrets-Manager CLI is to
[download prebuilt binaries](https://github.com/stackitcloud/stackit-secrets-manager-cli/releases).

Alternatively you can build it from source, if you have Go already installed:

```shell
$ go install github.com/stackitcloud/stackit-secrets-manager-cli/cmd/stackit-secrets-manager@latest
```

## Usage

Create an access token for the STACKIT project you want to interact with. The token needs at least project.member permissions.  
Now set the token and the project id with the `configure` subcommand:

```shell
$ stackit-secrets-manager configure
Authentication Token []: <your token>
Project UUID []:  <your project id>
Configuration successfully written
```

Alternatively, you can set these settings as environment variables in cases where you might have a read-only file
system:

```shell
$ export AUTHENTICATION_TOKEN=eyJraWQiO...zQXuLFGP3hMfw
$ export PROJECT_ID=54349a42-2fbc-4fed-b307-16673e3eaa1f
```

To create a new instance run:

```shell
$ stackit-secrets-manager create instance --name 'test'
```

To get a list of all secrets manager instances run:

```shell
$ stackit-secrets-manager get instances
```

To create a user with write access for that instance:

```shell
$ stackit-secrets-manager create user --instance-id 0069066b-b7d2-4e04-bda8-0f3f02efb920 --enable-write
```

Note down the password which is printed there. This is the only place where you will ever see the password for the user.

To list all users for that instance:

```shell
$ stackit-secrets-manager get users --instance-id 0069066b-b7d2-4e04-bda8-0f3f02efb920
```

For more information about the available sub-commands and flags use the `--help` command line flag.

Use the API URL and Secrets Engine name of the instance and the username and password of the user to configure
[the Hashicorp Vault client](https://developer.hashicorp.com/vault/downloads) to interact with the secrets engine on the command line.

```shell
$ export VAULT_ADDR=https://prod.sm.eu01.stackit.cloud
$ vault login -method=userpass username=h86c6it5228nn9d9 password="A{o'61eJzD]|hUH4"
$ vault kv put 0069066b-b7d2-4e04-bda8-0f3f02efb920/foo bar=baz
$ vault kv get 0069066b-b7d2-4e04-bda8-0f3f02efb920/foo
```

The web UI for the secrets engine can be opened in any web browser with the API url of the instance as target
location (i.e. "https://prod.sm.eu01.stackit.cloud"). Choose the "Username" login method and enter the username
and password as given by the CLI command to log in.

## Development

If you want to work with the source code of the Secrets-Manager CLI, you need to match these prerequisites:

* Go v1.19 or newer
* Have `make` and `git` available on your system

### Building

To build the binary, simply run:

```shell
$ make
```

This will download dependencies into the `bin` subdirectory, run some linting and build the `stackit-secrets-manager`
executable in the project root.

### Updating Dependencies

To update third party libraries:

```shell
$ go get -u
$ go mod tidy
```

To update the openapi spec for the Secrets-Manager API:

```shell
$ make update-openapi-spec
```

To update third party binaries, check the shell scripts in the `scripts` subdirectory and update the versions
in there.

### Building a release

To build a new release, set the version environment variable to the desired version you want to release. Also
provide your GitHub token as environment variable as well. Make sure that all your changes are committed and
call the release make target:

```shell
$ export VERSION=0.0.1
$ export GITHUB_TOKEN=<your github token>
$ make release
```

You can also test out a release locally first, without publishing it to GitHub. The release is put into the `dist`
subdirectory:

```shell
$ make release-local
```
