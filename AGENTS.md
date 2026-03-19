# AGENTS.md

## Core Principles

### Understand before acting
- Read the relevant code and understand the full context before making changes.
- Ask clarifying questions if the task is ambiguous — do not guess intent.
- When in doubt, do less and confirm with the developer.

### Fix the root cause, not the symptom
- When debugging, identify **why** the problem occurs, not just where.
- Do not suppress errors, mask them with fallbacks, or work around them unless explicitly asked.
- If a quick fix would introduce technical debt, flag it.

### Minimal footprint
- Make the smallest change that correctly solves the problem.
- Do not refactor unrelated code, rename things for style, or add features not requested.
- Avoid introducing new dependencies unless necessary and discussed.

### Leave the code better than you found it
- Fix obvious issues you notice in code you touch (typos, dead code, wrong comments).
- Do not do large unsolicited cleanups — flag them instead.

---

## Code Quality

### Types are contracts — respect them
- Always follow the type system of the language (TypeScript, Go, etc.).
- Avoid type casting and type assertions unless there is no alternative.
- If a type is wrong, fix the type — do not cast around it.
- Never use `any` in TypeScript. Use `unknown` with proper narrowing if the type is genuinely unknown.

### No magic values
- Never hardcode business logic values, limits, timeouts, or strings inline.
- Extract them to named constants with clear, descriptive names.

### Explicit over implicit
- Prefer explicit returns, explicit error handling, explicit types.
- Avoid clever one-liners that obscure intent.
- Code should be readable by a developer unfamiliar with the codebase.

### Handle errors properly
- Never silently swallow errors.
- Every error must be either handled with appropriate logic or propagated to the caller.
- Log errors with enough context to debug them (what happened, where, with what input).

### No dead code
- Do not leave commented-out code, unused imports, or unused variables.
- If code might be needed later, note it in a comment or a ticket — do not keep it in source.

---

## Functions and Structure

### Single responsibility
- Each function, method, or module should do one thing and do it well.
- If a function needs a long comment to explain what it does, it should probably be split.

### Pure functions where possible
- Prefer functions without side effects.
- Side effects (I/O, mutations, external calls) should be isolated and explicit.

### Avoid deep nesting
- Maximum 2–3 levels of nesting. Use early returns to reduce nesting.
- Extract deeply nested logic into named helper functions.

### Consistent naming
- Use names that describe **what** the thing is or does, not **how** it works.
- Be consistent with the naming conventions already used in the codebase.
- Boolean variables and functions should read like a question: `isLoading`, `hasError`, `canSubmit`.

---

## Testing

### Write tests for new logic
- Any new business logic, utility function, or non-trivial transformation should have tests.
- Tests should cover: the happy path, edge cases, and known failure scenarios.

### Tests are first-class code
- Apply the same quality standards to test code as to production code.
- Do not copy-paste test boilerplate without adapting it.

### Do not break existing tests
- Run the test suite before considering a task done.
- If a test fails due to a legitimate change in behavior, update the test and document why.

---

## Security

### Never hardcode secrets
- No API keys, tokens, passwords, or credentials in source code — ever.
- Use environment variables or a secrets manager.

### Validate all input
- Treat all external input as untrusted: user input, query params, API responses, file contents.
- Validate and sanitize at the boundary — do not assume callers do it.

### Principle of least privilege
- Request only the permissions and access scopes the feature actually needs.
- Do not store sensitive data longer than necessary.

---

## Communication

### Explain non-obvious decisions
- If you make a choice that isn't obviously the best one, add a short comment explaining why.
- If you skip something intentionally (e.g., skipping a check because it's handled upstream), say so.

### Flag concerns explicitly
- If a task as described would introduce a bug, security issue, or bad pattern — say so before doing it.
- Use `TODO:`, `FIXME:`, or `NOTE:` comments to surface issues you cannot fix in the current scope.

### Do not hallucinate APIs
- If you are unsure whether a function, method, or library feature exists, say so.
- Do not invent plausible-looking code that may not work — verify or flag uncertainty.
