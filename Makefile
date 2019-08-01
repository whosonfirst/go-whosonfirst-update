src:
	if test ! -d cmd/wof-bundles-index; then mkdir -p wof-bundles-index; fi
	if test ! -d cmd/wof-dist-index; then mkdir -p wof-dist-index; fi
	if test ! -d cmd/wof-dist-prune; then mkdir -p wof-dist-prune; fi
	if test ! -d cmd/wof-dist-publish; then mkdir -p wof-dist-publish; fi
	if test ! -d cmd/wof-list-repos; then mkdir -p wof-list-repos; fi
	if test ! -d cmd/wof-s3-sync; then mkdir -p wof-s3-sync; fi

