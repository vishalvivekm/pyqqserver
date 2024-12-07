package model

import "time"

const (
	BTECH = "btech"
)

var (
	SemesterList = map[string]string{
		"First Semester": "firstsemesters",
		"Second Semester": "secondsemesters",
	}
	SemListNew = map[string]int{
		"firstsemesters": 1,
		"secondsemesters": 2,
	}
	BranchList = map[string]string{
		"GCS": "GCS",
		"GCE": "GCE",
		"GEE": "GEE",
	}
)


type Subject struct {
	ID       string   `bson:"_id"`
	Name     string   `bson:"name"`
	Branches []string `bson:"branches"`
	Semester int      `bson:"semester"`
	SubjectID string `bson:"subjectID"` // subject name but in kebab case
}

type SubjectDetail struct {
	ID             string        `json:"_id"`
	Subject        string        `json:"subject"`
	Dept           []string      `json:"dept"`
	TheoryPaperCode string        `json:"theorypapercode"`
	LabPaperCode   string        `json:"labpapercode,omitempty"`
	TheoryCRedits  int           `json:"theorycredits"`
	LabCredits     int           `json:"labcredits,omitempty"`
	Theory         []TheoryUnit  `json:"theory,omitempty"`
	Lab []LabDetail  `json:"lab"`
}
type LabDetail struct {
	Experiment int `json:"experiment"`
	Aim AimDetail `json:"aim"`
}
type AimDetail struct {
	Objective string `json:"objective"`
	Steps []string `json:"steps"`
}

type Notes []Common
type PYQs []Common
type Books []Common

type Common struct {
	WebViewLink string    `json:"webViewLink"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"createdTime"`
}

type TheoryUnit struct {
	Unit   int      `json:"unit"`
	Topics []string `json:"topics"`
}
