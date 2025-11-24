# Tasks: CLI Tic-Tac-Toe Game

**Input**: Design documents from `/specs/001-cli-tictactoe/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, quickstart.md

**Tests**: Per TicTacToe Constitution Principle II (Test-First Development), tests are NON-NEGOTIABLE. All user stories MUST include test tasks written BEFORE implementation (Red-Green-Refactor cycle). Target: 80% coverage with test pyramid (70% unit, 20% integration, 10% e2e).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Path Conventions

- Go project structure with packages: `game/`, `validation/`, `ui/`
- Tests alongside source: `game/board_test.go`, `validation/input_test.go`
- Config files at root: `Makefile`, `.golangci.yml`, `.github/workflows/build.yml`

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Initialize Go module with `go mod init github.com/YOUR_USERNAME/tictactoe`
- [x] T002 Create project directory structure (game/, validation/, ui/, .github/workflows/)
- [x] T003 [P] Create Makefile with targets: build, test, lint, clean, build-all, coverage, fmt, install
- [x] T004 [P] Create .golangci.yml with enabled linters (gofmt, govet, staticcheck, errcheck, gocyclo, funlen)
- [x] T005 [P] Create .gitignore for Go (bin/, *.exe, go.sum, coverage files)
- [x] T006 [P] Create .github/workflows/build.yml for CI/CD (lint, test, build matrix for Linux/macOS/Windows)

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core data types that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T007 [P] Unit tests for Cell type in game/board_test.go (test Empty, X, O constants, String() method)
- [x] T008 Define Cell type and constants (Empty, X, O) in game/board.go
- [x] T009 [P] Unit tests for Board type in game/board_test.go (NewBoard, GetCell, SetCell, IsCellEmpty, IsFull)
- [x] T010 Implement Board type [3][3]Cell and basic operations in game/board.go
- [x] T011 [P] Unit tests for Player type in game/player_test.go (Player1, Player2, GetMark, Name, Other)
- [x] T012 Define Player type and constants (Player1, Player2) in game/player.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Play Complete Game (Priority: P1) 🎯 MVP

**Goal**: Two players can play a complete tic-tac-toe game with turn alternation, board display, win detection (horizontal, vertical, diagonal), and draw detection

**Independent Test**: Run the game binary, have two players alternate turns with valid coordinates, verify correct win/draw determination and game end

### Tests for User Story 1 (REQUIRED per Constitution) ⚠️

> **CONSTITUTION REQUIREMENT: Write these tests FIRST, ensure they FAIL before implementation (Red-Green-Refactor)**
> **Target: 80% coverage | Test pyramid: 70% unit, 20% integration, 10% e2e | Suite: <60s**

- [x] T013 [P] [US1] Unit tests for win detection in game/win_test.go (test all 8 win patterns: 3 rows, 3 cols, 2 diagonals, no win, partial boards)
- [x] T014 [P] [US1] Unit tests for draw detection in game/win_test.go (test full board no winner, partial boards)
- [x] T015 [P] [US1] Unit tests for MakeMove in game/board_test.go (test move application, turn switching, invalid moves)
- [x] T016 [US1] Integration test for complete game in integration_test.go (test full game flow: moves → win, moves → draw)

### Implementation for User Story 1

- [x] T017 [P] [US1] Implement CheckWin function in game/win.go (check 3 rows, 3 cols, 2 diagonals - O(1) algorithm)
- [x] T018 [P] [US1] Implement CheckDraw function in game/win.go (return true if board full and no winner)
- [x] T019 [US1] Implement MakeMove and turn management in game/board.go (apply move, switch player, maintain move count)
- [x] T020 [US1] Implement CLI app in main.go (display board, handle input, game loop)
- [x] T021 [US1] Implement board display function (render 3x3 grid with X/O/Empty)
- [x] T022 [US1] Implement game loop and win/draw announcement (check state after each move)
- [x] T023 [US1] Complete main entry point with input validation and error handling

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently. Two players can play a complete game with proper win/draw detection.

---

## Phase 4: User Story 2 - Input Validation and Error Handling (Priority: P2)

**Goal**: Players receive clear, specific error messages for invalid input (out-of-range, non-numeric, occupied cell, incomplete) and game re-prompts without switching turns

**Independent Test**: Run the game, deliberately enter invalid inputs (0, 4, "abc", occupied cells), verify appropriate error messages display and game continues with same player

### Tests for User Story 2 (REQUIRED per Constitution) ⚠️

> **CONSTITUTION REQUIREMENT: Write these tests FIRST, ensure they FAIL before implementation (Red-Green-Refactor)**
> **Target: 80% coverage | Test pyramid: 70% unit, 20% integration, 10% e2e | Suite: <60s**

- [ ] T024 [P] [US2] Unit tests for input validation in validation/input_test.go (test range 1-3, numeric only, cell empty, all error cases)
- [ ] T025 [US2] Integration test for error handling in integration_test.go (test invalid input → error message → reprompt → valid input → continue)

### Implementation for User Story 2

- [ ] T026 [US2] Implement input validation functions in validation/input.go (ValidateRange, ValidateNumeric, ValidateEmpty, ParseInput with sentinel errors)
- [ ] T027 [US2] Integrate validation into input handling in ui/input_view.go (call validation before MakeMove, catch errors, re-prompt on error)
- [ ] T028 [US2] Implement error message display in ui/messages.go (display specific error messages as modals: ErrInvalidRange, ErrInvalidFormat, ErrCellOccupied, ErrIncompleteInput)

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently. Game has robust input validation with helpful error messages.

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories and deployment automation

- [ ] T029 [P] Add godoc comments to all exported functions in game/, validation/, ui/ packages
- [ ] T030 [P] Test build for all platforms with `make build-all` (verify Linux amd64/arm64, macOS amd64/arm64, Windows amd64 binaries)
- [ ] T031 [P] Run `make lint` and fix any linting errors (ensure gocyclo ≤10, funlen ≤50, no errcheck failures)
- [ ] T032 [P] Run `make coverage` and verify ≥80% coverage (aim for 90%+ in game/, 95%+ in validation/, 50-60%+ in ui/)
- [ ] T033 Test GitHub Actions workflow locally or in PR (verify lint, test, and build jobs pass)
- [ ] T034 Create README.md at repository root with quickstart instructions, build commands, and project description

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Story 1 (Phase 3)**: Depends on Foundational phase completion - No dependencies on other stories
- **User Story 2 (Phase 4)**: Depends on Foundational phase completion - Also depends on User Story 1 (needs working game loop to test validation)
- **Polish (Phase 5)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - MVP core gameplay
- **User Story 2 (P2)**: Can start after User Story 1 completes - Requires working game to add validation

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Tests for components can be written in parallel (marked [P])
- Implementation follows dependency order:
  - US1: Core types (T007-T012 Foundational) → Game logic (T017-T019) → UI (T020-T022) → Main (T023)
  - US2: Validation logic (T026) → UI integration (T027-T028)
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel (T003-T006)
- All Foundational test tasks marked [P] can run in parallel (T007, T009, T011)
- All tests for User Story 1 marked [P] can run in parallel (T013-T015)
- Multiple implementation tasks in US1 marked [P] (T017-T018, once tests pass)
- User Story 2 tests and polishing tasks can run in parallel once US1 completes

---

## Parallel Example: User Story 1

```bash
# After Foundational phase completes:

# Launch all unit tests for User Story 1 together:
Task: "Unit tests for win detection in game/win_test.go"
Task: "Unit tests for draw detection in game/win_test.go"
Task: "Unit tests for MakeMove in game/board_test.go"

# After tests written and failing, launch parallel implementations:
Task: "Implement CheckWin function in game/win.go"
Task: "Implement CheckDraw function in game/win.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (project structure, tooling)
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1 (core gameplay)
4. **STOP and VALIDATE**: Test User Story 1 independently
   - Two players can complete a game
   - Win detection works for all 8 patterns
   - Draw detection works
   - Board displays correctly
5. Deploy/demo MVP if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP - playable game!)
3. Add User Story 2 → Test independently → Deploy/Demo (robust validation and UX)
4. Add Polish → Final release

### Test-First Workflow (Constitution Requirement)

For each user story phase:
1. **Red**: Write all test tasks for the story, run tests, verify they FAIL
2. **Green**: Implement features to make tests PASS (minimal code)
3. **Refactor**: Clean up code while keeping tests green
4. **Verify**: Check coverage ≥80%, lint passes, tests <60s
5. Move to next story

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- User Story 1 is independently completable and testable (MVP)
- User Story 2 depends on US1 (needs working game to add validation)
- Verify tests fail before implementing (Red-Green-Refactor)
- Run `make test` after each task group to ensure tests pass
- Run `make lint` frequently to catch code quality issues early
- Commit after each completed task or logical group
- Stop at any checkpoint to validate story independently
- Constitution requirements: 80% coverage, <60s test suite, functions ≤50 lines, complexity ≤10

---

## Task Summary

**Total Tasks**: 34
- **Setup (Phase 1)**: 6 tasks
- **Foundational (Phase 2)**: 6 tasks (3 test + 3 impl)
- **User Story 1 (Phase 3)**: 11 tasks (4 test + 7 impl)
- **User Story 2 (Phase 4)**: 5 tasks (2 test + 3 impl)
- **Polish (Phase 5)**: 6 tasks

**Parallel Opportunities**: 15 tasks marked [P]

**Independent Testing**:
- US1: Complete game playable after Phase 3
- US2: Robust validation after Phase 4

**MVP Scope**: Phases 1-3 (Setup + Foundational + User Story 1) = 23 tasks
