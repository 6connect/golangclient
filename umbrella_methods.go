package provisionclient

import (
	"encoding/json"
	"net/url"
)

type UmbrellaMethods struct {
	Client *Client
}

//	mods, err := client.Umbrella.GetModules(&map[string]string{
//		"offset": "2",
//		"limit": "10",
//	})
func (umbrella *UmbrellaMethods) GetModules(filters *map[string]string) ([]UmbrellaModule, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}
	body, err := umbrella.Client.doRequest("GET", "/umbrella/modules"+fquery, nil)
	if err != nil {
		return nil, err
	}

	modules_ret := []UmbrellaModule{}
	err = json.Unmarshal(body, &modules_ret)
	if err != nil {
		return nil, err
	}

	return modules_ret, nil
}

//	unbs, err := client.Umbrella.GetNetblocks(&map[string]string{
//		"cidr": "192.168.0.1/24",
//	    "module_id": "123",
//	})
func (umbrella *UmbrellaMethods) GetNetblocks(filters *map[string]string) ([]UmbrellaNetblock, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}
	body, err := umbrella.Client.doRequest("GET", "/umbrella/netblocks"+fquery, nil)
	if err != nil {
		return nil, err
	}

	netblocks_ret := []UmbrellaNetblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return netblocks_ret, nil
}

// values, err := client.Umbrella.GetModuleLinks(123, &map[string]string{})
func (umbrella *UmbrellaMethods) GetModuleLinks(moduleID string, filters *map[string]string) ([]ResourceLink, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}
	body, err := umbrella.Client.doRequest("GET", "/umbrella/modules/"+moduleID+"/links"+fquery, nil)
	if err != nil {
		return nil, err
	}

	resource_links_json := []ResourceLink_json{}
	err = json.Unmarshal(body, &resource_links_json)
	if err != nil {
		return nil, err
	}

	resource_links_ret := make([]ResourceLink, len(resource_links_json))
	for k, resource_link := range resource_links_json {
		resource_links_ret[k] = ResourceLink{
			ID:         resource_link.ID,
			ResourceID: resource_link.ResourceID,
			Relation:   resource_link.Relation,
			Data:       resource_link.Data,
			Resource:   resourceJsonToResource(resource_link.Resource),
		}

	}

	return resource_links_ret, nil
}
