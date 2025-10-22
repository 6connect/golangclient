package provisionclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
)

type ResourceMethods struct {
	Client *Client
}

//	client.DNS.GetResources(&map[string]string{
//		"offset": "2",
//		"limit": "10",
//	})
func (resources *ResourceMethods) GetResources(filters *map[string]string) ([]Resource, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}
	body, err := resources.Client.doRequest("GET", "/resources"+fquery, nil)
	if err != nil {
		return nil, err
	}

	resources_ret_json := []Resource_json{}

	err = json.Unmarshal(body, &resources_ret_json)
	if err != nil {
		return nil, err
	}
	resources_ret := make([]Resource, len(resources_ret_json))
	for k, resource := range resources_ret_json {
		resources_ret[k] = resourceJsonToResource(resource)
	}

	return resources_ret, nil
}

//	newResource := provisionclient.Resource{
//		Name:     "Test A Record",
//		Type:     "dnsrecord",
//		ParentID: "428964",
//		Attrs: map[string]string{
//			"record_host":  "aapi.example.com.",
//			"record_value": "1.1.1.1",
//			"record_type":  "A",
//		},
//	}
//
// values, err := client.Resources.AddResource(newResource)
func (resources *ResourceMethods) AddResource(resource Resource) (*Resource, error) {
	reqbody, err := json.Marshal(resource)
	if err != nil {
		return nil, err
	}

	body, err := resources.Client.doRequest("POST", "/resources", bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_resource Resource
	err = json.Unmarshal(body, &resp_resource)
	if err != nil {
		return nil, err
	}

	return &resp_resource, nil
}

func (resources *ResourceMethods) UpdateResource(resource Resource) (*Resource, error) {
	if string(resource.ID) == "" {
		return nil, errors.New("error: Resource ID is required for updating a resource")
	}

	reqbody, err := json.Marshal(resource)
	if err != nil {
		return nil, err
	}

	body, err := resources.Client.doRequest("PATCH", "/resources/"+string(resource.ID), bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_resource Resource
	err = json.Unmarshal(body, &resp_resource)
	if err != nil {
		return nil, err
	}

	return &resp_resource, nil
}

func (resources *ResourceMethods) DeleteResource(resource Resource) error {
	return resources.Client.Resources.DeleteResourceByID(string(resource.ID))
}

func (resources *ResourceMethods) DeleteResourceByID(resourceId string) error {
	if string(resourceId) == "" {
		return errors.New("error: Resource ID is required for updating a resource")
	}

	_, err := resources.Client.doRequest("DELETE", "/resources/"+resourceId, nil)
	return err
}

func (resources *ResourceMethods) GetResourceByID(resourceId string) (*Resource, error) {
	body, err := resources.Client.doRequest("GET", "/resources/"+resourceId, nil)
	if err != nil {
		return nil, err
	}
	resources_ret_json := Resource_json{}

	err = json.Unmarshal(body, &resources_ret_json)
	if err != nil {
		return nil, err
	}

	resources_ret := resourceJsonToResource(resources_ret_json)

	return &resources_ret, nil
}

func resourceJsonToResource(rj Resource_json) Resource {
	resource := Resource{
		ID:       rj.ID,
		ParentID: rj.ParentID,
		Name:     rj.Name,
		Slug:     rj.Slug,
		Modified: rj.Modified,
		Date:     rj.Date,
		Linkages: rj.Linkages,
	}

	switch dv := rj.Attrs.(type) {
	case map[string]interface{}:
		if len(dv) != 0 {
			m := make(map[string]string)
			for k, v := range dv {
				switch val := v.(type) {
				case string:
					m[k] = string(val)
				}
			}
			resource.Attrs = m
		}
	}

	switch dv := rj.Permissions.(type) {
	case map[string]interface{}:
		if len(dv) != 0 {
			m := make(map[string]string)
			for k, v := range dv {
				switch val := v.(type) {
				case string:
					m[k] = string(val)
				}
			}
			resource.Permissions = m
		}
	}

	switch dv := rj.Section.(type) {
	case map[string]interface{}:
		if len(dv) != 0 {
			m := make(map[string]string)
			for k, v := range dv {
				switch val := v.(type) {
				case string:
					m[k] = string(val)
				}
			}
			resource.Section = m
		}
	}

	return resource
}
