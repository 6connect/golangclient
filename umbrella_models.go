package provisionclient

type UmbrellaModule struct {
	ID              PVID     `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	ParentID        PVID     `json:"parent_id,omitempty"`
	CategoryID      PVID     `json:"category_id,omitempty"`
	Date            string   `json:"date,omitempty"`
	Modified        string   `json:"modified,omitempty"`
	IpamBackend     string   `json:"ipam_backend,omitempty"`
	Username        string   `json:"username,omitempty"`
	Token           string   `json:"token,omitempty"`
	Url             string   `json:"url,omitempty"`
	Password        string   `json:"password,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
	AcpID           string   `json:"acp_id,omitempty"`
	UtilizationType string   `json:"utilization_type,omitempty"`
	AcpWorkflowSlug []string `json:"acp_workflow_slug,omitempty"`
}

type UmbrellaNetblock struct {
	ID                     PVID     `json:"id,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Description            string   `json:"description,omitempty"`
	ParentID               PVID     `json:"parent_id,omitempty"`
	CategoryID             PVID     `json:"category_id,omitempty"`
	Date                   string   `json:"date,omitempty"`
	Modified               string   `json:"modified,omitempty"`
	RemoteID               PVID     `json:"remote_id,omitempty"`
	Address                string   `json:"address,omitempty"`
	EndAddress             string   `json:"end_address,omitempty"`
	AllAddresses           string   `json:"all_addresses,omitempty"`
	CIDR                   string   `json:"cidr,omitempty"`
	IsAssigned             string   `json:"is_assigned,omitempty"`
	RIR                    string   `json:"rir,omitempty"`
	Customer               string   `json:"customer,omitempty"`
	SupportedActions       string   `json:"supported_actions,omitempty"`
	ModuleType             string   `json:"module_type,omitempty"`
	UsedAddresses          string   `json:"used_addresses,omitempty"`
	HasChildren            string   `json:"has_children,omitempty"`
	IsAggregate            string   `json:"is_aggregate,omitempty"`
	AddressHuman           string   `json:"address_human,omitempty"`
	EndAddressHuman        string   `json:"end_address_human,omitempty"`
	Mask                   string   `json:"mask,omitempty"`
	ModuleID               PVID     `json:"module_id,omitempty"`
	Utilization            string   `json:"utilization,omitempty"`
	IsSyncing              string   `json:"is_syncing,omitempty"`
	Status                 string   `json:"status,omitempty"`
	LastSyncTime           string   `json:"last_sync_time,omitempty"`
	LastError              string   `json:"last_error,omitempty"`
	Permissions            []string `json:"permissions,omitempty"`
	AggregateModule        string   `json:"aggregate_module,omitempty"`
	ModuleTypeAggregate    string   `json:"module_type_aggregate,omitempty"`
	UsedAddressesAggregate string   `json:"used_addresses_aggregate,omitempty"`
	Asn                    string   `json:"asn,omitempty"`
	RemoteType             string   `json:"remote_type,omitempty"`
	Data                   string   `json:"data,omitempty"`
	AllocatedAddresses     string   `json:"allocated_addresses,omitempty"`
	ModuleIdAggregate      string   `json:"module_id_aggregate,omitempty"`
}
