# wazerox: wazero extended with bleeding-edge features and likely bugs

[![WebAssembly Core Specification Test](https://github.com/wasilibs/wazerox/actions/workflows/spectest.yaml/badge.svg)](https://github.com/wasilibs/wazerox/actions/workflows/spectest.yaml) [![Go Reference](https://pkg.go.dev/badge/github.com/wasilibs/wazerox.svg)](https://pkg.go.dev/github.com/wasilibs/wazerox) [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[wazero][18] is a wonderful WebAssembly runtime for Go developers with a priority
on stability. This often means it can take time to add new incubating WebAssembly
features or completely non-standard ones - there is only so much a project can
take on.

wazerox is a fork with a priority on running code for the wasilibs organization.
The reality is much existing code out there requires features that are not
supported by current WebAssembly or WASI specifications. wazerox allows wasilibs
to continue to provide wrappers with zero usage of cgo.

The additions to wazero should be considered mostly be an implementation detail
of wasilibs libraries. If you have code that requires these extensions to execute,
it is OK to try wazerox, but do expect bugs and best-effort or non-existing support.
It is never a good idea to try this library before trying wazero and see if it
supports your binary, if it does, use wazero.

Current extensions include:

- [WebAssembly Threads][19]
- Stack checkpoint/restore for [wasix][20] exception handling

For further details on the runtime, read about [wazero][18].

[1]: https://www.w3.org/TR/2019/REC-wasm-core-1-20191205/
[2]: https://www.w3.org/TR/2022/WD-wasm-core-2-20220419/
[4]: https://github.com/WebAssembly/meetings/blob/main/process/subgroups.md
[5]: https://github.com/WebAssembly/WASI
[6]: https://pkg.go.dev/golang.org/x/sys/unix
[7]: https://github.com/WebAssembly/spec/tree/wg-1.0/test/core
[8]: internal/engine/compiler/RATIONALE.md
[9]: https://github.com/wasilibs/wazerox/issues/506
[10]: https://go.dev/doc/devel/release
[11]: https://github.com/actions/virtual-environments
[12]: https://docs.docker.com/develop/develop-images/baseimages/#create-a-simple-parent-image-using-scratch
[13]: https://github.com/WebAssembly/WASI/blob/snapshot-01/phases/snapshot/docs.md
[14]: https://github.com/WebAssembly/spec/tree/d39195773112a22b245ffbe864bab6d1182ccb06/test/core
[15]: https://tetrate.io/blog/introducing-wazero-from-tetrate/
[16]: https://wazero.io/community/users/
[17]: https://github.com/wasilibs/wazerox/stargazers
[18]: https://github.com/testratelabs/wazero
[19]: https://github.com/WebAssembly/threads
[20]: https://wasix.org/
