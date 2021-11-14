load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/anchamber/lab",
)

go_library(
    name = "lab",
    srcs = ["main.go"],
    importpath = "github.com/anchamber/lab",
    visibility = ["//visibility:public"],
)
