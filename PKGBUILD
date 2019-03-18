# Maintainer: Dav Lord <vlord@free.fr>
pkgname=aurgasm
pkgver=0.1
pkgrel=1
pkgdesc="A simple ArchLinux AUR build tool"
arch=('x86_64')
url="https://github.com/davlord/aurgasm"
license=('GPL')
depends=('pacman','wget')
makedepends=('go')

prepare() {
	mkdir "$pkgname-$pkgver"
}

build() {
	cd "$pkgname-$pkgver"
	export GOPATH=$(pwd)
	go get github.com/davlord/aurgasm
	go install github.com/davlord/aurgasm
}

package() {
	cd "$pkgname-$pkgver"
	install -Dm755 "bin/aurgasm" "$pkgdir"/usr/bin/aurgasm
}
