// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/secrets/secrets.proto

package secrets

import policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.secrets.SecretsService/Create", "secrets", "create", "POST", "/secrets", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Secret); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "type":
					return m.Type
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.secrets.SecretsService/Read", "secrets:{id}", "read", "GET", "/secrets/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.secrets.SecretsService/Update", "secrets:{id}", "update", "PATCH", "/secrets/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Secret); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "type":
					return m.Type
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.secrets.SecretsService/Delete", "secrets:{id}", "delete", "DELETE", "/secrets/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.secrets.SecretsService/List", "secrets", "search", "POST", "/secrets/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
}
