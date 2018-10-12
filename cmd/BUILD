load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
    ],
    importpath = "github.com/thelark/hyperchain-go-sdk/cmd",
    deps = [
        "//api:go_default_library",
        "@in_gopkg_urfave_cli_v2//:go_default_library",
    ],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "hyperchain-cli",
    embed = [":go_default_library"],
    visibility = ["//visibility:private"],
)
