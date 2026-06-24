# Mobile + Backend Audit Report - 2026-06-23

## Summary

Audit scope: `mobile/` Android client and `api/` Go backend. `web/` and existing `test-report/` were used only as reference material.

Overall status:

| Area | Status | Evidence |
|---|---:|---|
| Backend unit tests | PASS | `go test ./...` |
| Backend race tests | BLOCKED | `go test ./... -race` requires CGO and a C compiler; current env has `CGO_ENABLED=0` and no `gcc/clang/cl` on PATH |
| Android unit/build | PASS | `.\gradlew.bat :app:testDebugUnitTest :app:assembleDebug --rerun-tasks` |
| Android lint | FAIL | `lintDebug` found 2 errors, 604 warnings |
| Android instrumented test | BLOCKED | `connectedDebugAndroidTest` built APKs but failed: `No connected devices!` |
| Backend clean DB runtime | FAIL | Health works, but public data endpoints fail on missing tables |
| Mobile/backend local config | FAIL | `mobile/local.properties` points `BASE_URL` to port `8000`, while local backend runs on `3001` |

## Findings

### P0 - Clean backend bootstrap does not recreate required schema

On a disposable PostgreSQL 15 database, backend starts and `/api/health` returns 200, but the app does not create the full schema. `api/internal/database/database.go:41` only migrates `TranscriptBookmark` and `LearningHistory`, leaving core tables like `categories`, `users`, `vocabulary_categories`, and `vocabulary_decks` missing.

Smoke status against clean DB after starting backend:

| Endpoint | Status |
|---|---:|
| `/api/health` | 200 |
| `/api/v1/lessons` | 200 |
| `/api/v1/categories` | 500 |
| `/api/v1/lessons/1/transcripts` | 200 |
| `/api/v1/vocabulary-categories` | 500 |
| `/api/v1/vocabulary-system-decks` | 500 |
| protected endpoints without token | 401 |

Server log examples:

- `ERROR: relation "categories" does not exist`
- `ERROR: relation "vocabulary_categories" does not exist`
- `ERROR: relation "vocabulary_decks" does not exist`

Impact: a fresh backend/mobile environment cannot reliably serve category or vocabulary screens. This blocks meaningful full mobile runtime verification on clean local setup.

### P0 - Seed command reports success even when seed data fails

`go run ./cmd/seed` exits with code 0 and prints `Seed completed!` even when inserts fail. Error handling in `api/cmd/seed/main.go:53`, `105`, `157`, and `191` logs errors but does not return a failing exit code.

Observed on clean DB before full schema exists:

- `categories: ERROR: relation "categories" does not exist`
- `users: ERROR: relation "users" does not exist`
- command still ended with `SEED_EXIT=0`

Observed after backend auto-migrated partial schema:

- `users: ERROR: relation "users" does not exist`
- `lessons: 7 rows`
- `transcripts: 15 rows`
- command still ended with `SEED_AFTER_SERVER_EXIT=0`

Impact: CI/local setup can falsely pass while required seed data is absent.

### P1 - Mobile local BASE_URL does not match backend port

`mobile/local.properties` has:

```properties
BASE_URL=http://10.0.2.2:8000/api/v1/
```

The backend README and audited server run on `http://localhost:3001`, so Android emulator should use:

```properties
BASE_URL=http://10.0.2.2:3001/api/v1/
```

Impact: mobile running on emulator will call port `8000`, not the audited local backend on `3001`, unless an undocumented proxy is running on `8000`.

### P1 - Shadowing pronunciation API contract is only partially implemented

Mobile declares `PronunciationApi.assessPronunciation()` at `mobile/app/src/main/java/com/example/app/data/remote/api/PronunciationApi.java:21` as an assessment API. It uploads:

- `audio`
- `referenceText`
- `lessonId`
- `transcriptId`

Mobile expects `PronunciationResponse` fields like `text`, `overallScore`, `scores`, `feedback`, `words`, and `attempt` at `mobile/app/src/main/java/com/example/app/data/remote/model/response/pronunciation/PronunciationResponse.java:5`.

Backend `TranscribeShadowing` at `api/internal/modules/shadowing_status/shadowing_status.controller.go:105` only reads form file `audio`, ignores `referenceText`, `lessonId`, and `transcriptId`, then returns only:

```json
{ "transcribed_text": "..." }
```

Impact: shadowing can transcribe audio, but there is no backend pronunciation scoring/feedback/attempt contract matching the mobile model. Score UI and feedback cannot be considered verified.

### P1 - Android lint fails the build

`lintDebug` fails with 2 errors and 604 warnings. Blocking errors:

| File | Line | Issue |
|---|---:|---|
| `mobile/app/src/main/res/layout/fragment_dictation.xml` | 172 | Must use `app:tint` instead of `android:tint` |
| `mobile/app/src/main/res/layout/fragment_listening.xml` | 172 | Must use `app:tint` instead of `android:tint` |

The warnings include many hardcoded strings, default-locale formatting calls, dependency version warnings, and WebView JavaScript security warnings.

Impact: CI or release checks that include lint will fail.

### P2 - Backend race test cannot run on current Windows environment

`go test ./... -race` first failed because race requires CGO. With `CGO_ENABLED=1`, it failed with:

```text
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
```

Impact: race safety is unverified locally until a C compiler is installed and CGO is enabled.

### P2 - Full mobile runtime/manual audit is blocked by environment

ADB is available and AVDs exist (`Pixel_8`, `Small_Phone`), but no device was connected:

```text
List of devices attached
```

`connectedDebugAndroidTest` result:

```text
DeviceException: No connected devices!
```

Manual login/runtime flows were not executed because there was no online emulator/device and no confirmed Clerk test credentials for mobile login.

## Mobile ↔ Backend Contract Matrix

| Mobile API | Backend route | Auth | Static contract | Runtime status |
|---|---|---:|---|---|
| `POST auth/sync` | `POST /api/v1/auth/sync` | Yes | Match | Blocked: needs Clerk token |
| `POST auth/complete-signup` | `POST /api/v1/auth/complete-signup` | Yes | Match | Blocked: needs Clerk token |
| `GET user/profile` | `GET /api/v1/user/profile` | Yes | Match | 401 without token |
| `PUT user/profile` | `PUT /api/v1/user/profile` | Yes | Match | Blocked: needs Clerk token |
| `GET lessons` | `GET /api/v1/lessons` | No | Match | 200 on clean DB, but empty because seed/bootstrap incomplete |
| `GET lessons/{lessonId}` | `GET /api/v1/lessons/:lessonId` | No | Match | Not fully verified with seeded data |
| `POST admin/lessons` | `POST /api/v1/admin/lessons` | Admin | Match | Blocked: needs admin token |
| `GET categories` | `GET /api/v1/categories` | No | Match | 500 on clean DB, missing `categories` table |
| `POST admin/categories` | `POST /api/v1/admin/categories` | Admin | Match | Blocked: needs admin token |
| `DELETE admin/categories/{id}` | `DELETE /api/v1/admin/categories/:id` | Admin | Match | Blocked: needs admin token |
| `GET lessons/{lessonId}/transcripts` | `GET /api/v1/lessons/:lessonId/transcripts` | No | Match | 200 on clean DB, but empty because seed/bootstrap incomplete |
| `GET transcript-bookmarks` | `GET /api/v1/transcript-bookmarks` | Yes | Match | 401 without token |
| `GET transcript-bookmarks/{lessonId}` | `GET /api/v1/transcript-bookmarks/:lessonId` | Yes | Match | 401 without token |
| `POST transcript-bookmarks` | `POST /api/v1/transcript-bookmarks` | Yes | Match | Blocked: needs token |
| `PUT transcript-bookmarks/{id/transcriptId}` | `PUT /api/v1/transcript-bookmarks/:transcriptId` | Yes | Match path, naming inconsistent in mobile clients | Blocked: needs token |
| `DELETE transcript-bookmarks/{id/transcriptId}` | `DELETE /api/v1/transcript-bookmarks/:transcriptId` | Yes | Match path, naming inconsistent in mobile clients | Blocked: needs token |
| `GET learning-history` | `GET /api/v1/learning-history` | Yes | Match | 401 without token |
| `POST learning-history` | `POST /api/v1/learning-history` | Yes | Match | Blocked: needs token |
| `GET learning-history/finished` | `GET /api/v1/learning-history/finished` | Yes | Match | Blocked: needs token |
| `GET learning-history/unfinished` | `GET /api/v1/learning-history/unfinished` | Yes | Match | Blocked: needs token |
| `GET learning-history/summary` | `GET /api/v1/learning-history/summary` | Yes | Match | Blocked: needs token |
| `GET learning-history/lessons/{lessonId}/summary` | `GET /api/v1/learning-history/lessons/:lessonId/summary` | Yes | Match | Blocked: needs token |
| `GET learning-history/{lessonId}` | `GET /api/v1/learning-history/:lessonId` | Yes | Match | Blocked: needs token |
| `GET dictation-status?lesson_id=` | `GET /api/v1/dictation-status` | Yes | Match | 401 without token |
| `POST dictation-status` | `POST /api/v1/dictation-status` | Yes | Match request body `transcript_id`, `lesson_id` | Blocked: needs token |
| `POST shadowing-status/transcribe` | `POST /api/v1/shadowing-status/transcribe` | Yes | Mismatch: backend only returns transcription, not mobile scoring model | Blocked: needs token and Deepgram |
| `GET shadowing-status?lesson_id=` | `GET /api/v1/shadowing-status` | Yes | Match | 401 without token |
| `POST shadowing-status` | `POST /api/v1/shadowing-status` | Yes | Match request body `transcript_id`, `lesson_id` | Blocked: needs token |
| `GET vocabulary-categories` | `GET /api/v1/vocabulary-categories` | No | Match | 500 on clean DB, missing table |
| `GET vocabulary-system-decks?category_id=` | `GET /api/v1/vocabulary-system-decks` | No | Match | 500 on clean DB, missing table |
| `GET vocabulary-decks` | `GET /api/v1/vocabulary-decks` | Yes | Match | 401 without token |
| `POST vocabulary-decks` | `POST /api/v1/vocabulary-decks` | Yes | Match | Blocked: needs token |
| `PUT vocabulary-decks/{id}` | `PUT /api/v1/vocabulary-decks/:id` | Yes | Match | Blocked: needs token |
| `DELETE vocabulary-decks/{id}` | `DELETE /api/v1/vocabulary-decks/:id` | Yes | Match | Blocked: needs token |
| `GET vocabulary-decks/{deckId}/items` | `GET /api/v1/vocabulary-decks/:id/items` | No | Match | Not fully verified with seeded data |
| `POST vocabulary-decks/{deckId}/items` | `POST /api/v1/vocabulary-decks/:id/items` | Yes | Match | Blocked: needs token |
| `PUT vocabulary-items/{itemId}` | `PUT /api/v1/vocabulary-items/:id` | Yes | Match | Blocked: needs token |
| `DELETE vocabulary-items/{itemId}` | `DELETE /api/v1/vocabulary-items/:id` | Yes | Match | Blocked: needs token |

## Command Evidence

Commands run from `D:\Backend_mobile_cover_parroto`:

```powershell
cd api
go test ./...
go test ./... -race
$env:CGO_ENABLED='1'; go test ./... -race
```

Backend results:

- `go test ./...`: PASS
- `go test ./... -race`: BLOCKED, CGO disabled
- `CGO_ENABLED=1 go test ./... -race`: BLOCKED, `gcc` missing

Android commands:

```powershell
cd mobile
$env:JAVA_HOME='C:\Program Files\Android\Android Studio\jbr'
$env:Path="$env:JAVA_HOME\bin;$env:Path"
.\gradlew.bat :app:testDebugUnitTest :app:assembleDebug --rerun-tasks --console=plain
.\gradlew.bat :app:lintDebug --console=plain
.\gradlew.bat :app:connectedDebugAndroidTest --console=plain
```

Android results:

- `testDebugUnitTest + assembleDebug`: PASS
- `lintDebug`: FAIL, 2 errors, 604 warnings
- `connectedDebugAndroidTest`: FAIL/BLOCKED, no connected devices

Backend runtime smoke was run against a disposable PostgreSQL container on port `54329`; container was removed after the audit.

## Recommended Next Steps

1. Add real schema bootstrap: migrations or full `AutoMigrate` for every model needed by backend/mobile, then make clean DB smoke part of CI.
2. Make `cmd/seed` fail fast or return non-zero when any seed insert fails.
3. Align `mobile/local.properties` `BASE_URL` with local backend port `3001`, or document/start the required port `8000` proxy.
4. Decide whether shadowing should be transcription-only or pronunciation scoring. If scoring is intended, backend response must match mobile `PronunciationResponse`.
5. Fix the 2 lint errors, then triage the 604 warnings by risk: hardcoded text, locale formatting, WebView JS security, dependency drift.
6. Re-run mobile runtime with a real emulator/device and Clerk test credentials after clean DB bootstrap is fixed.
