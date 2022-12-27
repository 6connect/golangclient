package golangsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

type IPAMMethods struct {
	Client *Client
}

// values, err := client.IPAM.SmartAssign("799399", "ipv4", "1918", 28, map[string]interface{}{})
func (ipam *IPAMMethods) SmartAssign(resource_id, ip_type, rir string, mask int, params map[string]interface{}) (*Netblock, error) {
	params["resource_id"] = resource_id
	params["type"] = ip_type
	params["rir"] = rir
	params["mask"] = mask

	reqbody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err := ipam.Client.doRequest("PUT", "/ipam/netblocks/smart_assign", bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_record Netblock
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

// values, err := client.IPAM.DirectAssign("799399", "192.168.1.0/24", "1918", map[string]interface{}{})
func (ipam *IPAMMethods) DirectAssign(resource_id, cidr, rir string, params map[string]interface{}) (*Netblock, error) {
	params["resource_id"] = resource_id
	params["rir"] = rir
	params["cidr"] = cidr

	reqbody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err := ipam.Client.doRequest("PUT", "/ipam/netblocks/smart_assign", bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_record Netblock
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

//	values, err := client.IPAM.GetNetblocks(&map[string]string{
//		"cidr": "192.168.192.176/28",
//	})
func (ipam *IPAMMethods) GetNetblocks(filters *map[string]string) ([]Netblock, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}
	body, err := ipam.Client.doRequest("GET", "/ipam/netblocks"+fquery, nil)
	if err != nil {
		return nil, err
	}

	netblocks_ret := []Netblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return netblocks_ret, nil
}

func (ipam *IPAMMethods) GetNetblockByID(netblock_id string) (*Netblock, error) {

	body, err := ipam.Client.doRequest("GET", "/ipam/netblocks/"+netblock_id, nil)
	if err != nil {
		return nil, err
	}

	netblocks_ret := Netblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return &netblocks_ret, nil
}

func (ipam *IPAMMethods) GetNetblockByCIDR(cidr string) (*Netblock, error) {

	body, err := ipam.Client.doRequest("GET", "/ipam/netblocks/"+cidr, nil)
	if err != nil {
		return nil, err
	}

	netblocks_ret := Netblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return &netblocks_ret, nil
}

func (ipam *IPAMMethods) AddNetblock(netblock Netblock) (*Netblock, error) {
	reqbody, err := json.Marshal(netblock)
	if err != nil {
		return nil, err
	}

	body, err := ipam.Client.doRequest("POST", "/ipam/netblocks", bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	netblocks_ret := Netblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return &netblocks_ret, nil
}

func (ipam *IPAMMethods) UpdateNetblock(netblock Netblock) (*Netblock, error) {
	reqbody, err := json.Marshal(netblock)
	if err != nil {
		return nil, err
	}

	body, err := ipam.Client.doRequest("POST", "/ipam/netblocks/"+string(netblock.ID), bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	netblocks_ret := Netblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return &netblocks_ret, nil
}

func (ipam *IPAMMethods) DeleteNetblock(netblock Netblock) (*Netblock, error) {
	return ipam.Client.IPAM.DeleteNetblockByID(string(netblock.ID))
}

func (ipam *IPAMMethods) DeleteNetblockByID(netblock_id string) (*Netblock, error) {

	body, err := ipam.Client.doRequest("DELETE", "/ipam/netblocks/"+netblock_id, nil)
	if err != nil {
		return nil, err
	}

	netblocks_ret := Netblock{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	return &netblocks_ret, nil
}

func (ipam *IPAMMethods) GetFirstAvailableByNetblock(netblock Netblock) (*string, error) {
	if string(netblock.ID) != "" {
		return ipam.GetFirstAvailable(string(netblock.ID))
	} else if string(netblock.CIDR) != "" {
		return ipam.GetFirstAvailable(string(netblock.CIDR))
	}

	return nil, errors.New("either ID or CIDR must be provided in the Netblock")
}

func (ipam *IPAMMethods) GetFirstAvailable(cidr_or_id string) (*string, error) {

	body, err := ipam.Client.doRequest("GET", "/ipam/netblocks/"+cidr_or_id+"/first_available", nil)
	if err != nil {
		return nil, err
	}

	netblocks_ret := map[string]string{}
	err = json.Unmarshal(body, &netblocks_ret)
	if err != nil {
		return nil, err
	}

	ret := netblocks_ret["cidr"]

	split := strings.Split(netblocks_ret["cidr"], "/")
	if len(split) == 2 && (split[1] == "32" || split[1] == "128") {
		ret = split[1]
	}

	return &ret, nil
}