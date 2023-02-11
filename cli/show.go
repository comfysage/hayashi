package cli

import (
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/crispybaccoon/hayashi/pkg"
)

func show(p pkg.Pkg) {
	printf(COLOR_CYAN + p.Collection + "/" + COLOR_MAGENTA + p.Name)

	if p.Desc != "" {
		printf(p.Desc)
	}
	if p.Url != "" {
		printf(COLOR_CYAN + "url  " + COLOR_RESET + p.Url)
	}
	if len(p.Install) > 0 {
		printf(COLOR_CYAN + "bash")
		for _, s := range p.Install {
			printf("  " + COLOR_GREEN + s)
		}
	}
}

func Show(name string) error {
	pkg, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	show(pkg)

	return nil
}

func ShowPkg(name string) error {
	var p pkg.Pkg
	var err error
	if cfg.local {
		p, err = pkg.GetPkgFromPath(name)
		if err != nil {
			return err
		}
	} else {
		p, err = pkg.GetPkg(name)
		if err != nil {
			return err
		}
	}
	str, err := p.String()
	if err != nil {
		return err
	}

	lexer := lexers.Get("yaml")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)
	style := styles.Monokai
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.TTY16
	if formatter == nil {
		formatter = formatters.Fallback
	}
	iterator, err := lexer.Tokenise(nil, str)
	err = formatter.Format(os.Stdout, style, iterator)

	return nil
}
