// SQLite3／暗号化したデータベースファイル利用する
package sqlite3

import (
	"os"
	"strings"
	"time"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"

	// SQLCipher is an SQLite extension
	_ "github.com/xeodou/go-sqlcipher"
)

const cipherCommand = "PRAGMA key = 'QJZA9SK5VDTAXY3V5B3JMF6YB8EI02P9';"

const (
	NotFoundFileMessage = "no such file or directory"
)

// Abstract : DBオブジェクトを返す。DBが初期化されていない場合は、DBの初期化を行う
//
// Reentrant: 可能
//
// Argument:
//  Type             Name                    In/Out  Comment
//  string           dbFilePath              (I)     DBファイルのパス
//
// Return:
//  Type             Comment
//  *gorm.DB         DB操作オブジェクト
//  error            エラー
func GetDatabase(dbFilePath string) (db *gorm.DB, err error) {
	if checkFileSize(dbFilePath) < 1000 {
		// ファイルサイズが一定サイズ以下なら不正なファイルなので削除する
		err = os.Remove(dbFilePath)
		if err != nil {
			strErr := err.Error()
			if -1 == strings.Index(strErr, NotFoundFileMessage) {
				_ = log.Error("[GetDatabase]: os.Remove(), ", err)
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	// データベースのコネクションを開く（ファイルがなければ新規作成される）
	db, err = gorm.Open("sqlite3", dbFilePath+"?_key=QJZA9SK5VDTAXY3V5B3JMF6YB8EI02P9")
	if err != nil {
		_ = log.Critical("[GetDatabase]: gorm.Open(), ", err)
		return
	}

	// データベースの暗号・復号キーの設定
	queryResult := db.Exec(cipherCommand)
	if queryResult.Error != nil {
		err = queryResult.Error
		_ = log.Critical("[GetDatabase]: db.Exec(), ", err)
		_ = db.Close()
		return nil, err
	}

	// DBファイルの異常確認のため、テーブル一覧をselect
	queryResult = db.Exec("select name from sqlite_master where type='table';")
	if queryResult.Error != nil {
		err = queryResult.Error
		_ = log.Critical("[GetDatabase]: db.Exec() 2, ", err)
		_ = db.Close()
		return nil, queryResult.Error
	}

	db.AutoMigrate(&Config{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&FontInformation{})
	db.AutoMigrate(&Collection{})

	return
}

func checkFileSize(filePath string) int64 {
	// ファイルが存在する場合は追記、存在しない場合は新規作成
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)

	if err != nil {
		strErr := err.Error()
		if -1 == strings.Index(strErr, NotFoundFileMessage) {
			_ = log.Error("[checkFileSize]: os.OpenFile(), ", err)
		} else {
			log.Info("[checkFileSize] os.OpenFile is create new file.")
		}
		return 0
	}

	// クローズ処理の遅延実行
	defer func() {
		_ = file.Close()
	}()

	// ファイル情報取得
	fileinfo, staterr := file.Stat()
	if staterr != nil {
		_ = log.Error("[checkFileSize]: file.Stat(), ", staterr)
		return 0
	}

	// ファイルサイズ
	return fileinfo.Size()
}

// Abstract : 起動情報を読み込む
//
// Reentrant: 可能
//
// Argument:
//  Type             Name                    In/Out  Comment
//  *gorm.DB         db                      (I)     DB操作オブジェクト
//
// Return:
//  Type             Comment
//  []Config         読み込んだレコード
//
// 制限事項:
//   DBオブジェクトが生成されていること
func GetConfig(db *gorm.DB) []Config {
	data := []Config{}
	db.Find(&data)

	return data
}

// Abstract : 会員情報を読み込む
//
// Reentrant: 可能
//
// Argument:
//  Type             Name                    In/Out  Comment
//  *gorm.DB         db                      (I)     DB操作オブジェクト
//
// Return:
//  Type             Comment
//  []Account        読み込んだレコード
//
// 制限事項:
//   DBオブジェクトが生成されていること
func GetAccount(db *gorm.DB) []Account {
	data := []Account{}
	db.Find(&data)

	return data
}

// Abstract : フォント情報を読み込む
//
// Reentrant: 可能
//
// Argument:
//  Type             Name                    In/Out  Comment
//  *gorm.DB         db                      (I)     DB操作オブジェクト
//
// Return:
//  Type              Comment
//  []FontInformation 読み込んだレコード
//
// 制限事項:
//   DBオブジェクトが生成されていること
func GetFonts(db *gorm.DB) []FontInformation {
	data := []FontInformation{}
	db.Find(&data)

	return data
}

// GetCollections はコレクションを返却します
func GetCollections(db *gorm.DB) []Collection {
	data := []Collection{}
	db.Find(&data)

	return data
}

// Abstract : 起動情報を書き込む
//
// Reentrant: 可能
//
// Argument:
//  Type             Name                    In/Out  Comment
//  *gorm.DB         db                      (I)     DB操作オブジェクト
//  []Config         records                 (I)     書き込み後のレコード
//
// Return:
//  Type             Comment
//  error            エラー
//
// 制限事項:
//   DBオブジェクトが生成されていること
func WriteConfig(db *gorm.DB, records []Config) (err error) {
	err = db.Delete(Config{}).Error

	for _, record := range records {
		if err = db.Create(record).Error; err != nil {
			return
		}
	}

	return
}

// Abstract : 会員情報を書き込む
//
// Reentrant: 可能
//
// Argument:
//  Type             Name                    In/Out  Comment
//  *gorm.DB         db                      (I)     DB操作オブジェクト
//  []Account        records                 (I)     書き込み後のレコード
//
// Return:
//  Type             Comment
//  error            エラー
//
// 制限事項:
//   DBオブジェクトが生成されていること
func WriteAccount(db *gorm.DB, records []Account) (err error) {
	err = db.Delete(Account{}).Error

	for _, record := range records {
		if err = db.Create(record).Error; err != nil {
			return
		}
	}

	return
}

// Abstract : フォント情報を書き込む
//
// Reentrant: 可能
//
// Argument:
//  Type              Name                    In/Out  Comment
//  *gorm.DB          db                      (I)     DB操作オブジェクト
//  []FontInformation records                 (I)     書き込み後のレコード
//
// Return:
//  Type             Comment
//  error            エラー
//
// 制限事項:
//   DBオブジェクトが生成されていること
func WriteFonts(db *gorm.DB, records []FontInformation) (err error) {
	err = db.Delete(FontInformation{}).Error
	//_ = log.Error("[WriteFonts]", err)

	for _, record := range records {
		if err = db.Create(record).Error; err != nil {
			//_ = log.Error("[WriteFonts]", err)
			return
		}
	}

	return
}

// WriteCollections はコレクション情報を書き出します
func WriteCollections(db *gorm.DB, records []Collection) (err error) {
	err = db.Delete(Collection{}).Error

	for _, record := range records {
		if err = db.Create(record).Error; err != nil {
			return
		}
	}

	return
}
