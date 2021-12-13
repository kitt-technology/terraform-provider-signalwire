package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSignalwireDomainApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSignalwireDomainAppCreate,
		Read:   resourceSignalwireDomainAppRead,
		Update: resourceSignalwireDomainAppUpdate,
		Delete: resourceSignalwireDomainAppDelete,
		Schema: map[string]*schema.Schema{
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The space your SIP endpoint will be created in (https://your-space.signalwire.com)",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string representing the friendly name for this domain application.",
			},
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string representing the identifier portion of the domain application. Must be unique across your project.",
			},
			"ip_auth_enabled": {
				Type:        schema.TypeBool,
				Description: "Whether the domain application will enforce IP authentication for incoming requests.",
				Optional: true,
			},
			"ip_auth": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list containing whitelisted IP addresses and IP blocks used if ip_auth_enabled is true.",
				Optional: true,
			},
			"call_handler": {
				Type:        schema.TypeString,
				Description: "A string representing how this domain application handles calls. Valid values are relay_context, laml_webhooks, and laml_application.",
				Required: true,
			},
			"call_request_url": {
				Type:        schema.TypeString,
				Description: "A string representing the LaML URL to access when a call is received. This is only used when call_handler is set to laml_webhooks.",
				Optional: true,
			},
			"call_request_method": {
				Type:        schema.TypeString,
				Description: "A string representing the HTTP method to use with call_request_url. Valid values are GET and POST. This is only used when call_handler is set to laml_webhooks.",
				Optional: true,
				Default: "POST",
			},
			"call_fallback_url": {
				Type:        schema.TypeString,
				Description: "A string representing the LaML URL to access when the call to call_request_url fails. This is only used when call_handler is set to laml_webhooks.",
				Optional: true,
			},
			"call_fallback_method": {
				Type:        schema.TypeString,
				Description: "A string representing the HTTP method to use with call_fallback_url. Valid values are GET and POST. This is only used when call_handler is set to laml_webhooks.",
				Optional: true,
				Default: "POST",
			},
			"call_status_callback_url": {
				Type:        schema.TypeString,
				Description: "A string representing a URL to send status change messages to. This is only used when call_handler is set to laml_webhooks.",
				Optional: true,
			},
			"call_status_callback_method": {
				Type:        schema.TypeString,
				Description: "A string representing the HTTP method to use with call_status_callback_url. Valid values are GET and POST. This is only used when call_handler is set to laml_webhooks.",
				Optional: true,
				Default: "POST",
			},
			"call_relay_context": {
				Type:        schema.TypeString,
				Description: "A string representing the Relay context to forward incoming calls to. This is only used when call_handler is set to relay_context.",
				Optional: true,
			},
			"call_laml_application_id": {
				Type:        schema.TypeString,
				Description: "A string representing the ID of the LaML application to forward incoming calls to. This is only used when call_handler is set to laml_application.",
				Optional: true,
			},
			"encryption": {
				Type:        schema.TypeString,
				Description: "A string representing whether connections to this domain application require encryption or if encryption is optional. Encryption will always be used if possible. Valid values are optional and required.",
				Required: true,
			},
			"ciphers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "A list of encryption ciphers this endpoint will support.",
			},
			"codecs": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "A list of codecs this endpoint will support.",
			},
		},
	}
}

func resourceSignalwireDomainAppCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	resp, err := client.Req("POST", d.Get("space").(string), "domain_applications", schemaToDomainApp(d))

	if err != nil {
		return err
	}

	if _, ok := resp["id"]; !ok {
		return fmt.Errorf("no id on response: %v", resp)
	}

	d.SetId(resp["id"].(string))

	return resourceSignalwireDomainAppRead(d, meta)
}

func resourceSignalwireDomainAppRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	resp, err := client.Req("GET", d.Get("space").(string), "domain_applications/"+d.Id(), nil)

	if err != nil {
		return err
	}

	respToDomainAppData(d, resp)
	return nil
}

func resourceSignalwireDomainAppUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	resp, err := client.Req("PUT", d.Get("space").(string), "domain_applications/"+d.Id(), schemaToDomainApp(d))

	if err != nil {
		return err
	}

	respToDomainAppData(d, resp)
	return nil
}

func resourceSignalwireDomainAppDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	_, err := client.Req("DELETE", d.Get("space").(string), "domain_applications/"+d.Id(), nil)

	return err
}

func schemaToDomainApp(d *schema.ResourceData) map[string]interface{} {
	request := map[string]interface{}{
		"name": d.Get("name").(string),
		"identifier": d.Get("identifier").(string),
		"ip_auth_enabled": d.Get("ip_auth_enabled").(bool),
		"call_handler": d.Get("call_handler").(string),
		"call_relay_context": d.Get("call_relay_context").(string),
		"call_laml_application_id": d.Get("call_laml_application_id").(string),
		"encryption": d.Get("encryption").(string),
	}

	optionalsStr := []string{"call_request_url", "call_request_method", "call_fallback_url",  "call_laml_application_id", "call_fallback_method", "call_status_callback_url", "call_status_callback_method", "call_relay_context", "encryption"}
	for _, optional := range optionalsStr {
		if val, ok := d.GetOkExists(optional); ok {
			request[optional] = val.(string)
		}
	}
	optionalsSlices := []string{"ip_auth", "ciphers", "codecs"}
	for _, optionalSlice := range optionalsSlices {
		if val, ok := d.GetOkExists(optionalSlice); ok {
			var strSlice []string
			for _, strVal := range val.([]interface{}) {
				strSlice = append(strSlice, strVal.(string))
			}
			request[optionalSlice] = strSlice
		}
	}
	return request
}

func respToDomainAppData(d *schema.ResourceData, resp map[string]interface{}) {
	d.Set("id", resp["id"].(string))
	d.Set("type", resp["type"].(string))
	d.Set("name", resp["name"].(string))
	d.Set("identifier", resp["identifier"].(string))
	d.Set("domain", resp["domain"].(string))
	d.Set("ip_auth_enabled", resp["ip_auth_enabled"].(bool))
	d.Set("call_handler", resp["call_handler"].(string))

	optionalStr := []string{
		"call_request_url",
		"call_request_method",
		"call_fallback_url",
		"call_fallback_method",
		"call_status_callback_url",
		"call_status_callback_method",
		"call_relay_context",
		"call_laml_application_id",
	}
	for _, key := range optionalStr {
		if val, ok := resp[key]; ok && val != nil {
			d.Set(key, resp[key].(string))
		}
	}

	d.Set("encryption", resp["encryption"].(string))

	slices := []string{"ip_auth", "ciphers", "codecs"}
	for _, slice := range slices {
		var strSlice []string
		for _, strVal := range resp[slice].([]interface{}) {
			strSlice = append(strSlice, strVal.(string))
		}
		d.Set(slice, strSlice)
	}
}
