
.PHONY: clean export assets build
clean:
	rm -rf dist
export:
	go run .
assets:
	# Ensure dist/assets exists
	mkdir -p dist/assets
	# Copy everything from public/assets into dist/assets
	[ -d public/assets ] && cp -R public/assets/* dist/assets/ || true
	# Copy everything else in public/ (HTML, docs, etc.)
	[ -d public ] && cp -R public/* dist/ || true
build: clean export assets
	@echo "Built static site in ./dist"
