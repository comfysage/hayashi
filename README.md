# hayashi

:seedling: a tiny distro-independent package manager written in Go.

## Usage

### Pkg Config

A pkg is defined in a yaml file at one of the [collections](#collections) in `~/.hayashi/pkg/`.

```yaml
pkg: pkg_name
collection: pkg_collection
desc: description
url: git_url

install:
 - install_script
update:
 - update_script
remove:
 - remove_script
```

### Installing a pkg

Install pkg defined in `~/.hayashi/pkg/`:

    $ hayashi add pkg_name

> This will search for a [pkg config](#pkg-config), clone the git repo and run the install script.

### Adding a pkg config

Add a pkg config at `~/.hayashi/pkg/custom/pkg_name.yaml`:

    $ hayashi pkg add pkg_name

You can also specify a git url for the pkg:

    $ hayashi pkg add pkg_name git_url

### Removing a pkg config

    $ hayashi pkg remove pkg_name

### Showing info about a pkg

	$ hayashi show pkg_name

### Updating a pkg

    $ hayashi update pkg_name

## Installation

hayashi can be installed with go:

```bash
go install github.com/CrispyBaccoon/hayashi
```

or you can [build from source](#building-from-source).

### Building from source

	$ cd ~/
    $ git clone --depth 1 https://github.com/CrispyBaccoon/hayashi.git
    $ cd ~/hayashi
    $ make && sudo make install

## Configured packages

### Collections

Packages are stored in collections. hayashi includes some base packages in the core collection.

### Core packages

There are multiple packages included in the repo.
- [cbonsai](https://gitlab.com/jallbrit/cbonsai). a bonsai tree generator written in C.
- [neovim](https://github.com/neovim/neovim). Vim-fork focused on extensibility and usability.
- [fzf](https://github.com/junegunn/fzf). a general-purpose command-line fuzzy finder.
- [bat](https://github.com/sharkdp/bat). A cat(1) clone with wings.
- [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh). A delightful framework for managing your zsh configuration.
- [vim-plug](https://github.com/junegunn/vim-plug). A minimalist Vim plugin manager.

