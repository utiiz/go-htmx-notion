CREATE TABLE IF NOT EXISTS "FAVORITE" (
    user_id INT,
    project_id INT,
    PRIMARY KEY (user_id, project_id),
    FOREIGN KEY (user_id) REFERENCES "USER"(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES "PROJECT"(id) ON DELETE CASCADE
);
