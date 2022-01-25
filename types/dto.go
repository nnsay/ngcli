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

type CreatedUploadDTO struct {
	Upload struct {
		ID int `json:"id"`
	} `json:"upload"`
}
type CreatedJobDTO struct {
	ID              int    `json:"id"`
	SubjectID       int    `json:"subjectId"`
	ProjectID       int    `json:"projectId"`
	CreatedByUserID int    `json:"createdByUserId"`
	JobType         int    `json:"jobType"`
	Status          int    `json:"status"`
	CreatedAt       int64  `json:"createdAt"`
	RequestRef      string `json:"requestRef"`
	UploadID        int    `json:"uploadId"`
	JobInfo         struct {
		ExperimentInfo struct {
			CommonInfo struct {
				SubjectCustID string `json:"subjectCustId"`
			} `json:"commonInfo"`
			ExperimentType   int `json:"experimentType"`
			ConfirmStatus    int `json:"confirmStatus"`
			TimeofRepetition int `json:"timeofRepetition"`
		} `json:"experimentInfo"`
	} `json:"jobInfo"`
}
