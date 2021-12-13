---
page_title: "Signalwire SIP Endpoint"
subcategory: "SIP"
---

# signalwire_sip_endpoint Resource

Manages a Signalwire SIP endpoint. See the [API docs](https://docs.signalwire.com/topics/relay-rest/#resources-sip-endpoints) for more information

## Example Usage

```hcl
resource "signalwire_sip_endpoint" "test_endpoint" {
  space = "your_space"
  username = "c3p0"
  password = "password"
  caller_id = "C-3P0"
  ciphers = [
    "AEAD_AES_256_GCM_8",
    "AES_256_CM_HMAC_SHA1_80",
    "AES_CM_128_HMAC_SHA1_80",
    "AES_256_CM_HMAC_SHA1_32",
    "AES_CM_128_HMAC_SHA1_32"
  ]
  codecs = [
    "OPUS",
    "G722",
    "PCMU",
    "PCMA",
    "VP8",
    "H264"
  ]
  encryption = "optional"
}
```

## Argument Reference

The following arguments are supported:

- `space` - (Mandatory) The space your SIP endpoint will be created in (https://your-space.signalwire.com)
- `username` - (Mandatory) String representing the username portion of the endpoint. Must be unique across your project and must not container white space characters or @.
- `password` - (Mandatory) A password to authenticate registrations to this endpoint.
- `caller_id` - (Optional) Friendly Caller ID used as the CNAM when dialing a phone number or the From when dialing another SIP Endpoint.
- `ciphers` - (Optional) A list of encryption ciphers this endpoint will support.
- `codecs` - (Optional) A list of codecs this endpoint will support.
- `encryption` - (Optional) A string representing whether connections to this endpoint require encryption or if encryption is optional. Encryption will always be used if possible.
