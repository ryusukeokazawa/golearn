package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var n = flag.Bool("n", false, "通し番号を付与する")
	// option を設定
	flag.Parse()
	// file pathをターミナルから取得
	// TODO: filepathをハイフンでちゃんと区切らないと取得できない
	var files = flag.Args()
	// TODO: 絶対パスじゃないとひらけない
	var path, err = os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
	}

	//実行ファイルのディレクトリ名を取得
	path = filepath.Dir(path)
	//通し番号の用意
	i := 1

	for x := 0; x < len(files); x++ {
		sf, err := os.Open(filepath.Base(files[x]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		} else {
			scanner := bufio.NewScanner(sf)
			for ; scanner.Scan(); i++ {
				if *n {
					//オプションがある場合
					fmt.Printf("%v: ", i)
				}
				fmt.Println(scanner.Text())
			}
		}
	}
}
