# 00 – Introduction

## What You Will Learn
What Go is, who made it and why, and what makes it worth learning in 2026. By the end of this lesson you should be able to explain Go in a sentence to a friend and know what kind of problems it's good at solving.

## Why This Matters
Before writing a single line of code, it helps to know *why* a language exists. Every language is a set of trade-offs its designers made on purpose. Understanding those trade-offs up front means the syntax you learn later (short variable names, explicit error handling, no exceptions, strict compiler rules) will make sense as deliberate choices instead of arbitrary rules to memorize.

## Concept Explanation
Go (often called **Golang**, mostly because `golang.org` was the only domain name available) is a statically typed, compiled programming language created at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Ken Thompson co-created UNIX and the B programming language — one of C's ancestors — so Go was designed by people with decades of experience building the infrastructure the rest of the software world runs on.

A few terms worth defining right away, since the rest of the course leans on them constantly:

- **Statically typed** — every value's type (whether it's a number, text, true/false, etc.) is known and checked *before* the program ever runs, at compile time. This catches a whole category of bugs early instead of at 2am in production.
- **Compiled** — Go source code is translated ahead of time into a single, standalone binary that the machine runs directly. This is different from a language like Python or JavaScript, where the code is read and executed line-by-line by an interpreter at runtime. Compiled code generally starts faster and runs faster.
- **Concurrency** — the ability for a program to make progress on multiple tasks at once (e.g., handling many network requests without one blocking the others). Go was designed from day one with lightweight tools for this, though we won't touch them until much later in the course.

Go was conceived in 2007, open-sourced in 2009, and reached its stable 1.0 release in 2012. The motivation was practical, not academic: Google had enormous codebases, huge engineering teams, and servers with many CPU cores, and existing languages were making that combination painful. C++ compiled fast but was complex to write and slow to build at scale; Python and similar languages were pleasant to write but slower to execute and lacked compile-time type safety. Go set out to keep the performance and safety of a compiled, statically typed language while keeping the syntax as simple and fast to write as a scripting language.

## Common Mistakes
- Assuming "Go" and "Golang" are different things — they're the same language; "Golang" is just the nickname that stuck because of the website's domain.
- Expecting Go to behave like Python or JavaScript because it "looks simple." Simple syntax does not mean loose rules — Go's compiler is strict (you'll see this concretely in Lesson 02).
- Thinking concurrency is something you need to understand before writing basic programs. It's a headline feature of Go, but it's an advanced topic — you can (and will) write plenty of useful Go code before ever touching a goroutine.

## Best Practices
- Read error messages from the Go compiler carefully — they are usually precise and tell you exactly what's wrong and where.
- Don't rush past the fundamentals to get to "the interesting stuff" like concurrency. Go's standard library and idioms reward a solid grasp of the basics.

## Real-World Example
Go is not an academic curiosity — it's the language behind Docker, Kubernetes, Terraform, and large parts of the infrastructure at Google, Uber, Dropbox, and Cloudflare. Companies chose it because it compiles to a single fast binary with no external runtime required, which makes it ideal for command-line tools, servers, and cloud infrastructure — exactly the kind of software those companies needed to build and deploy at massive scale.

## Summary
Go is a statically typed, compiled language built by experienced systems engineers at Google to make large-scale software development simpler, faster to build, and easier to run concurrently. It trades some flexibility for safety and speed, a theme that will show up again and again throughout this course.

## What to Learn Next
Next, you'll install Go on your machine and run your very first program.

Next: [../01Setup/README.md](../01Setup/README.md)
