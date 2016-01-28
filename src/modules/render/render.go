package render

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/contrib/renders/multitemplate"

	"conf"
	"templates"

	"modules/log"
)

func LoadTemplates() multitemplate.Render {
	switch conf.TMPL_TYPE {
	case conf.BINDATA:
		return loadTemplatesBindata(conf.TMPL_DIR)
	default:
		return loadTemplatesDefault(conf.TMPL_DIR)
	}
}

func loadTemplatesDefault(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templatesDir + "/layouts/"
	layouts, err := filepath.Glob(layoutDir + "*/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includeDir := templatesDir + "/includes/"
	includes, err := filepath.Glob(includeDir + "*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(template.ParseFiles(files...))
		tmplName := strings.TrimPrefix(layout, layoutDir)
		tmplName = strings.TrimSuffix(tmplName, ".tmpl")
		log.DebugPrint("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return r
}

func loadTemplatesBindata(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templatesDir + "/layouts"
	layoutDirs, err := templates.AssetDir(layoutDir)
	if err != nil {
		panic(err.Error())
	}

	var layouts []string
	for _, dir := range layoutDirs {
		files, err := templates.AssetDir(layoutDir + "/" + dir)
		if err != nil {
			panic(err.Error())
		}

		// 过滤非.tmpl后缀模板
		layoutFiels, err := tmplsFilter(files, layoutDir+"/"+dir)
		if err != nil {
			panic(err.Error())
		}

		layouts = append(layouts, layoutFiels...)
	}

	includeDir := templatesDir + "/includes"
	includeFiels, err := templates.AssetDir(includeDir)
	if err != nil {
		panic(err.Error())
	}
	// 过滤非.tmpl后缀模板
	includes, err := tmplsFilter(includeFiels, includeDir)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(parseBindataFiles(files...))
		tmplName := strings.TrimPrefix(layout, layoutDir+"/")
		tmplName = strings.TrimSuffix(tmplName, ".tmpl")
		log.DebugPrint("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return r
}

// 过滤非tmpl后缀模板文件
func tmplsFilter(files []string, dir string) ([]string, error) {
	var tmpls []string
	for _, file := range files {
		if strings.HasSuffix(file, ".tmpl") {
			tmpls = append(tmpls, dir+"/"+file)
		}
	}
	return tmpls, nil
}

// parseFiles is the helper for the method and function. If the argument
// template is nil, it is created from the first file.
func parseBindataFiles(filenames ...string) (*template.Template, error) {
	var t *template.Template
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
	}
	for _, filename := range filenames {
		b, err := templates.Asset(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		name := filepath.Base(filename)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
