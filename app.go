/*

ターミナル移動、ファイルオープン操作等を、楽しみながら学ぶためのプログラムです。


go build -v -o takarasagashi app.go

*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"time"
)

const (
	ATARI_MESSAGE  = "あたり"
	HAZURE_MESSAGE = "はずれ"
	SUBDIR_RATIO   = 40
)

/*
作成する、ディレクトリ名を設定できます。
*/
var DIR_NAMES []string = []string{"isu", "tukue", "hako", "kokuban", "kuruma", "densya", "sora"}

/*
作成する、ファイル名を設定できます。
*/
var FILE_NAMES []string = []string{"kuji", "takarabako"}

var setting struct {
	DirCount  int
	FileCount int
	RootName  string
}

// この init() 関数は　main() より先に呼ばれます。
func init() {
	flag.IntVar(&setting.DirCount, "dir_count", 10, "ディレクトリ数")
	flag.IntVar(&setting.FileCount, "file_count", 20, "ファイル数")
	flag.StringVar(&setting.RootName, "root_name", "takara", "宝物のディレクトリ名")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	root, _ := os.Getwd()
	fmt.Printf("いまの場所は\n\n「%s」\n\nだよ。\nここに宝物ディレクトリ\n\n「%s」\n\nをつくるよ。作成開始！\n", root, setting.RootName)

	root = path.Join(root, setting.RootName)

	_, e := os.Stat(root)

	if e == nil {
		fmt.Printf("\n\n\nあれ？\n\n\nすでに宝物ディレクトリ「%s」があるよ。削除するか、違う名前で作ろう\n\n", setting.RootName)
		return
	}

	dirs := genDirList(setting.DirCount)
	createDir(root, dirs)
	hideTakara(root, dirs, setting.FileCount)

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("隠し中....")

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("隠し中....")

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("隠し中....")

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("\n宝を隠したよ！「%s」とかかれたファイルを探してね！\n", ATARI_MESSAGE)

}

// 宝をかくす（ファイルを作成する）
func hideTakara(root string, dirs []string, count int) {

	atari := rand.Intn(count)

	for i := 0; i < count; i++ {
		name := FILE_NAMES[rand.Intn(len(FILE_NAMES))]
		dir := dirs[rand.Intn(len(dirs))]

		file := path.Join(root, dir, fmt.Sprintf("%s.txt", name))
		_, e := os.Stat(file)

		if e == nil {
			file = path.Join(root, dir, fmt.Sprintf("%s%d.txt", name, i))
		}

		if i == atari {
			ioutil.WriteFile(file, []byte(ATARI_MESSAGE), os.ModePerm)
		} else {
			ioutil.WriteFile(file, []byte(HAZURE_MESSAGE), os.ModePerm)
		}

	}
}

// ディレクトリを作成する
func createDir(root string, dirs []string) {
	for _, dir := range dirs {
		dir = path.Join(root, dir)
		os.MkdirAll(dir, 0777)
	}
}

// 作成するディレクトリの一覧を取得する
func genDirList(max int) []string {

	tree := map[string]bool{}
	list := []string{}

	for i := 0; i < max; i++ {
		name := DIR_NAMES[rand.Intn(len(DIR_NAMES))]

		// subdir
		if rand.Intn(100) < SUBDIR_RATIO && len(list) > 0 {
			parent := list[rand.Intn(len(list))]
			name = parent + "/" + name
		}

		if _, has := tree[name]; has {
			name = fmt.Sprintf("%s%d", name, i)
		}
		tree[name] = true
		list = append(list, name)
	}

	return list
}
