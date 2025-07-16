# Gemini CLI Agent Instructions for 'self-taught' Repository

This document outlines specific instructions for the Gemini CLI agent when interacting with the 'self-taught' repository.

## Q&A Logging Policy

When a user asks a question related to a specific technology or library within this repository (e.g., `golang/third-party/gin`, `kubernetes/some-component`), the following procedure must be followed to log the Q&A:

1. **Determine Target `QA.md` File:**
   * If the questions for a language/technology are primarily about its core aspects and do not branch into multiple distinct sub-topics, create a single `QA.md` file in the language's root directory (e.g., `golang/QA.md`).
   * If there are multiple distinct sub-topics (e.g., `third-party/gin`, `third-party/testify`), create a `QA/` directory within the language's root, and then create separate Markdown files for each topic within that directory. The filename should be the topic name (e.g., `golang/QA/gin.md`, `golang/QA/testify.md`).

2. **Identify Content Heading:**
   * For single `QA.md` files, the main topic (e.g., "Golang Source Code") can be a second-level heading (`## Golang Source Code`).
   * For files within a `QA/` directory, the file itself represents the topic, so the main heading inside the file should be a first-level heading (`# Gin Framework Q&A`).
   * Specific questions within any `QA.md` file should be third-level headings (`### Q:`).

3. **Format as Q&A Sections:** Each question should be a third-level heading (`### Q:`) followed by its answer. Answers can include multiple lines, code blocks, or lists.
4. **Group Related Content:** If multiple questions are asked about the same sub-topic, group them together under the same heading.
5. **Continuous Update:** The `QA.md` file(s) should be updated with new Q&A as they occur. If the file(s) do not exist, create them. If they exist, append or update the relevant section.

## Code and Markdown Quality Standards

As a senior-level software engineer, I must ensure all generated content adheres to high quality standards:

* **Markdown Linting:** All generated Markdown files (e.g., `QA.md`) must strictly conform to common Markdown linting standards (e.g., `markdownlint` rules). This includes proper spacing, heading levels, list formatting, and code block syntax.
* **Code Best Practices:** All generated code snippets must follow industry best practices, idiomatic patterns, and senior-level coding standards for the respective language. This includes considerations for readability, maintainability, error handling, and efficiency.

**Example `QA.md` Structure (for `golang/QA/gin.md`):**

```markdown
# Gin Framework Q&A

### Q: How to initialize a Gin router?


A:

    ```go
    r := gin.Default()
    ```

```
