/*
SQL Script to setup the database for the application
Tables:
    - Tasks
    - Lists
    - Tags
    - TaskTags

Relations:
    - Tasks and Lists (1:N)
    - Tasks and Tags (N:M)

TASKS:
    - id: int
    - title: varchar
    - description: text
    - due_date: date
    - completed: boolean
    - list_id: int
    - priority: int
    - created_at: timestamp

LISTS:
    - id: int
    - name: varchar
    - created_at: timestamp

TAGS:
    - id: int
    - name: varchar
    - created_at: timestamp

TASK_TAGS:
    - task_id: int
    - tag_id: int
*/

-- Create the Lists table if it does not exist
CREATE TABLE IF NOT EXISTS Lists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the Tasks table if it does not exist
CREATE TABLE IF NOT EXISTS Tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date DATE,
    completed BOOLEAN DEFAULT FALSE,
    list_id INTEGER,
    priority INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (list_id) REFERENCES Lists(id)
);

-- Create the Tags table if it does not exist
CREATE TABLE IF NOT EXISTS Tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the TaskTags table if it does not exist
CREATE TABLE IF NOT EXISTS TaskTags (
    task_id INTEGER,
    tag_id INTEGER,
    FOREIGN KEY (task_id) REFERENCES Tasks(id),
    FOREIGN KEY (tag_id) REFERENCES Tags(id),
    PRIMARY KEY (task_id, tag_id)
);
