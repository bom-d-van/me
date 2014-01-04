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
	"time"
)

var (
	pkgPath      = os.Getenv("GOPATH") + "/src/github.com/bom-d-van/me"
	thoughtsPath = pkgPath + "/thoughts"
	log          = configs.Log
	apptmpl      *template.Template
	creationInfo = map[string]time.Time{}
	timeNow      = time.Now()
)

func init() {
	mustParseViews()
	creationInfoBytes, err := ioutil.ReadFile(pkgPath + "/thoughts_creation_file.txt")
	if err != nil {
		log.Panic(err)
	}

	for _, item := range strings.Split(string(creationInfoBytes), "\n") {
		if item == "" {
			continue
		}
		segs := strings.Split(item, "://")
		if len(segs) != 2 {
			log.Panic(item + " Is Illegal Format.")
		}
		tim, err := time.Parse("01/02/2006 15:04:05", segs[1])
		if err != nil {
			log.Panic(err)
		}
		creationInfo[segs[0]] = tim
	}
}

func genCreationTime(path string) time.Time {
	path = strings.Replace(path, thoughtsPath, "", 1)
	tim, ok := creationInfo[path]
	if !ok {
		tim = timeNow
	}
	log.Println(tim)

	return tim
}

func GetArticle(params martini.Params) string {
	name := params["artile_name"]

	log.Println("Get Article", name)

	artilePath := thoughtsPath + "/" + name + ".md"

	info, err := os.Stat(artilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
			return ""
		}

		artilePath = thoughtsPath + "/" + name + "/index.md"
		info, err = os.Stat(artilePath)
		if err != nil {
			log.Println(err)
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
		NoMod   bool
	}{
		Content: html,
		Title:   getTitle(name),
		Mod:     info.ModTime(),
	}

	log.Println(artilePath)
	thoughtData.Cre = genCreationTime(artilePath)
	thoughtData.NoMod = thoughtData.Cre.Format(atFmt) != thoughtData.Mod.Format(atFmt)

	return genPage("thought", thoughtData)
}

func getTitle(name string) string {
	return strings.Title(strings.Join(strings.Split(name, "_"), " "))
}

type thought struct {
	Name  string
	Title string
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
			path = path + "/index.md"
			info, err = os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}

				if log != nil {
					log.Println(err)
				}
				return err
			}
			t.Name = name
		} else {
			t.Name = strings.Replace(name, ".md", "", -1)
		}

		t.Title = getTitle(t.Name)
		// t.Mod = info.ModTime()
		t.Cre = genCreationTime(path)

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

var atFmt = "Mon Jan 2 15:04"

func mustParseViews() {
	var err error
	apptmpl = template.New("MeTmpl")
	apptmpl.Funcs(map[string]interface{}{
		"fmtOn": func(tim time.Time) string {
			return tim.Format("Mon Jan 2")
		},
		"fmtAt": func(tim time.Time) string {
			return tim.Format(atFmt)
		},
	})
	apptmpl, err = apptmpl.ParseGlob(pkgPath + "/app/views/*.html")
	if err != nil {
		panic(err)
	}
}
