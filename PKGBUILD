pkgname='hayashi'
pkgdesc='a tiny distro-independent package manager written in Go.'
url="https://github.com/crispybaccoon/$pkgname"
pkgver='0.1'
pkgrel=1
license=( 'GPT', 'MIT' )
arch=('any')

makedepends=('go')
depends=('bash' 'git')

prepare() {
  mkdir -p $HOME/.hayashi/repo
  cd $HOME/.hayashi/repo
  git clone --filter=blob:none https://github.com/crispybaccoon/hayashi
}

build() {
  cd $HOME/.hayashi/repo/hayashi
  go build -o ./hayashi .
}

package() {
  cd $HOME/.hayashi/repo/hayashi
  ./hayashi config init
  ./hayashi task pack hayashi
}
