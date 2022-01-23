package types

type LoginDTO struct {
	ApplicationType int    `json:"applicationType"`
	UserName        string `json:"email"`
	Password        string `json:"password"`
}

type LoinResultUserDTO struct {
	ID    int `json:"id"`
	OrgID int `json:"orgId"`
}

// LoinResultDTO -- login api response dto
type LoinResultDTO struct {
	Message string            `json:"message"`
	Token   string            `json:"token"`
	User    LoinResultUserDTO `json:"user"`
}

type UploadSTSDTO struct {
	OSSKey          string `json:"ossKey"`
	OSSRegion       string `json:"ossRegion"`
	OSSUri          string `json:"ossUri"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken   string `json:"securityToken"`
	ExpirationMS    int    `json:"expirationMS"`
}
