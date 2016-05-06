PREFIX = /usr
performance = normal

all: test

prepare:
	@echo build with performance: $(performance)
	mkdir -p result
	@find schemas -name "*.xml" -exec cp {} result \;
	@find overrides -name "*.override" -exec cat {} \; > result/99_deepin-default-gsettings.gschema.override
	@find overrides -name "*.override.$(performance)" -exec cat {} \; >> result/99_deepin-default-gsettings.gschema.override

test: prepare 
	@echo "Testing schemas with glib-compile-shemas..."
	@glib-compile-schemas --dry-run result

install:
	mkdir -p $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas
	@cp result/* $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/

clean:
	@-rm -rf result
