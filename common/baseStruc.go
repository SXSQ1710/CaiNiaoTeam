package common

// Response 统一响应结构
type Response struct {
	StatusCode int32  `json:"status_code"`          //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态描述
}

type Video struct {
	Id            int64 `json:"id,omitempty" ;gorm:"primary_key;AUTO_INCREMENT"`
	AuthorId      int64
	Author        User   `json:"author" ;gorm:"foreignKey:AuthorId;references:Id;"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"` //通过本机网关访问本地文件
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty " ;gorm:"size:32;"`
}

type View_video_favorites struct {
	Id            int64 `json:"id,omitempty" ;gorm:"primary_key;AUTO_INCREMENT"`
	AuthorId      int64
	Author        User   `json:"author" ;gorm:"foreignKey:AuthorId;references:Id;"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"` //通过本机网关访问本地文件
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty " ;gorm:"size:32;"`
}

type Comment struct {
	Id         int64 `json:"id,omitempty" ;gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int64
	VideoId    int64
	User       User   `json:"user" ;gorm:"foreignKey:UserId;references:Id;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type Favorite struct {
	User_id  string `gorm:"primary_key"`
	Video_id string `gorm:"primary_key"`
}

type Relation struct {
	User_id        string `gorm:"primary_key"`
	Follow_user_id string `gorm:"primary_key"`
	Follow_user    User   `gorm:"foreignKey:Follow_user_id;references:Id;"` //被关在的人
	Follower_user  User   `gorm:"foreignKey:User_id;references:Id;"`        //关注的人
}

type User struct {
	Id            int64  `json:"id,omitempty" ;gorm:"primary_key;AUTO_INCREMENT"` //用户ID，自增
	IdPass        string `gorm:"size:64"`                                         //用户鉴权，唯一标识用户,现阶段数据库保存token
	Name          string `json:"name,omitempty" ;gorm:"size:32"`                  //用户名
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func (video *View_video_favorites) SetPlayUrl(playUrl string) {
	video.PlayUrl = playUrl
}

func (video *View_video_favorites) SetCoverUrl(coverUrl string) {
	video.CoverUrl = coverUrl
}
