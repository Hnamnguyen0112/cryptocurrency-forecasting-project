build:
	bazel build //...

test:
	bazel test //...

gazelle:
	bazel run //:gazelle

update:
	bazel run //:gazelle-update-repos
