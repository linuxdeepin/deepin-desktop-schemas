PREFIX = /usr

all: test

prepare:
	mkdir -p result
	@find schemas -name "*.xml" -exec cp {} result \;

test: prepare 
	@echo "Testing schemas with glib-compile-shemas..."
	@glib-compile-schemas --dry-run result

install:
	mkdir -p $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas
	@cp result/* $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/

clean:
	@-rm -rf result
