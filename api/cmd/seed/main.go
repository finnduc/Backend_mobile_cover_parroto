package main

// cmd/seed/main.go
// Run: go run ./cmd/seed

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/database/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback for local dev
		dsn = "host=localhost user=postgres password=postgres dbname=parroto port=5433 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("🌱 Starting seed...")

	seedCategories(db)
	seedUsers(db)
	seedLessons(db)
	seedTranscripts(db)
	seedBookmarks(db)
	fmt.Println("✅ Seed completed!")
}

// ─── Categories ──────────────────────────────────────────────

func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{ID: 1, Name: "Daily Conversation"},
		{ID: 2, Name: "Business English"},
		{ID: 3, Name: "Travel & Tourism"},
		{ID: 4, Name: "News & Media"},
	}
	result := db.CreateInBatches(&categories, len(categories))
	if result.Error != nil {
		log.Printf("⚠️  categories: %v", result.Error)
	} else {
		fmt.Printf("   categories: %d rows\n", result.RowsAffected)
	}
}

// ─── Users ────────────────────────────────────────────────────

func hashPassword(plain string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("bcrypt error: %v", err)
	}
	return string(b)
}

func seedUsers(db *gorm.DB) {
	hash := hashPassword("Password123!")

	users := []models.User{
		{
			ID:       "1",
			Email:    "admin@parroto.com",
			Name:     "Admin",
			UserRole: enums.UserRole("admin"),
			Password: hash,
		},
		{
			ID:        "2",
			Email:     "alice@example.com",
			Name:      "Alice Nguyen",
			UserRole:  enums.UserRole("user"),
			Password:  hash,
			AvatarURL: "https://api.dicebear.com/7.x/thumbs/svg?seed=alice",
			CreatedAt: time.Now(),
		},
		{
			ID:        "3",
			Email:     "bob@example.com",
			Name:      "Bob Tran",
			UserRole:  enums.UserRole("user"),
			Password:  hash,
			AvatarURL: "https://api.dicebear.com/7.x/thumbs/svg?seed=bob",
			CreatedAt: time.Now(),
		},
		{
			ID:        "4",
			Email:     "carol@example.com",
			Name:      "Carol Le",
			UserRole:  enums.UserRole("user"),
			Password:  hash,
			AvatarURL: "https://api.dicebear.com/7.x/thumbs/svg?seed=carol",
			CreatedAt: time.Now(),
		},
		{
			ID:        "5",
			Email:     "david@example.com",
			Name:      "David Pham",
			UserRole:  enums.UserRole("user"),
			Password:  hash,
			AvatarURL: "https://api.dicebear.com/7.x/thumbs/svg?seed=david",
			CreatedAt: time.Now(),
		},
	}

	result := db.CreateInBatches(&users, len(users))
	if result.Error != nil {
		log.Printf("⚠️  users: %v", result.Error)
	} else {
		fmt.Printf("   users: %d rows\n", result.RowsAffected)
	}
}

// ─── Lessons ──────────────────────────────────────────────────

func ptr(u uint) *uint { return &u }

func seedLessons(db *gorm.DB) {
	now := time.Now()
	lessons := []models.Lesson{
		{ID: 1, CategoryID: ptr(1), Title: "Greeting People",
			Description:  "Learn how to greet people naturally in everyday situations.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "beginner", Duration: 180.0, CreatedAt: now},
		{ID: 2, CategoryID: ptr(1), Title: "Ordering Food at a Restaurant",
			Description:  "Practice ordering food and drinks confidently at a restaurant.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "beginner", Duration: 240.0, CreatedAt: now},
		{ID: 3, CategoryID: ptr(2), Title: "Job Interview Essentials",
			Description:  "Key phrases and answers for common job interview questions.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "intermediate", Duration: 360.0, CreatedAt: now},
		{ID: 4, CategoryID: ptr(3), Title: "Asking for Directions",
			Description:  "How to ask for and understand directions when traveling.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "beginner", Duration: 200.0, CreatedAt: now},
		{ID: 5, CategoryID: ptr(4), Title: "Talking About Current Events",
			Description:  "Discuss news and current events using natural English expressions.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "advanced", Duration: 420.0, CreatedAt: now},
		{ID: 6, CategoryID: ptr(2), Title: "Email Writing for Professionals",
			Description:  "Write clear and professional emails in English.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "intermediate", Duration: 300.0, CreatedAt: now},
		{ID: 7, CategoryID: ptr(1), Title: "Small Talk at Work",
			Description:  "Build rapport with colleagues through casual conversation.",
			VideoURL:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			ThumbnailURL: "https://img.youtube.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
			Level:        "beginner", Duration: 220.0, CreatedAt: now},
	}

	result := db.CreateInBatches(&lessons, len(lessons))
	if result.Error != nil {
		log.Printf("⚠️  lessons: %v", result.Error)
	} else {
		fmt.Printf("   lessons: %d rows\n", result.RowsAffected)
	}
}

// ─── Transcripts ──────────────────────────────────────────────

func seedTranscripts(db *gorm.DB) {
	transcripts := []models.Transcript{
		// Lesson 1 — Greeting People
		{ID: 1, LessonID: 1, Sequence: 1, Content: "Hey, how are you doing?", Phonetic: "/heɪ haʊ ɑːr juː ˈduːɪŋ/", Vietnamese: "Này, bạn có khỏe không?", StartTimestamp: 0.0, EndTimestamp: 3.5},
		{ID: 2, LessonID: 1, Sequence: 2, Content: "I'm doing great, thanks for asking.", Phonetic: "/aɪm ˈduːɪŋ ɡreɪt θæŋks fər ˈæskɪŋ/", Vietnamese: "Tôi đang rất tốt, cảm ơn bạn đã hỏi.", StartTimestamp: 3.5, EndTimestamp: 7.0},
		{ID: 3, LessonID: 1, Sequence: 3, Content: "It's nice to meet you.", Phonetic: "/ɪts naɪs tə miːt juː/", Vietnamese: "Rất vui được gặp bạn.", StartTimestamp: 7.0, EndTimestamp: 9.5},
		{ID: 4, LessonID: 1, Sequence: 4, Content: "Likewise! What brings you here today?", Phonetic: "/ˈlaɪkwaɪz wɒt brɪŋz juː hɪər təˈdeɪ/", Vietnamese: "Tôi cũng vậy! Hôm nay bạn đến đây có việc gì?", StartTimestamp: 9.5, EndTimestamp: 13.0},
		{ID: 5, LessonID: 1, Sequence: 5, Content: "I just moved to the neighborhood.", Phonetic: "/aɪ dʒʌst muːvd tə ðə ˈneɪbəhʊd/", Vietnamese: "Tôi vừa chuyển đến khu phố này.", StartTimestamp: 13.0, EndTimestamp: 16.5},

		// Lesson 2 — Ordering Food
		{ID: 6, LessonID: 2, Sequence: 1, Content: "Excuse me, could I see the menu please?", Phonetic: "/ɪkˈskjuːz miː kʊd aɪ siː ðə ˈmenjuː pliːz/", Vietnamese: "Xin lỗi, cho tôi xem thực đơn được không?", StartTimestamp: 0.0, EndTimestamp: 4.0},
		{ID: 7, LessonID: 2, Sequence: 2, Content: "Of course! Here you go.", Phonetic: "/əv kɔːrs hɪər juː ɡoʊ/", Vietnamese: "Tất nhiên! Đây ạ.", StartTimestamp: 4.0, EndTimestamp: 6.0},
		{ID: 8, LessonID: 2, Sequence: 3, Content: "I'd like to order the grilled salmon.", Phonetic: "/aɪd laɪk tə ˈɔːrdər ðə ɡrɪld ˈsæmən/", Vietnamese: "Tôi muốn gọi cá hồi nướng.", StartTimestamp: 6.0, EndTimestamp: 9.5},
		{ID: 9, LessonID: 2, Sequence: 4, Content: "And what would you like to drink?", Phonetic: "/ænd wɒt wʊd juː laɪk tə drɪŋk/", Vietnamese: "Và bạn muốn uống gì?", StartTimestamp: 9.5, EndTimestamp: 12.0},
		{ID: 10, LessonID: 2, Sequence: 5, Content: "Just water for now, thank you.", Phonetic: "/dʒʌst ˈwɔːtər fər naʊ θæŋk juː/", Vietnamese: "Tạm thời chỉ nước thôi, cảm ơn.", StartTimestamp: 12.0, EndTimestamp: 15.0},

		// Lesson 3 — Job Interview
		{ID: 11, LessonID: 3, Sequence: 1, Content: "Tell me about yourself.", Phonetic: "/tel miː əˈbaʊt jɔːrˈself/", Vietnamese: "Hãy giới thiệu về bản thân bạn.", StartTimestamp: 0.0, EndTimestamp: 2.5},
		{ID: 12, LessonID: 3, Sequence: 2, Content: "I have five years of experience in software development.", Phonetic: "/aɪ hæv faɪv jɪərz əv ɪkˈspɪəriəns ɪn ˈsɒftweər dɪˈveləpmənt/", Vietnamese: "Tôi có năm năm kinh nghiệm trong lĩnh vực phát triển phần mềm.", StartTimestamp: 2.5, EndTimestamp: 8.0},
		{ID: 13, LessonID: 3, Sequence: 3, Content: "What are your greatest strengths?", Phonetic: "/wɒt ɑːr jɔːr ˈɡreɪtɪst streŋθs/", Vietnamese: "Điểm mạnh lớn nhất của bạn là gì?", StartTimestamp: 8.0, EndTimestamp: 11.0},
		{ID: 14, LessonID: 3, Sequence: 4, Content: "I'm a fast learner and a great team player.", Phonetic: "/aɪm ə fæst ˈlɜːrnər ænd ə ɡreɪt tiːm ˈpleɪər/", Vietnamese: "Tôi học nhanh và là một người đồng đội xuất sắc.", StartTimestamp: 11.0, EndTimestamp: 15.5},
		{ID: 15, LessonID: 3, Sequence: 5, Content: "Where do you see yourself in five years?", Phonetic: "/weər duː juː siː jɔːrˈself ɪn faɪv jɪərz/", Vietnamese: "Bạn thấy mình ở đâu trong năm năm tới?", StartTimestamp: 15.5, EndTimestamp: 19.0},
	}

	result := db.CreateInBatches(&transcripts, len(transcripts))
	if result.Error != nil {
		log.Printf("⚠️  transcripts: %v", result.Error)
	} else {
		fmt.Printf("   transcripts: %d rows\n", result.RowsAffected)
	}
}

// ─── Transcript Bookmarks ─────────────────────────────────────

func seedBookmarks(db *gorm.DB) {
	bookmarks := []models.TranscriptBookmark{
		{UserID: "2", TranscriptID: 1, Note: "", CreatedAt: time.Now()},
		{UserID: "2", TranscriptID: 3, Note: "important", CreatedAt: time.Now()},
		{UserID: "3", TranscriptID: 2, Note: "", CreatedAt: time.Now()},
		{UserID: "4", TranscriptID: 1, Note: "review later", CreatedAt: time.Now()},
		{UserID: "4", TranscriptID: 5, Note: "", CreatedAt: time.Now()},
		{UserID: "5", TranscriptID: 3, Note: "", CreatedAt: time.Now()},
	}

	result := db.CreateInBatches(&bookmarks, len(bookmarks))
	if result.Error != nil {
		log.Printf("⚠️  transcript bookmarks: %v", result.Error)
	} else {
		fmt.Printf("   transcript bookmarks: %d rows\n", result.RowsAffected)
	}
}


