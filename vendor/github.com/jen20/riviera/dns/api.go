package dns

import "fmt"

const apiVersion = "2015-05-04-preview"
const apiProvider = "Microsoft.Network"

func dnsZoneDefaultURLPathFunc(resourceGroupName, dnsZoneName string) func() string {
	return func() string {
		return fmt.Sprintf("resourceGroups/%s/providers/%s/dnsZones/%s", resourceGroupName, apiProvider, dnsZoneName)
	}
}
