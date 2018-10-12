load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_test(
    name = "go_default_test",
    srcs = ["api_test.go"],
    deps = [":go_default_library"],
    importpath = "github.com/thelark/hyperchain-go-sdk/api",
)

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "errors.go",
        "type.go",
    ],
    importpath = "github.com/thelark/hyperchain-go-sdk/api",
    deps = [
    ],
    visibility = ["//visibility:public"],
)

