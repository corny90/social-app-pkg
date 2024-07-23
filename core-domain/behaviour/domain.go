package behaviour

import (
	"encoding/json"
	"time"

	"github.com/gocql/gocql"
)

type TrackRecord struct {
	UserID        gocql.UUID `json:"userId"`
	PostID        gocql.UUID `json:"postId"`
	ActionView    bool       `json:"actionView"`
	ActionLike    bool       `json:"actionLike"`
	ActionComment bool       `json:"actionComment"`
	ViewDuration  int        `json:"viewDuration"`
	ViewCount     int        `json:"viewCount"`
	Clicks        Clicks     `json:"clicks"`
	CreatedAt     time.Time  `json:"createdAt"`
	Keywords      []string   `json:"keywords"`
	Score         int        `json:"score"`
	TrackType     string     `json:"trackType"`
}

type Clicks struct {
	PostUser bool `json:"post_user"`
	Views    bool `json:"views"`
	Likes    bool `json:"likes"`
	Comments bool `json:"comments"`
}

type KeywordStats struct {
	Keyword    string `json:"keyword"`
	TotalViews int    `json:"totalViews"`
	TotalTime  int    `json:"totalTime"`
}

type StatisticData struct {
	TopSpentTime []TrackRecord  `json:"topRecords"`
	TopAccessed  []TrackRecord  `json:"topPosts"`
	TopKeywords  []KeywordStats `json:"topKeywords"`
}

type TrackRequest struct {
	UserID string `json:"userId"`
}

type TrackResponse struct {
	UserID string          `json:"userId"`
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}
