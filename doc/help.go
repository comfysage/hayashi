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
			&TASK_DOC,
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
var TASK_DOC = Doc{
	Name: "task",
	Man: ManDoc{
		Short: "Run pkg tasks",
		Long:  "",
		Usage: "ysi task <command> <pkg_name>",
		Cmddoc: docs{
			&TASK_CLONE_DOC,
			&TASK_BUILD_DOC,
			&TASK_PACK_DOC,
		},
	},
}
var (
	TASK_CLONE_DOC = Doc{
		Name: "task-clone",
		Man: ManDoc{
			Name: "clone",
			Short: "Clone pkg repo",
			Long:  "Clone the package repository",
			Usage: "ysi task clone <pkg_name>",
		},
	}
	TASK_BUILD_DOC = Doc{
		Name: "task-build",
		Man: ManDoc{
			Name: "build",
			Short: "Build pkg",
			Long:  "Build a cloned package from source",
			Usage: "ysi task build <pkg_name>",
		},
	}
	TASK_PACK_DOC = Doc{
		Name: "task-pack",
		Man: ManDoc{
			Name: "pack",
			Short: "Pack pkg",
			Long:  "Install installation files for a package",
			Usage: "ysi task pack <pkg_name>",
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
	&TASK_DOC,
	&TASK_CLONE_DOC,
	&TASK_BUILD_DOC,
	&TASK_PACK_DOC,
})
