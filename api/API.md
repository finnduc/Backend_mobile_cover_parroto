# Engflix API — Frontend Data Models

Base URL: `http://localhost:3001/api/v1`  
Auth: `Authorization: Bearer <clerk_session>` (required for protected endpoints)  
Pagination: `?page=1&limit=10` (defaults: page=1, limit=10)

---

## Category

```
GET  /categories                      (public)
```

```json
{ "id": 1, "name": "Daily Conversation" }
```

---

## Lesson

```
GET    /lessons?category_id=&level=&search=   (public)
GET    /lessons/:lessonId                      (public)
```

```json
{ "id": 1, "category_id": 1, "title": "Greeting People", "description": "...",
  "video_url": "https://...", "thumbnail_url": "https://...",
  "level": "beginner", "duration": 180, "created_at": "2025-01-01T00:00:00Z" }
```

Level values: `beginner`, `intermediate`, `advanced`

---

## Transcript

```
GET  /lessons/:lessonId/transcripts   (public)
```

```json
{ "id": 1, "lesson_id": 1, "sequence": 1, "content": "Hey, how are you doing?",
  "phonetic": "/heɪ haʊ ɑːr juː ˈduːɪŋ/", "vietnamese": "Này, bạn có khỏe không?",
  "start_timestamp": 0, "end_timestamp": 3.5 }
```

---

## Bookmark

```
GET    /bookmarks                     (auth)
POST   /bookmarks/:lessonId           (auth, no body)
DELETE /bookmarks/:lessonId           (auth, no body)
```

```json
{ "user_id": "clerk_xxx", "lesson_id": 1, "created_at": "...",
  "lesson": { "id": 1, "title": "...", "thumbnail_url": "...", "level": "beginner", "duration": 180 } }
```

---

## ShadowingStatus

```
POST  /shadowing-status/:transcriptId  (auth, no body, returns 409 if duplicate)
GET   /shadowing-status?lesson_id=     (auth)
```

```json
{ "user_id": "clerk_xxx", "transcript_id": 1, "lesson_id": 1, "completed_at": "..." }
```

---

## DictationStatus

```
POST  /dictation-status/:transcriptId  (auth, no body, idempotent — ON CONFLICT DO NOTHING)
GET   /dictation-status?lesson_id=     (auth)
```

```json
{ "user_id": "clerk_xxx", "transcript_id": 1, "lesson_id": 1, "completed_at": "..." }
```

---

## VocabularyCategory

```
GET   /vocabulary-categories                          (public)
// admin:
GET   /admin/vocabulary-categories                    (admin)
GET   /admin/vocabulary-categories/:id                 (admin)
POST  /admin/vocabulary-categories                    (admin)
PUT   /admin/vocabulary-categories/:id                 (admin)
DELETE /admin/vocabulary-categories/:id               (admin)
```

```json
{ "id": 1, "name": "Oxford Word List", "description": "...",
  "created_at": "...", "updated_at": "..." }
```

POST/PUT body: `{ "name": "...", "description": "..." }`

---

## VocabularyDeck

```
GET   /vocabulary-decks?category_id=                  (public)
POST  /vocabulary-decks                               (auth)
PUT   /vocabulary-decks/:id                            (auth, own deck only)
DELETE /vocabulary-decks/:id                           (auth, own deck only)
// admin:
GET   /admin/vocabulary-decks                         (admin)
POST  /admin/vocabulary-decks                         (admin, creates system deck)
PUT   /admin/vocabulary-decks/:id                      (admin)
DELETE /admin/vocabulary-decks/:id                     (admin)
```

```json
{ "id": 1, "user_id": null, "category_id": 1, "name": "Oxford A1",
  "description": "...", "thumbnail_url": "...", "level": "beginner",
  "is_default": true, "created_at": "...", "updated_at": "..." }
```

- `user_id: null` = system deck, `user_id: "clerk_xxx"` = user-created
- `is_default: true` for system decks

POST/PUT body: `{ "name": "...", "description": "...", "thumbnail_url": "...", "level": "..." }`

---

## VocabularyItem

```
GET   /vocabulary-decks/:deckId/items?page=&limit=    (public)
POST  /vocabulary-decks/:deckId/items                 (auth, own deck only)
PUT   /vocabulary-items/:id                            (auth, own item only)
DELETE /vocabulary-items/:id                           (auth, own item only)
// admin:
GET   /admin/vocabulary-decks/:deckId/items           (admin)
POST  /admin/vocabulary-decks/:deckId/items           (admin)
PUT   /admin/vocabulary-items/:id                      (admin)
DELETE /admin/vocabulary-items/:id                     (admin)
```

```json
{ "id": 1, "deck_id": 1, "lesson_id": null, "transcript_id": null,
  "phrase": "hello", "normalized_phrase": "hello",
  "meaning": "xin chào", "example_sentence": "Hello, how are you?",
  "note": "", "created_at": "...", "updated_at": "..." }
```

POST body: `{ "phrase": "...", "normalized_phrase": "...", "meaning": "...", "example_sentence": "...", "note": "...", "lesson_id": null, "transcript_id": null }`  
PUT body: same fields (no deck_id/lesson_id/transcript_id)

---

## Auth

```
POST /auth/complete-signup  (auth, no body)
```

Returns string: `"User profile finalized successfully"`

---

## FE Integration Flows

### Shadowing / Dictation Progress
```
1. GET /categories                          → display tabs
2. GET /lessons?category_id=X              → display lesson list
3. GET /lessons/:lessonId/transcripts      → load all sentences
4. For each sentence on complete:
   POST /shadowing-status/:transcriptId    → mark shadowing done
   POST /dictation-status/:transcriptId    → mark dictation done (client validates answer first)
5. GET /shadowing-status?lesson_id=X       → get completed IDs
   const progress = completed.length / totalTranscripts
```

### Vocabulary Browser
```
1. GET /vocabulary-categories              → show sections
2. GET /vocabulary-decks?category_id=X     → show decks
3. GET /vocabulary-decks/:id/items         → show words
4. POST /vocabulary-decks                  → user creates own deck
5. POST /vocabulary-decks/:id/items        → user adds word
```

---

## Status Codes

| Code | When |
|------|------|
| 200 | Success |
| 400 | Validation error (bad request body/params) |
| 401 | Missing/invalid auth token |
| 403 | No permission (wrong owner, not admin) |
| 404 | Resource not found |
| 409 | Duplicate (shadowing already completed) |
| 500 | Server error |
