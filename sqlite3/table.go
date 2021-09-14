package sqlite3

// 会員情報
type Account struct {
	UserToken    string // ユーザートークン
	ID           string // ID
	Name         string // ログインID（e-mail）
	DesktopUser  string // サインインしているデスクトップユーザー名
	AccessToken  string // Auth0のアクセストークン
	RefreshToken string // Auth0のリフレッシュトークン
	Expiry       int64  // Auth0のアクセストークンの有効期限
	SessionId    string // 現在使用中のセッションＩＤ
	DeviceName   string // デバイス名
}

func (Account) TableName() string {
	return "account"
}

// 起動情報
type Config struct {
	TimeOnServer        int64  `gorm:"not null"` // セッションを取得したときの基準時刻
	TimeCounter         int64  `gorm:"not null"` // セッションを取得した後の経過時間
	SessionPeriod       int64  `gorm:"not null"` // セッション取得周期のカウント数
	OfflineExpiredCount int64  `gorm:"not null"` // オフラインを許容できるカウント数
	OfflineTimeCounter  int64  `gorm:"not null"` // オフライン経過時間に対応したカウント数
	TimeAtShutdown      int64  `gorm:"not null"` // シャットダウンの時の時間
	LastCheckUpdateTime string // 前回アップデート情報を確認した時間【Serviceは未使用】
	ProxyAddress        string // プロキシのホストアドレス
	ProxyPort           int32  // プロキシのポート番号
	ProxyUser           string // プロキシ認証ユーザー名
	ProxyPassword       string // プロキシ認証パスワード
	CountryCode         string // Agentから渡された国コード
}

func (Config) TableName() string {
	return "config"
}

// フォント情報
// SDD テーブル定義　6.1.1節 フォント情報｜Serviceに対応した構造体
type FontInformation struct {
	ID               string `gorm:"primary_key"` // 購入フォントＩＤ
	Name             string // フォント名 （用例）フォント一覧表の一段目に表示するフォント名
	EnglishName      string
	PostScriptName   string
	FileName         string // ファイル名
	FileSize         int
	Version          string // バージョン
	URL              string // フォントデータの保存場所（ダウンロード URL) （用例）フォントデータをフォント配布サーバーからダウンロードする
	Digest           string // フォントファイルのメッセージダイジェスト (MD5) （用例）正しくダウンロード出来たか確認するためにフォントデータと比較する
	OrderUpdateTime  int64  // 購入フォント情報変更日時 (UNIX時間)
	ProductCode      string // 製品のコード
	ServiceCode      string // サービスのコード
	OrderId          string // 注文ID （用例）MorisawaCore Libraryによりフォントファイルに購買情報を記録する
	DescriptionAsset string // フォントの説明
	AttrFontCode     string // フォントのコード （用例）MorisawaCore Libraryによりフォントファイルに購買情報を記録する
	ExpiryTime       int64  // 当該製品のライセンス期限切れ時間(UNIX時間) （用例）フォントの利用期限を計算する　2145884400を超える値は無期限とする
	AccountId        string // フォントを購入した会員アカウント （用例）MorisawaCore Libraryによりフォントファイルに購買情報を記録する
	CurrentVersion   string // 現在インストールしているフォントのバージョン （用例）サーバーのフォントとローカルのフォントをリビジョンチェックする。
	Status           int    // フォントの同期状態 （用例）フォント一覧表にインストール状態を表示する
	// 0 : 状態不明
	// 1: 未同期／未インストール　Not sync
	// 2: 使用可能	Download and sync font success
	// 3: 使用不可
	// 4: ダウンロード中
	// 5: パッケージフォント有り／同名フォント有り　FontList exist on OS
	FilePath string // ディスクに保存したフォントファイルの絶対パス
}

func (FontInformation) TableName() string {
	return "fonts"
}

// Collection はコレクション情報です
type Collection struct {
	ID   string // ID
	Name string // 名称
}

func (Collection) TableName() string {
	return "collections"
}
