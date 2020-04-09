load("@bazel_gazelle//:def.bzl", "gazelle")

genrule(
    name = "app",
    srcs = [
        "//cmd:server",
    ],
    outs = [
        "application",
    ],
    cmd = "cp ./bazel-out/*/bin/cmd/*/server /bazel/\"$@\"",
    executable = 1,
    local = 1,
)

gazelle(
    name = "gazelle",
    prefix = "github.com/r351574nc3/ultramagnus",
)