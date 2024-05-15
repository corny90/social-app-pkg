package behaviour

import (
	"github.com/gocql/gocql"
	"time"
)

type TrackRecord struct {
	UserID        gocql.UUID `json:"userId"`
	PostID        gocql.UUID `json:"postId"`
	ActionView    bool       `json:"actionView"`
	ActionLike    bool       `json:"actionLike"`
	ActionComment bool       `json:"actionComment"`
	ViewDuration  int32      `json:"viewDuration"`
	Clicks        string     `json:"clicks"`
	CreatedAt     time.Time  `json:"createdAt"`
	Keywords      []string   `json:"keywords"`
	Score         int        `json:"score"`
}

type KeywordStats struct {
	Keyword    string `json:"keyword"`
	TotalViews int32  `json:"totalViews"`
	TotalTime  int32  `json:"totalTime"`
}

type StatisticData struct {
	TopSpentTime []TrackRecord  `json:"topRecords"`
	TopAccessed  []TrackRecord  `json:"topPosts"`
	TopKeywords  []KeywordStats `json:"topKeywords"`
}
