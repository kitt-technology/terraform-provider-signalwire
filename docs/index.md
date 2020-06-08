# Signalwire Provider

The Signalwire provider is used to interact with the many resources supported by Signalwire. The provider needs to be configured with your Twilio credentials before it can be used.

> ⚠️ **Disclaimer**: This project is not an official Signalwire project and is not supported or endorsed by Signalwire in any way.

## Authentication

The following authentication methods are supported, in precedence order:

- Environment variables
    - Project ID & Auth Token
  
#### Project ID & Auth Token

### Environment variables

#### Project ID & Auth Token

You can provide your credentials via the `SIGNALWIRE_PROJECT_ID` and `SIGNALWIRE_AUTH_TOKEN` environment variables, representing your Signalwire Project ID and Auth Token respectively.

```hcl
provider "signalwire" {}
```

Usage:

```sh
export SIGNALWIRE_PROJECT_ID="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export SIGNALWIRE_AUTH_TOKEN="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
terraform plan
```

or

```sh
SIGNALWIRE_PROJECT_ID="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" SIGNALWIRE_AUTH_TOKEN="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" terraform plan
```
