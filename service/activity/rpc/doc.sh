# from https://github.com/pseudomuto/protoc-gen-doc
docker run --rm \
  -v $(pwd)/doc:/out \
  -v $(pwd):/protos \
  pseudomuto/protoc-gen-doc --doc_opt=markdown,docs.md
