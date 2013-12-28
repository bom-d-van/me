package app

import (
	"bytes"
	"github.com/bom-d-van/me/configs"
	"github.com/codegangsta/martini"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"
)

var (
	pkgPath      = os.Getenv("GOPATH") + "/src/github.com/bom-d-van/me"
	thoughtsPath = pkgPath + "/thoughts"
	log          = configs.Log
	apptmpl      *template.Template
)

func init() {
	mustParseViews()
}

func GetArticle(params martini.Params) string {
	name := params["artile_name"]

	log.Println("Get Article", name)

	artilePath := thoughtsPath + "/" + name + ".md"

	info, err := os.Stat(artilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
			return ""
		}

		artilePath = thoughtsPath + "/" + name + "/index.md"
		info, err = os.Stat(artilePath)
		if err != nil {
			log.Fatal(err)
			return ""
		}
	}

	content, err := ioutil.ReadFile(artilePath)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	html := template.HTML(blackfriday.MarkdownCommon(content))

	thoughtData := struct {
		Content template.HTML
		Title   string
		Cre     time.Time
		Mod     time.Time
	}{
		Content: html,
		Title:   getTitle(name),
	}

	sys := info.Sys()
	if sys != nil {
		stat := sys.(*syscall.Stat_t)
		thoughtData.Cre = time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec)
		thoughtData.Mod = info.ModTime()
	} else {
		thoughtData.Cre = info.ModTime()
		// thoughtData.Mod = nil
	}

	return genPage("thought", thoughtData)
}

func getTitle(name string) string {
	return strings.Title(strings.Join(strings.Split(name, "_"), " "))
}

type thought struct {
	Path  string
	Name  string
	Title string
	Mod   time.Time
	Cre   time.Time
}

type thoughtSorter struct {
	thoughts []thought
}

func (t *thoughtSorter) Len() int {
	return len(t.thoughts)
}

func (t *thoughtSorter) Less(i, j int) bool {
	return t.thoughts[i].Cre.UnixNano() > t.thoughts[j].Cre.UnixNano()
}

func (t *thoughtSorter) Swap(i, j int) {
	tmp := t.thoughts[i]
	t.thoughts[i] = t.thoughts[j]
	t.thoughts[j] = tmp
}

type IntsReverse struct {
	ints []int
}

func (t *IntsReverse) Len() int {
	return len(t.ints)
}

func (t *IntsReverse) Less(i, j int) bool {
	return t.ints[i] > t.ints[j]
}

func (t *IntsReverse) Swap(i, j int) {
	tmp := t.ints[i]
	t.ints[i] = t.ints[j]
	t.ints[j] = tmp
}

func GetThoughts() string {
	thoughts := []thought{}
	filepath.Walk(thoughtsPath, func(path string, info os.FileInfo, err error) error {
		if thoughtsPath == path {
			return err
		}
		if strings.Contains(strings.Replace(path, thoughtsPath+"/", "", -1), "/") {
			return err
		}

		name := info.Name()
		if strings.HasPrefix(name, ".") {
			return err
		}

		t := thought{}
		if info.IsDir() {
			info, err = os.Stat(path + "/index.md")
			if err != nil {
				return err
			}
			t.Name = name
		} else {
			t.Name = strings.Replace(name, ".md", "", -1)
		}

		t.Title = getTitle(t.Name)
		t.Mod = info.ModTime()

		sys := info.Sys()
		if sys != nil {
			stat := sys.(*syscall.Stat_t)
			t.Cre = time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec)
		} else {
			t.Cre = t.Mod
		}

		thoughts = append(thoughts, t)

		return err
	})

	sort.Sort(&thoughtSorter{thoughts})

	thoughtMap := map[int][]thought{}
	years := []int{}
	for i, t := range thoughts {
		year := t.Cre.Year()
		ts, ok := thoughtMap[year]
		if !ok {
			years = append(years, year)
		}

		ts = append(ts, thoughts[i])
		thoughtMap[year] = ts
	}

	sort.Sort(&IntsReverse{years})

	data := struct {
		Years      []int
		ThoughtMap map[int][]thought
	}{years, thoughtMap}

	return genPage("thoughts", data)
}

func GetAbout() string {
	return genPage("about", nil)
}

func genPage(tmpl string, data interface{}) string {
	if configs.ReLoadTemplate {
		mustParseViews()
	}

	buf := bytes.NewBufferString("")
	err := apptmpl.ExecuteTemplate(buf, tmpl, data)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return buf.String()
}

func mustParseViews() {
	var err error
	apptmpl = template.New("MeTmpl")
	apptmpl.Funcs(map[string]interface{}{
		"fmtOn": func(tim time.Time) string {
			return tim.Format("Mon Jan 2")
		},
		"fmtAt": func(tim time.Time) string {
			return tim.Format("Mon Jan 2 15:04")
		},
	})
	apptmpl, err = apptmpl.ParseGlob(pkgPath + "/app/views/*.html")
	if err != nil {
		panic(err)
	}
}
