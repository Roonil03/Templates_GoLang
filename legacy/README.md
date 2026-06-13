# Legacy Codebase (Go 1.20 and Below)

This directory contains the legacy, non-generic implementations of the data structures. These files are maintained explicitly for backward compatibility with environments running Go versions 1.20 and below that lack support for modern type parameters or the standard `cmp` package.

## Contents
- **u/**: Legacy utility modules and auxiliary packages.
- **Data Structures**: Stripped, explicit-type variations (e.g., `int` based variants) of standard linear and tree structures.

For the modern, type-safe, generic implementations optimized for Go 1.21+, refer back to the root directory modules.
