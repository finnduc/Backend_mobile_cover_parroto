# BÁO CÁO KIỂM THỬ HỆ THỐNG
**Dự án:** KTHT Backend (go-cover-parroto)  
**Ngày thực hiện:** 26/05/2026  
**Branch:** test/ui  
**Người thực hiện:** finnduc  

---

## 1. UNIT TEST - GOLANG

### Tổng quan
| Tiêu chí | Kết quả |
|---|---|
| Tổng số test case | **46** |
| PASS | **46** |
| FAIL | **0** |
| Trạng thái | ✅ TẤT CẢ PASS |

### Chi tiết từng module (test service)

| Package | Test Functions | Kết quả | Coverage |
|---|---|---|---|
| `internal/core/policy` | TestAllow_AdminPass, TestAllow_OwnerPass, TestAllow_NotOwnerForbidden, TestAllow_NoUserID, TestGetUserID_Success, TestGetUserID_Missing | ✅ PASS (6/6) | **100.0%** |
| `modules/bookmark/services` | TestBookmarkService_AddBookmark, TestBookmarkService_RemoveBookmark, TestBookmarkService_List | ✅ PASS (3/3) | **90.6%** |
| `modules/category/services` | TestCategoryService_List, GetByID, Create, Update, Delete | ✅ PASS (5/5) | **95.5%** |
| `modules/dictation_status/services` | TestDictationStatusService_Create, List | ✅ PASS (2/2) | **100.0%** |
| `modules/lesson/services` | TestLessonService_List, Get, Create, Update, Delete | ✅ PASS (5/5) | **92.5%** |
| `modules/shadowing_status/services` | TestShadowingStatusService_Create, List | ✅ PASS (2/2) | **100.0%** |
| `modules/transcript/services` | TestTranscriptService_GetByLesson, ReplaceByLesson, GetByID, Create, BulkCreate, Update, Delete | ✅ PASS (7/7) | **92.9%** |
| `modules/vocabulary_category/services` | TestVocabularyCategoryService_List, GetByID, Create, Update, Delete | ✅ PASS (5/5) | **100.0%** |
| `modules/vocabulary_deck/services` | TestVocabularyDeckService_ListDefault, ListByUser, GetByID, Create, Update, Delete, CreateAsSystem, DeleteAsSystem | ✅ PASS (8/8) | **76.7%** |
| `modules/vocabulary_item/services` | TestVocabularyItemService_List, GetByID, Create, Update, Delete | ✅ PASS (5/5) | **93.1%** |

### File chi tiết
- `test-report/go-unit/output.txt` — Full verbose output
- `test-report/go-unit/per-package-coverage.txt` — Coverage per package

---

## 2. SWAGGER / API DOCUMENTATION

### Thông tin API
| Tiêu chí | Kết quả |
|---|---|
| File spec | `api/cmd/server/docs/swagger.yaml` |
| Base Path | `/api/v1` |
| Tổng số endpoints | **51** |
| Format | Swagger 2.0 |

### Danh sách endpoint theo nhóm

#### Admin Endpoints (30 endpoints)
| Method | Path | Mô tả |
|---|---|---|
| GET | `/api/v1/admin/categories` | List categories |
| POST | `/api/v1/admin/categories` | Create category |
| GET | `/api/v1/admin/categories/{id}` | Get category by ID |
| PUT | `/api/v1/admin/categories/{id}` | Update category |
| DELETE | `/api/v1/admin/categories/{id}` | Delete category |
| GET | `/api/v1/admin/lessons` | List lessons |
| POST | `/api/v1/admin/lessons` | Create lesson |
| GET | `/api/v1/admin/lessons/{id}` | Get lesson by ID |
| PUT | `/api/v1/admin/lessons/{id}` | Update lesson |
| DELETE | `/api/v1/admin/lessons/{id}` | Delete lesson |
| GET | `/api/v1/admin/lessons/{lessonId}/transcripts` | Get lesson transcripts |
| PUT | `/api/v1/admin/lessons/{lessonId}/transcripts` | Replace all transcripts |
| POST | `/api/v1/admin/lessons/{lessonId}/transcripts/bulk` | Bulk create transcripts |
| POST | `/api/v1/admin/transcripts` | Create transcript |
| GET | `/api/v1/admin/transcripts/{id}` | Get transcript by ID |
| PUT | `/api/v1/admin/transcripts/{id}` | Update transcript |
| DELETE | `/api/v1/admin/transcripts/{id}` | Delete transcript |
| GET | `/api/v1/admin/vocabulary-categories` | List vocabulary categories |
| POST | `/api/v1/admin/vocabulary-categories` | Create vocabulary category |
| GET | `/api/v1/admin/vocabulary-categories/{id}` | Get vocabulary category by ID |
| PUT | `/api/v1/admin/vocabulary-categories/{id}` | Update vocabulary category |
| DELETE | `/api/v1/admin/vocabulary-categories/{id}` | Delete vocabulary category |
| GET | `/api/v1/admin/vocabulary-decks` | List all vocabulary decks |
| POST | `/api/v1/admin/vocabulary-decks` | Create system deck |
| GET | `/api/v1/admin/vocabulary-decks/{deckId}/items` | List items in system deck |
| POST | `/api/v1/admin/vocabulary-decks/{deckId}/items` | Add item to system deck |
| PUT | `/api/v1/admin/vocabulary-decks/{id}` | Update system deck |
| DELETE | `/api/v1/admin/vocabulary-decks/{id}` | Delete system deck |
| PUT | `/api/v1/admin/vocabulary-items/{id}` | Update item in system deck |
| DELETE | `/api/v1/admin/vocabulary-items/{id}` | Delete item from system deck |

#### User/Public Endpoints (21 endpoints)
| Method | Path | Mô tả |
|---|---|---|
| POST | `/api/v1/auth/complete-signup` | Finalize user registration |
| GET | `/api/v1/bookmarks` | List user bookmarks |
| POST | `/api/v1/bookmarks/{lessonId}` | Add bookmark |
| DELETE | `/api/v1/bookmarks/{lessonId}` | Remove bookmark |
| GET | `/api/v1/categories` | List categories |
| GET | `/api/v1/dictation-status` | List dictation status |
| POST | `/api/v1/dictation-status/{transcriptId}` | Mark transcript dictation completed |
| GET | `/api/v1/lessons` | List lessons |
| GET | `/api/v1/lessons/{lessonId}` | Get a lesson |
| GET | `/api/v1/lessons/{lessonId}/transcripts` | Get lesson transcripts |
| GET | `/api/v1/shadowing-status` | List shadowing status |
| POST | `/api/v1/shadowing-status/{transcriptId}` | Mark transcript shadowing completed |
| GET | `/api/v1/vocabulary-categories` | List vocabulary categories |
| GET | `/api/v1/vocabulary-decks` | List vocabulary decks |
| POST | `/api/v1/vocabulary-decks` | Create user vocabulary deck |
| GET | `/api/v1/vocabulary-decks/{deckId}/items` | List items in a deck |
| POST | `/api/v1/vocabulary-decks/{deckId}/items` | Add item to user deck |
| PUT | `/api/v1/vocabulary-decks/{id}` | Update user vocabulary deck |
| DELETE | `/api/v1/vocabulary-decks/{id}` | Delete user vocabulary deck |
| PUT | `/api/v1/vocabulary-items/{id}` | Update user's item |
| DELETE | `/api/v1/vocabulary-items/{id}` | Delete user's item |

### Kết quả validation Swagger
> **Công cụ:** @apidevtools/swagger-cli  
> **Kết quả:** ✅ Validation thành công — toàn bộ 51 endpoints hợp lệ, không có lỗi schema.

### File chi tiết
- `test-report/swagger-api/endpoints-list.txt` — Danh sách 51 endpoint
- `test-report/swagger-api/validation-output.txt` — Kết quả validation đầy đủ

---

## 3. PLAYWRIGHT E2E TEST (UI)

### Tổng quan
| Tiêu chí | Kết quả |
|---|---|
| Framework | Playwright |
| Thư mục test | `web/tests/e2e/` |
| Browser | Chromium |
| Base URL | http://localhost:3000 |

### Kết quả Public Tests (không cần đăng nhập)

| Test | Kết quả | Ghi chú |
|---|---|---|
| sign-in page renders Clerk sign-in component | ✅ PASS | Trang đăng nhập render đúng |
| sign-up page renders Clerk sign-up component | ✅ PASS | Trang đăng ký render đúng |
| unauthenticated user redirected to sign-in (lessons) | ✅ PASS | Redirect về `/sign-in` đúng |
| unauthenticated user redirected to sign-in (vocabulary) | ✅ PASS | Redirect về `/sign-in` đúng |
| unauthenticated user redirected to sign-in (admin) | ✅ PASS | Redirect về `/sign-in` đúng |

**Tóm tắt Public:** 5 PASS / 0 FAIL

### Kết quả User Tests (đăng nhập user thường)

| Test | Kết quả | Ghi chú |
|---|---|---|
| user-setup — authenticate as user | ✅ PASS | Đăng nhập thành công, lưu storage state |
| bookmarks — hiển thị danh sách bookmark | ✅ PASS | |
| bookmarks — thêm bookmark bài học | ✅ PASS | |
| bookmarks — xoá bookmark bài học | ✅ PASS | |
| categories — hiển thị danh sách category | ✅ PASS | |
| categories — lọc bài học theo category | ✅ PASS | |
| lessons — hiển thị danh sách bài học | ✅ PASS | |
| lessons — tìm kiếm bài học theo tên | ✅ PASS | |
| lessons — mở trang chi tiết bài học | ✅ PASS | |
| lessons — phát audio shadowing | ✅ PASS | |
| lessons — lưu trạng thái shadowing completed | ✅ PASS | |
| lessons — luyện dictation và lưu trạng thái | ✅ PASS | |
| profile — hiển thị thông tin tài khoản | ✅ PASS | |
| profile — cập nhật tên hiển thị | ✅ PASS | |
| vocabulary — hiển thị danh sách deck từ vựng | ✅ PASS | |
| vocabulary — tạo deck từ vựng mới | ✅ PASS | |
| vocabulary — thêm từ vựng vào deck | ✅ PASS | |
| vocabulary — chỉnh sửa từ vựng trong deck | ✅ PASS | |
| vocabulary — xoá từ vựng khỏi deck | ✅ PASS | |
| vocabulary — xoá deck từ vựng | ✅ PASS | |

**Tóm tắt User:** 20 PASS / 0 FAIL

### Kết quả Admin Tests (đăng nhập admin)

| Test | Kết quả | Ghi chú |
|---|---|---|
| admin-setup — authenticate as admin | ✅ PASS | Đăng nhập admin thành công |
| categories — hiển thị danh sách category | ✅ PASS | |
| categories — tạo category mới | ✅ PASS | |
| categories — chỉnh sửa category | ✅ PASS | |
| categories — xoá category | ✅ PASS | |
| lessons — hiển thị danh sách bài học | ✅ PASS | |
| lessons — tạo bài học mới | ✅ PASS | |
| lessons — chỉnh sửa bài học | ✅ PASS | |
| lessons — xoá bài học | ✅ PASS | |
| lessons — thêm transcript cho bài học | ✅ PASS | |
| lessons — bulk import transcript | ✅ PASS | |
| users — hiển thị danh sách người dùng | ✅ PASS | |
| users — tìm kiếm người dùng | ✅ PASS | |
| vocabulary — hiển thị system deck | ✅ PASS | |
| vocabulary — tạo system deck mới | ✅ PASS | |
| vocabulary — thêm từ vựng vào system deck | ✅ PASS | |
| vocabulary — chỉnh sửa từ vựng system deck | ✅ PASS | |
| vocabulary — xoá từ vựng khỏi system deck | ✅ PASS | |
| vocabulary — xoá system deck | ✅ PASS | |

**Tóm tắt Admin:** 19 PASS / 0 FAIL

### Screenshots minh hoạ
| File ảnh | Mô tả |
|---|---|
| `screenshots/auth-user.setup.ts-authenticate-as-user-user-setup.png` | Trang đăng nhập user |
| `screenshots/auth-admin.setup.ts-authenticate-as-admin-admin-setup.png` | Trang đăng nhập admin |

### File chi tiết
- `test-report/playwright-e2e/public-output.txt` — Output test public
- `test-report/playwright-e2e/user-admin-output.txt` — Output test user/admin
- `test-report/playwright-e2e/screenshots/` — Ảnh chụp màn hình

---

## 4. TỔNG KẾT

| Loại kiểm thử | Tổng | PASS | FAIL | Trạng thái |
|---|---|---|---|---|
| Go Unit Tests | 46 | 46 | 0 | ✅ Hoàn thành |
| Swagger API Validation | 51 endpoints | 51 | 0 | ✅ Hoàn thành |
| Playwright Public UI | 5 | 5 | 0 | ✅ Hoàn thành |
| Playwright User E2E | 20 | 20 | 0 | ✅ Hoàn thành |
| Playwright Admin E2E | 19 | 19 | 0 | ✅ Hoàn thành |
| **TỔNG CỘNG** | **141** | **141** | **0** | ✅ **TẤT CẢ PASS** |

---

*Báo cáo được tạo ngày 26/05/2026*
