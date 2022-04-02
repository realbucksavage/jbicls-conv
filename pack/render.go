package pack

import (
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/realbucksavage/jbicls-conv/conv"
)

type Renderer struct {
	tmpl *template.Template
}

func NewRenderer(source string) (*Renderer, error) {
	if source == "" {
		return nil, errors.New("theme pack source is required")
	}

	stat, err := os.Stat(source)
	if err != nil {
		return nil, fmt.Errorf("cannot stat %q: %v", source, err)
	}

	if !stat.IsDir() {
		return nil, fmt.Errorf("%q is not a directory", source)
	}

	g := fmt.Sprintf("%s/*", source)
	tmpl, err := template.ParseGlob(g)

	if err != nil {
		return nil, fmt.Errorf("cannot parse templates in %q: %v", source, err)
	}

	return &Renderer{tmpl: tmpl}, nil
}

func (r *Renderer) Run(scheme *conv.Scheme, target string) error {
	if err := r.tmpl.ExecuteTemplate(os.Stdout, target, scheme); err != nil {
		return err
	}

	return nil
}
