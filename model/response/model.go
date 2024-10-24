package response

type ListRoleRsp []*RoleRsp

type RoleRsp struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
