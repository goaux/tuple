package main

import (
	"bufio"
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"text/template"

	"github.com/divan/num2words"
	"golang.org/x/sync/errgroup"
)

//go:embed *.go.tmpl
var tmpls embed.FS

type generator struct {
	basedir string
	tmpl    *template.Template
}

type data struct {
	N int
}

func main() {
	var (
		basedir string
		from    int
		to      int
	)
	flag.StringVar(&basedir, "d", "..", "output directory")
	flag.IntVar(&from, "f", 2, "from")
	flag.IntVar(&to, "t", 3, "to")
	flag.Parse()

	g, err := newGenerator(basedir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init generator: %v\n", err)
		os.Exit(1)
	}

	sem := make(chan struct{}, max(2, runtime.NumCPU()))
	var eg errgroup.Group
	for i := max(2, from); i <= to; i++ {
		sem <- struct{}{}
		eg.Go(func() error {
			defer func() { <-sem }()
			return g.gen(i)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := g.runGoimports(basedir); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func newGenerator(basedir string) (*generator, error) {
	t := template.New("").Funcs(template.FuncMap{
		"Add":       func(i, j int) int { return i + j },
		"Sub":       func(i, j int) int { return i - j },
		"Mod":       func(i, j int) int { return i % j },
		"IntToType": IntToType,
		"IntToVar":  IntToVar,
		"IntToWord": num2words.Convert,
	})
	t, err := t.ParseFS(tmpls, "*.go.tmpl")
	if err != nil {
		return nil, err
	}
	return &generator{basedir: basedir, tmpl: t}, nil
}

func (g *generator) gen(i int) error {
	dir := filepath.Join(g.basedir, fmt.Sprintf("tuple%d", i))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir error [%d]: %w", i, err)
	}

	base := fmt.Sprintf("tuple%d", i)
	list := []struct {
		Name string
		Tmpl string
	}{
		{base + ".go", "tuple.go.tmpl"},
		{base + "_compare.go", "tuple_compare.go.tmpl"},
		{base + "_getter.go", "tuple_getter.go.tmpl"},
		{base + "_test.go", "tuple_test.go.tmpl"},
	}
	for _, v := range list {
		if err := g.writeTemplate(filepath.Join(dir, v.Name), i, v.Tmpl); err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) writeTemplate(name string, n int, templateName string) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()
	bw := bufio.NewWriter(f)
	if err = g.tmpl.ExecuteTemplate(bw, templateName, &data{N: n}); err != nil {
		return err
	}
	return bw.Flush()
}

func (g *generator) runGoimports(dir string) error {
	if v, err := exec.LookPath("goimports"); err != nil || v == "" {
		fmt.Fprintln(os.Stderr, "warning: goimports not found, skipping formatting")
		return nil
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	cmd := exec.Command("goimports", "-w", dir)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("goimports error: %v, output: %s", err, string(out))
	}
	return nil
}

const alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var alphaLower = strings.ToLower(alphaUpper)

func IntToType(n int) string { return intToAlpha(n, alphaUpper) }
func IntToVar(n int) string  { return intToAlpha(n, alphaLower) }

func intToAlpha(n int, alphabet string) string {
	if n < 0 {
		return ""
	}
	var res []byte
	for n >= 0 {
		res = append(res, alphabet[n%len(alphabet)])
		n = n/len(alphabet) - 1
		if n < 0 {
			break
		}
	}
	slices.Reverse(res)
	return string(res)
}
