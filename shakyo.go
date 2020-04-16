package main

// 写経 10章p36 ミドルウェアを作る
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MiddleWare interface {
	ServeNext(h http.Handler) http.Handler
}
type MiddleWareFunc func(h http.Handler) http.Handler

func (f MiddleWareFunc) ServeNext(h http.Handler) http.Handler {
	return f(h)
}
func With(h http.Handler, ms ...MiddleWare) http.Handler {
	for _, m := range ms {
		h = m.ServeNext(h)
	}
	return h
}

// 写経 10章p39 ハンドラのテスト
// テスト対象
func handler(w, http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, net/http!")
}

// テストコード
func TestSample(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK { t.Fatal("unexpected status code") }
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil { t.Fatal("unexpected error")}
	const expected = "Hello, net/http!"
	if s := string(b); s != expected {t.Fatalf("unexpected response: %s", s)}
}

// レスポンスを読み取る
func main() {
	resp, err := http.Get("http:example.com")
	if err != nil { /* エラー処理 */
	}
	defer resp.Body.Close()
	var p Person
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&p); err != nil {
		// エラー処理
	}
	fmt.Println(p)
}

// リクエストを指定する
req, err := http.NewRequest("GET", "http://example.con", nil)
req.Header.Add("If-None-Match", `W/wyxxy`)
resp, err := client.Do(req)

//  リクエストとコンテキスト
ctx := req.Context()
req = req.WithConntext(ctx)

// DB
// DB のオープン
db, err := sql.Open("sqlite", "database.db")

// SQLの実行
// INSERTやDELETEなど
func (db *DB) Exec(query string, args ...interface{}) (Result, error) 
// SELECTなどで複数レコードを取得する場合
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
// SELECTなどで1つ􏰀レコードを取得する場合
func (db *DB) QueryRow(query string, args ...interface{}) *Row

// テーブルの作成
const sql = `
CREATE TABLE IF NOT EXSISTS user (
	id INTENGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	age INTEGER NOT NULL
);
`
if _, err := db.Exec(spl); err != nil {}

// レコードの挿入
type User struct {
	ID int64
	Name string
	Age int64
}
users := []*User{{Name: "tenntenn", Age: 32}, {Name: "Gopher", Age: 10}}
for i := range users {
	const sql = "INSERT INTO user(name, age) values (?, ?)"
	r, err := db.Exec(sql, users[i].Name, users[i].Age)
	if err != nil {}
	id, err := r.LastInsertId()
	if err != nil {}
	users[i].ID = id
	fmt.Println("INSERT", users[i])
}

// 複数レコードのスキャン
rows, err := db.Query("SELECT * FROM user WHERE age = ?", age)
if rows.Next() {
	var u User
	if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {}
	fmt.Println(u)
}
if err := rows.Err(); err != nil {}

// レコードの更新
r, err := db.Exec("UPDATE user SET age = age + 1 WHERE id = 1")
cnt , err := r.RowsAffected()

// トランザクションの開始
func (db *DB) Begin() (*Tx, error)

// contenxtに渡す場合
func (db *DB) BeginTx(context.Context, *TxOptions) (*TX, error)

// INSERT/DELETE
func (tx *Tx) Exec(query string, args ...interfase{}) (Result, error)
// SELECTで複数レコード取得
func (tx *Tx) Query(query string, args ...interfase{}) (*Rows, error)
// コミット
func (tx *Tx) Commit() error
// ロールバック
func (tx *Tx) Rollback() error