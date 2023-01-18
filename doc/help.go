package doc

var HELP_DOC = Doc{
	Name: "help",
	Man: ManDoc{
		Short: "hayashi. a tiny distro-independent package manager written in Go.",
		Long: "",
		Usage: "",
		Cmddoc: docs{
			&CONFIG_DOC,
			&PKG_DOC,
		},
	},
}
var CONFIG_DOC = Doc{
	Name: "config",
	Man: ManDoc{
		Short: "Manage global configuration",
		Long: "",
		Usage: "ysi config <command>",
		Cmddoc: docs{
			&CONFIG_INIT_DOC,
			&CONFIG_CREATE_DOC,
		},
	},
}
var (
	CONFIG_INIT_DOC = Doc{
		Name: "config-init",
		Man: ManDoc{
			Name: "init",
			Short: "Setup hayashi",
			Long:  "Create configuration structure and setup core collection.\nFetch core collection pkg and install.",
			Usage: "ysi config init",
		},
	}
	CONFIG_CREATE_DOC = Doc{
		Name: "config-create",
		Man: ManDoc{
			Name: "create",
			Short: "Setup .hayashi.yaml",
			Long:  "Create default configuration file",
			Usage: "ysi config create",
		},
	}
)
var PKG_DOC = Doc{
	Name: "pkg",
	Man: ManDoc{
		Short: "Manage pkg configurations",
		Long:  "",
		Usage: "ysi pkg <command> <pkg_name>",
		Cmddoc: docs{
			&PKG_SHOW_DOC,
			&PKG_ADD_DOC,
		},
	},
}
var (
	PKG_SHOW_DOC = Doc{
		Name: "pkg-show",
		Man: ManDoc{
			Name: "show",
			Short: "Show pkg configuration",
			Long:  "Show yaml contents of a pkg",
			Usage: "ysi pkg show <pkg_name>",
		},
	}
	PKG_ADD_DOC = Doc{
		Name: "pkg-add",
		Man: ManDoc{
			Name: "add",
			Short: "Add pkg configuration",
			Long:  "Add a pkg from the commandline or add a local configuration to global configurations",
			Usage: "ysi pkg add <pkg_name> [<git_url>]",
			Flagdoc: []Shortdoc{
				{"--local", "add a local configuration to a global collection",
					"ysi pkg add --local <pkg_path>"},
			},
		},
	}
)

var DOCS = DefineDocs(docs{
	&HELP_DOC,
	&CONFIG_DOC,
	&CONFIG_INIT_DOC,
	&CONFIG_CREATE_DOC,
	&PKG_DOC,
	&PKG_SHOW_DOC,
	&PKG_ADD_DOC,
})
