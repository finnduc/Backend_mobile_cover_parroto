-- 1. Users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    avatar_url VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. Categories
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- 3. Lessons (One Category to Many Lessons)
CREATE TABLE lessons (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT,
    title VARCHAR(255),
    description TEXT,
    video_url VARCHAR(500) NOT NULL,
    thumbnail_url VARCHAR(500),
    level VARCHAR(20),
    duration FLOAT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_lessons_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE SET NULL
);

-- 4. Transcripts
CREATE TABLE transcripts (
    id BIGSERIAL PRIMARY KEY,
    lesson_id BIGINT NOT NULL,
    sequence INTEGER NOT NULL,
    content TEXT NOT NULL,
    phonetic VARCHAR(500),
    vietnamese TEXT,
    start_timestamp FLOAT,
    end_timestamp FLOAT,

    CONSTRAINT fk_transcripts_lesson
        FOREIGN KEY (lesson_id)
        REFERENCES lessons(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_transcripts_sequence
        UNIQUE (lesson_id, sequence)
);

-- 5. Learning History
CREATE TABLE learning_history (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    lesson_id BIGINT NOT NULL,
    duration_watched FLOAT,
    completed BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_history_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_history_lesson
        FOREIGN KEY (lesson_id)
        REFERENCES lessons(id)
        ON DELETE CASCADE
);

-- 6. Bookmarks
CREATE TABLE bookmarks (
    user_id BIGINT NOT NULL,
    lesson_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, lesson_id),

    CONSTRAINT fk_bookmarks_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_bookmarks_lesson
        FOREIGN KEY (lesson_id)
        REFERENCES lessons(id)
        ON DELETE CASCADE
);

-- Indexes để tối ưu truy vấn
CREATE INDEX idx_lessons_category_id ON lessons(category_id);
CREATE INDEX idx_transcripts_lesson_id ON transcripts(lesson_id);
CREATE INDEX idx_learning_history_user ON learning_history(user_id);
CREATE INDEX idx_bookmarks_user ON bookmarks(user_id);
