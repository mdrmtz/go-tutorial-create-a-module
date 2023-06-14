# Tutorial: Create a Go module

## Table of Contents

- [Prerequisites](#prerequisites)
- [Start a module that others can use](#start)

This is the first part of a tutorial that introduces a few fundamental features of the Go language. If you're just getting started with Go, be sure to take a look at [Tutorial: Get started with Go](/doc/tutorial/getting-started.html), which introduces the `go` command, Go modules, and very simple Go code.

In this tutorial, you'll create two modules. The first is a library which is intended to be imported by other libraries or applications. The second is a caller application which will use the first.

This tutorial's sequence includes seven brief topics that each illustrate a different part of the language.

1. Create a module -- Write a small module with functions you can call from another module.
2. [Call your code from another module](/doc/tutorial/call-module-code.html) -- Import and use your new module.
3. [Return and handle an error](/doc/tutorial/handle-errors.html) -- Add simple error handling.
4. [Return a random greeting](/doc/tutorial/random-greeting.html) -- Handle data in slices (Go's dynamically-sized arrays).
5. [Return greetings for multiple people](/doc/tutorial/greetings-multiple-people.html) -- Store key/value pairs in a map.
6. [Add a test](/doc/tutorial/add-a-test.html) -- Use Go's built-in unit testing features to test your code.
7. [Compile and install the application](/doc/tutorial/compile-install.html) -- Compile and install your code locally.

**Note:** For other tutorials, see [Tutorials](/doc/tutorial/index.html).

## Prerequisites

- **Some programming experience.** The code here is pretty simple, but it helps to know something about functions, loops, and arrays.
- **A tool to edit your code.** Any text editor you have will work fine. Most text editors have good support for Go. The most popular are VSCode (free), GoLand (paid), and Vim (free).
- **A command terminal.** Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

