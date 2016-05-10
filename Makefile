PREFIX = /usr
ARCH = x86

all: build

build:
	@echo build with Arch: $(ARCH)
	mkdir -p result
	@find schemas -name "*.xml" -exec cp {} result \;
	@find -L overrides/$(ARCH) -name "*.override" -exec cat {} \; >> result/99_deepin-default-gsettings.gschema.override

test: 
	@echo "Testing schemas with glib-compile-shemas..."
	@glib-compile-schemas --dry-run result

install:
	mkdir -p $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas
	@cp result/* $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/

clean:
	@-rm -rf result
