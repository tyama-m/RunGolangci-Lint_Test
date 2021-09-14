package sqlite3

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
)

var dbFilePath = "./test.db"

// DBファイルの削除
func removeDbFile() {
	os.Remove(dbFilePath)

	if _, err := os.Stat(dbFilePath); err == nil {
		fmt.Println("unable to remove db file")
		return
	}
}

// 2つのアカウントが一致するかどうかを返す
func equalsAccount(object1 []Account, object2 []Account) bool {
	if len(object1) != len(object2) {
		return false
	}

	// ソート
	sort.SliceStable(object1, func(i, j int) bool {
		return object1[i].Name < object1[j].Name
	})
	sort.SliceStable(object2, func(i, j int) bool {
		return object2[i].Name < object2[j].Name
	})

	// 各要素が一致するかの確認
	for i := range object1 {
		if object1[i] != object2[i] {
			return false
		}
	}

	return true
}

// 2つの起動情報が一致するかどうかを返す
func equalsConfig(object1 []Config, object2 []Config) bool {
	if len(object1) != len(object2) {
		return false
	}

	// ソート
	sort.SliceStable(object1, func(i, j int) bool {
		return object1[i].ProxyPort < object1[j].ProxyPort
	})
	sort.SliceStable(object2, func(i, j int) bool {
		return object2[i].ProxyPort < object2[j].ProxyPort
	})

	// 各要素が一致するかの確認
	for i := range object1 {
		if object1[i] != object2[i] {
			return false
		}
	}

	return true
}

// 2つのフォント情報が一致するかどうかを返す
func equalsFont(font1 []FontInformation, font2 []FontInformation) bool {
	if len(font1) != len(font2) {
		return false
	}

	// ソート
	sort.SliceStable(font1, func(i, j int) bool {
		return font1[i].AttrFontCode < font1[j].AttrFontCode
	})
	sort.SliceStable(font2, func(i, j int) bool {
		return font2[i].AttrFontCode < font2[j].AttrFontCode
	})

	// 各要素が一致するかの確認
	for i := range font1 {
		if font1[i] != font2[i] {
			return false
		}
	}

	return true
}

// DBファイルがない状態でDB操作オブジェクト取得
func TestGetDatabase(t *testing.T) {
	// テスト準備
	expectedError := error(nil)
	removeDbFile()

	// テスト実施
	db, err := GetDatabase(dbFilePath)
	if db != nil {
		defer db.Close()
	}

	// 検証
	if db == nil {
		t.Errorf("db is null")
	}

	if err != expectedError {
		t.Errorf("err is not %v but %v", expectedError, err)
	}

	_, err = os.Stat(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
	}
}

// DBファイルがある状態でDB操作オブジェクト取得
func TestGetDatabase2(t *testing.T) {
	// テスト準備
	expectedError := error(nil)
	removeDbFile()

	db, err := GetDatabase(dbFilePath)
	if db != nil {
		db.Close()
	}
	_, err = os.Stat(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
	}

	// テスト実施
	db, err = GetDatabase(dbFilePath)
	if db != nil {
		defer db.Close()
	}

	// 検証
	if db == nil {
		t.Errorf("db is null")
	}

	if err != expectedError {
		t.Errorf("err is not %v but %v", expectedError, err)
	}

	_, err = os.Stat(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
	}
}

// 不正なDBファイルがある状態でDB操作オブジェクト取得
func TestGetDatabase3(t *testing.T) {
	// テスト準備
	expectedError := error(nil)
	removeDbFile()

	ioutil.WriteFile(dbFilePath, []byte(strings.Repeat("abc", 334)), 0644)
	_, err := os.Stat(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
	}

	// テスト実施
	db, err := GetDatabase(dbFilePath)

	// 検証
	if db != nil {
		t.Errorf("bad db file opened")
		defer db.Close()
	}

	if err == expectedError {
		t.Errorf("err is not %v but %v", expectedError, err)
	}

	_, err = os.Stat(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
	}

	t.Log("Ok. ignore the error massage -> (file is encrypted or is not a database) ")
}

// DBファイルからレコード読み込み（0件）
func TestGetConfig(t *testing.T) {
	// テスト準備
	removeDbFile()

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	config := GetConfig(db)

	// 検証
	if len(config) != 0 {
		t.Errorf("wrong record count")
	}
}

// DBファイルからレコード読み込み（3件）
func TestGetConfig2(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedConfig := []Config{
		{1, 2, 3, 4, 5, 6,
			"a", "b",
			7,
			"c", "d", "e"},
		{11, 12, 13, 14, 15, 16,
			"aa", "bb",
			17,
			"cc", "dd", "ee"},
		{21, 22, 23, 24, 25, 26,
			"aaa", "bbb",
			27,
			"ccc", "ddd", "eee"},
	}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	WriteConfig(db, expectedConfig)

	// テスト実施
	config := GetConfig(db)

	// 検証
	if !equalsConfig(config, expectedConfig) {
		t.Error("config error")
		t.Error("    result  :", config)
		t.Error("    expected:", expectedConfig)
	}
}

// DBファイルからレコード読み込み（0件）
func TestGetAccount(t *testing.T) {
	// テスト準備
	removeDbFile()

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	account := GetAccount(db)

	// 検証
	if len(account) != 0 {
		t.Errorf("wrong record count")
	}
}

// DBファイルからレコード読み込み（3件）
func TestGetAccount2(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedAccount := []Account{
		{"a", "b", "c", "o", "r", 1, "u"},
		{"d", "e", "f", "p", "s", 2, "v"},
		{"g", "h", "i", "q", "t", 3, "w"},
	}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	WriteAccount(db, expectedAccount)

	// テスト実施
	account := GetAccount(db)

	// 検証
	if !equalsAccount(account, expectedAccount) {
		t.Error("account error")
		t.Error("    result  :", account)
		t.Error("    expected:", expectedAccount)
	}
}

// DBファイルからレコード読み込み（0件）
func TestGetFonts(t *testing.T) {
	// テスト準備
	removeDbFile()

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	account := GetFonts(db)

	// 検証
	if len(account) != 0 {
		t.Errorf("wrong record count")
	}
}

// DBファイルからレコード読み込み（3件）
func TestGetFonts2(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedFont := []FontInformation{
		{"A", 1, "b", "c", "d", "e", "f",
			"g", "h", "i", 2, "k", "l", "m",
			"n", "o", 3, "q", "r", 4, "t"},
		{"B", 1, "b", "c", "d", "e", "f",
			"g", "h", "i", 2, "k", "l", "m",
			"n", "o", 3, "q", "r", 4, "t"},
		{"C", 1, "b", "c", "d", "e", "f",
			"g", "h", "i", 2, "k", "l", "m",
			"n", "o", 3, "q", "r", 4, "t"},
	}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	WriteFonts(db, expectedFont)

	// テスト実施
	font := GetFonts(db)

	// 検証
	if !equalsFont(font, expectedFont) {
		t.Error("font error")
		t.Error("    result  :", font)
		t.Error("    expected:", expectedFont)
	}
}

// DBファイルにレコード書き込み（0件）
func TestWriteConfig(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedConfig := []Config{}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	err = WriteConfig(db, expectedConfig)

	// 検証
	if err != nil {
		t.Errorf("unable to write record")
	}

	wroteRecord := GetConfig(db)
	if !equalsConfig(wroteRecord, expectedConfig) {
		t.Error("config error")
		t.Error("    result  :", wroteRecord)
		t.Error("    expected:", expectedConfig)
	}
}

// DBファイルにレコード書き込み（3件）
func TestWriteConfig2(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedConfig := []Config{
		{1, 2, 3, 4, 5, 6,
			"a", "b",
			7,
			"c", "d", "e"},
		{11, 12, 13, 14, 15, 16,
			"aa", "bb",
			17,
			"cc", "dd", "ee"},
		{21, 22, 23, 24, 25, 26,
			"aaa", "bbb",
			27,
			"ccc", "ddd", "eee"},
	}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	err = WriteConfig(db, expectedConfig)

	// 検証
	if err != nil {
		t.Errorf("unable to write record")
	}

	wroteRecord := GetConfig(db)
	if !equalsConfig(wroteRecord, expectedConfig) {
		t.Error("config error")
		t.Error("    result  :", wroteRecord)
		t.Error("    expected:", expectedConfig)
	}
}

// DBファイルにレコード書き込み（0件）
func TestWriteAccount(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedAccount := []Account{}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	err = WriteAccount(db, expectedAccount)

	// 検証
	if err != nil {
		t.Errorf("unable to write record")
	}

	wroteRecord := GetAccount(db)
	if !equalsAccount(wroteRecord, expectedAccount) {
		t.Error("account error")
		t.Error("    result  :", wroteRecord)
		t.Error("    expected:", expectedAccount)
	}
}

// DBファイルにレコード書き込み（3件）
func TestWriteAccount2(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedAccount := []Account{
		{"a", "b", "c", "o", "r", 1, "u"},
		{"d", "e", "f", "p", "s", 2, "v"},
		{"g", "h", "i", "q", "t", 3, "w"},
	}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	err = WriteAccount(db, expectedAccount)

	// 検証
	if err != nil {
		t.Errorf("unable to write record")
	}

	wroteRecord := GetAccount(db)
	if !equalsAccount(wroteRecord, expectedAccount) {
		t.Error("account error")
		t.Error("    result  :", wroteRecord)
		t.Error("    expected:", expectedAccount)
	}
}

// DBファイルにレコード書き込み（0件）
func TestWriteFonts(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedFont := []FontInformation{}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	err = WriteFonts(db, expectedFont)

	// 検証
	if err != nil {
		t.Errorf("unable to write record")
	}

	font := GetFonts(db)
	if !equalsFont(font, expectedFont) {
		t.Error("font error")
		t.Error("    result  :", font)
		t.Error("    expected:", expectedFont)
	}
}

// DBファイルにレコード書き込み（3件）
func TestWriteFonts2(t *testing.T) {
	// テスト準備
	removeDbFile()
	expectedFont := []FontInformation{
		{"A", 1, "b", "c", "d", "e", "f",
			"g", "h", "i", 2, "k", "l", "m",
			"n", "o", 3, "q", "r", 4, "t"},
		{"B", 1, "b", "c", "d", "e", "f",
			"g", "h", "i", 2, "k", "l", "m",
			"n", "o", 3, "q", "r", 4, "t"},
		{"C", 1, "b", "c", "d", "e", "f",
			"g", "h", "i", 2, "k", "l", "m",
			"n", "o", 3, "q", "r", 4, "t"},
	}

	db, err := GetDatabase(dbFilePath)
	if err != nil {
		t.Errorf("db file is not created")
		return
	}
	defer db.Close()

	// テスト実施
	err = WriteFonts(db, expectedFont)

	// 検証
	if err != nil {
		t.Errorf("unable to write record")
	}

	font := GetFonts(db)
	if !equalsFont(font, expectedFont) {
		t.Error("font error")
		t.Error("    result  :", font)
		t.Error("    expected:", expectedFont)
	}
}

func TestMain(m *testing.M) {
	// テスト準備

	// DBファイルの削除
	removeDbFile()

	ret := m.Run()

	// テスト後始末

	// DBファイルの削除
	removeDbFile()

	os.Exit(ret)
}
