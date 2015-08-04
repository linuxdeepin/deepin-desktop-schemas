PREFIX=/usr
SCHEMA_DIR=schemas
DEST=share/glib-2.0/schemas

install:
	mkdir -p ${PREFIX}/${DEST}
	cp -R ${SCHEMA_DIR} ${PREFIX}/${DEST}
