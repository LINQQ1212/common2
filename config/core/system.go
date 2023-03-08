package core

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 端口值
	AdminAddr     int    `mapstructure:"admin-addr" json:"admin-addr" yaml:"admin-addr"` // 端口值
	UserName      string `mapstructure:"username" json:"username" yaml:"username"`
	PassWord      string `mapstructure:"password" json:"password" yaml:"password"`
	MainDir       string `mapstructure:"main-dir" json:"main-dir" yaml:"main-dir"`
	ImageGrcp     string `mapstructure:"image-grcp" json:"image-grcp" yaml:"image-grcp"`
	TGToken       string `mapstructure:"tg-token" json:"tg-token" yaml:"tg-token"`
	TGChatId      string `mapstructure:"tg-chat-id" json:"tg-chat-id" yaml:"tg-chat-id"`
	Info          string `mapstructure:"info" json:"info" yaml:"info"`
	TestKey       string `mapstructure:"test-key" json:"test-key" yaml:"test-key"`
	HandleProduct int    `mapstructure:"handle-product" json:"handle-product" yaml:"handle-product"`
	JumpMethod    int    `mapstructure:"jump-method" json:"jump-method" yaml:"jump-method"`
	Scheme        string `mapstructure:"scheme" json:"scheme,omitempty" yaml:"scheme"`
	ListSize      int    `mapstructure:"list_size" json:"list_size,omitempty" yaml:"list_size"`
}

type NewVersionOption struct {
	TopDir string `json:"top_dir,omitempty" yaml:"top_dir" mapstructure:"top_dir"`

	RemoteCopy      bool   `json:"remote_copy,omitempty" yaml:"remote_copy" mapstructure:"remote_copy"` // 是否远程复制
	RemoteHost      string `json:"remote_host,omitempty" yaml:"remote_host" mapstructure:"remote_host"`
	RemotePort      string `json:"remote_port,omitempty" yaml:"remote_port" mapstructure:"remote_port"`
	RemoteUser      string `json:"remote_user,omitempty" yaml:"remote_user" mapstructure:"remote_user"`
	RemotePwd       string `json:"remote_pwd,omitempty" yaml:"remote_pwd" mapstructure:"remote_pwd"`
	RemoteEndRemove bool   `json:"remote_end_remove,omitempty" yaml:"remote_end_remove" mapstructure:"remote_end_remove"` // 结束同时删除远程服务器的文件

	AutoFilePath bool   `json:"auto_file_path,omitempty" yaml:"auto_file_path" mapstructure:"auto_file_path"`
	GoogleImg    string `json:"google_img,omitempty" yaml:"google_img" mapstructure:"google_img"`
	YahooDsc     string `json:"yahoo_dsc,omitempty" yaml:"yahoo_dsc" mapstructure:"yahoo_dsc"`
	BingDsc      string `json:"bing_dsc,omitempty" yaml:"bing_dsc" mapstructure:"bing_dsc"`
	YoutubeDsc   string `json:"youtube_dsc,omitempty" yaml:"youtube_dsc" mapstructure:"youtube_dsc"`
	GErrorSkip   bool   `json:"g_error_skip,omitempty" yaml:"g_error_skip" mapstructure:"g_error_skip"`
	YErrorSkip   bool   `json:"y_error_skip,omitempty" yaml:"y_error_skip" mapstructure:"y_error_skip"`
	BErrorSkip   bool   `json:"b_error_skip,omitempty" yaml:"b_error_skip" mapstructure:"b_error_skip"`
	YtErrorSkip  bool   `json:"yt_error_skip,omitempty" yaml:"yt_error_skip" mapstructure:"yt_error_skip"`

	EndRemove   bool `json:"end_remove,omitempty" yaml:"end_remove" mapstructure:"end_remove"`
	DownMainPic bool `json:"down_main_pic,omitempty" yaml:"down_main_pic" mapstructure:"down_main_pic"`
	VersionOption
}

type VersionOption struct {
	UseG  bool `json:"use_g,omitempty" yaml:"use_g" mapstructure:"use_g"`
	UseY  bool `json:"use_y,omitempty" yaml:"use_y" mapstructure:"use_y"`
	UseB  bool `json:"use_b,omitempty" yaml:"use_b" mapstructure:"use_b"`
	UseYT bool `json:"use_yt,omitempty" yaml:"use_yt" mapstructure:"use_yt"`

	Option        int32     `mapstructure:"option" json:"option,omitempty" yaml:"option"`
	RandTemp      bool      `mapstructure:"rand_temp" json:"rand_temp" yaml:"rand_temp"`
	Category      bool      `mapstructure:"category" json:"category,omitempty" yaml:"category"`
	CategoryLink  int32     `mapstructure:"category_link" json:"category_link,omitempty" yaml:"category_link"`
	ProductLink   int32     `mapstructure:"product_link" json:"product_link,omitempty" yaml:"product_link"`
	List          string    `mapstructure:"list" json:"list,omitempty" yaml:"list"`
	Article       string    `mapstructure:"article" json:"article,omitempty" yaml:"article"`
	UseBigSitemap bool      `mapstructure:"use_big_sitemap" json:"use_big_sitemap,omitempty" yaml:"use_big_sitemap"`
	BigSitemap    Sitemap   `mapstructure:"big_sitemap" json:"big_sitemap,omitempty" yaml:"big_sitemap"`
	SubSitemap    Sitemap   `mapstructure:"sub_sitemap" json:"sub_sitemap,omitempty" yaml:"sub_sitemap"`
	GoogleImgs    GoogleImg `mapstructure:"google_imgs" json:"google_imgs,omitempty" yaml:"google_imgs"`
	Paging        bool      `mapstructure:"paging" json:"paging,omitempty" yaml:"paging"`
	Word          bool      `mapstructure:"word" json:"word,omitempty" yaml:"word"`
}

type GoogleImg struct {
	Size      int32 `mapstructure:"size" json:"size" yaml:"size"`
	Option    int32 `mapstructure:"option" json:"option" yaml:"option"`
	GroupSize int32 `mapstructure:"group_size" json:"group_size" yaml:"group_size"`
}

type Sitemap struct {
	Size   int32 `mapstructure:"size" json:"size" yaml:"size"`
	Option int32 `mapstructure:"option" json:"option" yaml:"option"`
}
