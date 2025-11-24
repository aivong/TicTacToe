<!--
Sync Impact Report:
Version change: Initial → 1.0.0
Modified principles: N/A (initial constitution)
Added sections:
  - Core Principles (5 principles covering code quality, testing, UX, performance, maintainability)
  - Quality Gates (Pre-Implementation, Pre-Merge, Pre-Deployment)
  - Development Workflow (Feature Development Process, Code Review Standards, Branch Strategy)
  - Governance (Constitution Authority, Amendment Process, Compliance)
Templates updated:
  ✅ .specify/templates/plan-template.md - Constitution Check section expanded with specific checklist items for all 5 principles
  ✅ .specify/templates/spec-template.md - Success Criteria section enhanced with Constitution Compliance subsection
  ✅ .specify/templates/tasks-template.md - Tests changed from OPTIONAL to REQUIRED (NON-NEGOTIABLE), test pyramid targets added, task IDs renumbered
Follow-up TODOs: None
Rationale: This constitution establishes non-negotiable standards to prevent technical debt, ensure consistent quality, and maintain high UX/performance standards from day one.
-->

# TicTacToe Constitution

## Core Principles

### I. Code Quality (NON-NEGOTIABLE)

Code must be clean, readable, and maintainable. Every commit must meet these standards:

- **Single Responsibility**: Each function/class/module has one clear purpose
- **Descriptive Naming**: Names reveal intent without needing comments (e.g., `calculateWinningPositions` not `calc`)
- **No Magic Numbers**: Constants must be named (e.g., `BOARD_SIZE = 3` not hardcoded `3`)
- **DRY Principle**: No code duplication beyond 3 lines - extract to reusable functions
- **Maximum Function Length**: 50 lines per function; exceeding requires written justification
- **Cyclomatic Complexity**: Maximum complexity of 10; higher values require refactoring
- **Code Reviews**: All code must pass peer review before merge

**Rationale**: Technical debt accumulates exponentially. Preventing it at source is 10x cheaper than fixing it later.

### II. Test-First Development (NON-NEGOTIABLE)

Testing is not optional. Tests must be written BEFORE implementation:

- **Red-Green-Refactor Cycle**: Write failing test → Implement → Pass test → Refactor → Repeat
- **Test Coverage Minimum**: 80% code coverage for all new features; exceptions require approval
- **Test Pyramid**:
  - 70% unit tests (fast, isolated, test single functions)
  - 20% integration tests (test component interactions)
  - 10% end-to-end tests (test complete user flows)
- **Test Independence**: Each test must be runnable in isolation without side effects
- **Fast Test Suite**: Unit tests must complete in <5 seconds; full suite in <60 seconds
- **Test Documentation**: Each test must clearly document what scenario it covers

**Rationale**: Tests written after implementation are biased and miss edge cases. Test-first ensures design testability and catches defects early.

### III. User Experience Consistency (NON-NEGOTIABLE)

User experience must be consistent, intuitive, and delightful:

- **Response Time Standards**:
  - User actions: <100ms perceived response (loading indicators for longer operations)
  - Page loads: <1 second initial render
  - API responses: <200ms p95 latency
- **Error Handling**:
  - All errors must show user-friendly messages (never raw stack traces)
  - Clear actionable guidance (e.g., "Invalid move. Please select an empty cell.")
  - Graceful degradation (system remains usable despite errors)
- **Accessibility**:
  - Keyboard navigation support for all interactions
  - Screen reader compatibility (ARIA labels where needed)
  - Color contrast ratio minimum 4.5:1 (WCAG AA standard)
- **Visual Consistency**:
  - Single source of truth for design tokens (colors, spacing, typography)
  - Consistent component patterns across all screens
  - Responsive design for mobile, tablet, desktop

**Rationale**: Inconsistent UX erodes trust. Users judge quality by interface, not code elegance.

### IV. Performance Requirements (NON-NEGOTIABLE)

Performance is a feature, not an optimization:

- **Memory Management**:
  - Maximum memory footprint: <50MB for game state
  - No memory leaks (must pass leak detection tools)
  - Efficient data structures (appropriate algorithms for problem size)
- **Computational Efficiency**:
  - Game logic operations: O(1) or O(n) maximum complexity
  - AI move calculation: <500ms for any board state
  - No blocking operations on main thread/event loop
- **Network Efficiency** (if applicable):
  - Minimize request payload size (<10KB per request)
  - Implement request debouncing/throttling for rapid user actions
  - Cache static assets with appropriate TTLs
- **Battery Efficiency** (mobile):
  - Minimize background processing
  - Efficient rendering (60 fps minimum for animations)

**Performance Budgets**: Any feature exceeding these limits requires optimization or architectural review.

**Rationale**: Poor performance directly impacts user satisfaction. Performance problems compound at scale.

### V. Maintainability & Documentation

Code must be maintainable by others without heroic effort:

- **Self-Documenting Code**: Code clarity > cleverness; simple solutions preferred over complex ones
- **Necessary Comments**: Document WHY, not WHAT (code explains what it does; comments explain why)
- **API Documentation**: All public functions/methods must have docstrings describing:
  - Purpose and behavior
  - Parameters (types, constraints, defaults)
  - Return values
  - Example usage
- **Architecture Documentation**: Keep architecture diagrams current (update with significant changes)
- **Change Documentation**: Non-trivial PRs must include:
  - Problem being solved
  - Approach chosen and alternatives considered
  - Testing strategy
  - Migration plan (if breaking changes)

**Rationale**: Code is read 10x more than written. Optimize for reader comprehension.

## Quality Gates

All features must pass these gates before deployment:

### Pre-Implementation Gate
- [ ] Feature specification approved
- [ ] Test cases documented
- [ ] Performance requirements defined
- [ ] UX design reviewed (for user-facing features)

### Pre-Merge Gate
- [ ] All tests passing (unit, integration, e2e)
- [ ] Code coverage ≥80% for new code
- [ ] Code review approved by ≥1 reviewer
- [ ] No critical linting/formatting violations
- [ ] Performance benchmarks met
- [ ] Documentation updated

### Pre-Deployment Gate
- [ ] End-to-end tests passing in staging environment
- [ ] Performance metrics within acceptable ranges
- [ ] Accessibility audit passed (for UI changes)
- [ ] Security scan passed (no critical vulnerabilities)

**Enforcement**: Automated CI/CD pipeline blocks deployment if gates fail.

## Development Workflow

### Feature Development Process

1. **Specification Phase**: Write clear, testable requirements with acceptance criteria
2. **Design Phase**: Document approach, identify edge cases, define test scenarios
3. **Test Phase**: Write failing tests for all scenarios (unit → integration → e2e)
4. **Implementation Phase**: Write minimal code to pass tests
5. **Refactor Phase**: Improve code quality while keeping tests green
6. **Review Phase**: Peer review for code quality, test coverage, documentation
7. **Integration Phase**: Merge to main branch after all gates pass

### Code Review Standards

Reviewers must verify:
- Code meets all principles in this constitution
- Tests are comprehensive and test the right things
- No obvious security vulnerabilities
- Performance implications considered
- Documentation is clear and accurate

Reviewers should NOT:
- Nitpick style (automate with linters)
- Approve code they don't understand
- Skip running the code locally

### Branch Strategy

- `main` branch: Always deployable, protected
- Feature branches: `###-feature-name` format (e.g., `001-game-logic`)
- All changes via Pull Requests (no direct commits to main)
- Squash commits on merge for clean history

## Governance

### Constitution Authority

This constitution supersedes all other practices, preferences, and "the way we've always done it." When in doubt, this document is the source of truth.

### Amendment Process

Constitution amendments require:
1. Written proposal with rationale
2. Review by all team members
3. Unanimous approval (or escalation to project lead)
4. Version bump (semantic versioning)
5. Migration plan for existing code (if applicable)
6. Update to all dependent templates and documentation

### Complexity Justification

When violating a principle is unavoidable:
1. Document the violation in plan.md Complexity Tracking table
2. Explain why it's needed
3. Document what simpler alternatives were considered and why they're insufficient
4. Get explicit approval before proceeding

### Compliance

- All Pull Requests must verify compliance with this constitution
- Reviewers have authority to block PRs that violate principles
- Technical debt must be tracked and prioritized for resolution
- Regular audits (monthly) to identify drift from standards

### Living Document

This constitution should evolve with the project. If a principle repeatedly blocks progress or proves impractical, propose an amendment rather than working around it.

**Version**: 1.0.0 | **Ratified**: 2025-11-23 | **Last Amended**: 2025-11-23
