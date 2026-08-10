## Description
This Pull Request introduces a fully automated CI/CD pipeline via GitHub Actions to build, package, and release TMSU natively for Windows.

## Technical Implementation Details
Compiling TMSU on Windows natively requires specific environmental setups due to its dependencies:
* CGo & SQLite: The `go-sqlite3` dependency requires a C compiler. Microsoft's MSVC is incompatible with this specific CGo integration. Therefore, this workflow utilizes MSYS2 with the MinGW-w64 UCRT64 toolchain (`mingw-w64-ucrt-x86_64-gcc`).
* VFS Bypass: Because Windows lacks native Linux FUSE support, the build command explicitly passes the `-tags novfs` flag to cleanly omit the VFS components, ensuring a successful compilation of the core tagging features.

## Artifacts Generated
Upon pushing a new `v*` tag, the workflow automatically generates and attaches the following to a GitHub Release:
1. Portable Archive (`tmsu-windows-portable.zip`): A standalone executable alongside documentation.
2. Windows Installer (`tmsu-windows-installer.exe`): A fully silent-capable setup generated using Inno Setup 6, configured to install reliably to the system's Program Files directory.

## Verification
The generated artifacts have been rigorously tested locally inside isolated Windows Server Core Docker containers to ensure clean execution and correct installation paths.

## Related Issues
* Resolves #34
* Relates to #115
* Relates to #137
