---
page_title: "Signalwire Domain App"
subcategory: "SIP"
---

# domain_app Resource

Manages a Signalwire Domain App. See the [API docs](https://docs.signalwire.com/topics/relay-rest/#resources-domain-applications) for more information

## Example Usage

```hcl
resource "signalwire_domain_app" "test_app" {
  space = "your_space"
  name = "your_app"
  identifier = "some_id"
  ip_auth_enabled = true
  ip_auth = ["8.8.8.8", "4.4.4.4"]
  encryption = "required"
  call_handler = "laml_webhooks"
  call_relay_context = "incoming"
  ciphers = [
    "AEAD_AES_256_GCM_8",
  ]
  codecs = [
    "OPUS",
    "G722",
    "PCMU",
    "PCMA"
  ]
}
```

## Argument Reference

The following arguments are supported:

- `space` - (Mandatory) The space your SIP endpoint will be created in (https://your-space.signalwire.com)
- `name` - (Mandatory) A string representing the friendly name for this domain application.
- `identifier` - (Mandatory) A string representing the identifier portion of the domain application. Must be unique across your project.
- `call_handler` - (Mandatory) A string representing how this domain application handles calls. Valid values are relay_context, laml_webhooks, and laml_application.
- `call_request_url` - (Optional) A string representing the LaML URL to access when a call is received. This is only used when call_handler is set to laml_webhooks.
- `call_request_method` - (Optional) A string representing the HTTP method to use with call_request_url. Valid values are GET and POST. This is only used when call_handler is set to laml_webhooks.
- `call_fallback_url` - (Optional) A string representing the LaML URL to access when the call to call_request_url fails. This is only used when call_handler is set to laml_webhooks.
- `call_fallback_method` - (Optional) A string representing the HTTP method to use with call_fallback_url. Valid values are GET and POST. This is only used when call_handler is set to laml_webhooks.
- `call_status_callback_url` - (Optional) A string representing a URL to send status change messages to. This is only used when call_handler is set to laml_webhooks.
- `call_status_callback_method` - (Optional) A string representing the HTTP method to use with call_status_callback_url. Valid values are GET and POST. This is only used when call_handler is set to laml_webhooks.
- `call_relay_context` - (Optional) A string representing the Relay context to forward incoming calls to. This is only used when call_handler is set to relay_context.
- `call_laml_application_id` - (Optional) A string representing the ID of the LaML application to forward incoming calls to. This is only used when call_handler is set to laml_application.
- `ip_auth_enabled` - (Optional) Whether the domain application will enforce IP authentication for incoming requests.
- `ip_auth` - (Optional) A list containing whitelisted IP addresses and IP blocks used if ip_auth_enabled is true.
- `encryption` - (Optional) A string representing whether connections to this endpoint require encryption or if encryption is optional. Encryption will always be used if possible.
- `ciphers` - (Optional) A list of encryption ciphers this endpoint will support.
- `codecs` - (Optional) A list of codecs this endpoint will support.
