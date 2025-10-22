package provisionclient

type Resource struct {
	ID          PVID              `json:"id"`
	ParentID    PVID              `json:"parent_id"`
	Name        string            `json:"name"`
	Slug        string            `json:"slug,omitempty"`
	Type        string            `json:"type"`
	Date        string            `json:"date"`
	Modified    string            `json:"modified"`
	Attrs       map[string]string `json:"attrs,omitempty"`
	Permissions map[string]string `json:"permissions,omitempty"`
	Section     map[string]string `json:"section,omitempty"`
	Linkages    []Linkage         `json:"linkages,omitempty"`
}

type Resource_json struct {
	ID          PVID        `json:"id"`
	ParentID    PVID        `json:"parent_id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug,omitempty"`
	Type        string      `json:"type"`
	Date        string      `json:"date"`
	Modified    string      `json:"modified"`
	Attrs       interface{} `json:"attrs,omitempty"`
	Permissions interface{} `json:"permissions,omitempty"`
	Section     interface{} `json:"section,omitempty"`
	Linkages    []Linkage   `json:"linkages,omitempty"`
}

type Linkage struct {
	ID          PVID   `json:"id,omitempty"`
	ResourceID1 PVID   `json:"resource_id1,omitempty"`
	ResourceID2 PVID   `json:"resource_id2,omitempty"`
	Relation    string `json:"relation,omitempty"`
	Data        string `json:"data,omitempty"`
}
