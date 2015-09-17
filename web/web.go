package web

import (
	"net/http"
	"os"
	"path/filepath"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/ncw/rclone/fs"
	"github.com/skratchdot/open-golang/open"
	//"strings"
)

var homePath string

func init() {
	homePath, _ = os.Getwd() // FIXME: Should we embed static content in executable?
}

func StartServer() {

	r := mux.NewRouter()
	staticPath := filepath.Join(homePath, "web", "static")
	r.Handle("/static", http.FileServer(http.Dir(staticPath)))
	r.HandleFunc("/", MainPage)
	//r.HandleFunc("/articles", ArticlesHandler)

	open.Start("http://127.0.01:5678/")
	http.ListenAndServe("127.0.0.1:5678", r)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
	s := fmt.Sprintf("Config:%#v\n", fs.Config)
	remotes := fs.ConfigFile.GetSectionList()
	s += "Remotes:\n"
	//s += strings.Join(remotes, "\n")
	for _, remote := range remotes {
		s += fmt.Sprintf("%-20s %s\n", remote, fs.ConfigFile.MustValue(remote, "type"))
	}

	w.Write([]byte(s))
}
