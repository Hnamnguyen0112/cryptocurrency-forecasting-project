load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/Hnamnguyen0112/cryptocurrency-forecasting-project
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
