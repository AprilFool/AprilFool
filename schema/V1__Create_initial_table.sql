-- Group
CREATE TABLE IF NOT EXISTS "Group" (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(50),
    description     TEXT
);

-- User
CREATE TABLE IF NOT EXISTS "User" (
    id          SERIAL PRIMARY KEY,
    created     TIMESTAMP WITHOUT TIME ZONE,
    mail        VARCHAR(50),
    password    VARCHAR(60),
    name        VARCHAR(50),
    photo       VARCHAR(20),
    "group"     INTEGER REFERENCES "Group" (id),
    status      INTEGER
);


-- Problem
CREATE TABLE IF NOT EXISTS Problem (
    id              SERIAL PRIMARY KEY,
    created         TIMESTAMP WITHOUT TIME ZONE,
    modified        TIMESTAMP WITHOUT TIME ZONE,
    "user"          INTEGER REFERENCES "User" (id),
    title           VARCHAR(100),
    content         TEXT,
    level           INTEGER,
    constraints     TEXT
);

CREATE TABLE IF NOT EXISTS Problem_testcase (
    id          SERIAL PRIMARY KEY,
    problem     INTEGER REFERENCES Problem (id),
    level       INTEGER,
    testcase    TEXT
);

-- Tag
CREATE TABLE IF NOT EXISTS Tag (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(50)
);

-- Problem-Tag
CREATE TABLE IF NOT EXISTS Problem_Tag (
    problem     INTEGER REFERENCES Problem (id),
    tag         INTEGER REFERENCES Tag (id),
    CONSTRAINT unique__problem_tag UNIQUE (problem, tag)
);

-- Language
CREATE TABLE IF NOT EXISTS Language (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(30),
    description     VARCHAR(100),
    build           VARCHAR(200),
    run             VARCHAR(200)
);
    
-- Submission
CREATE TABLE IF NOT EXISTS "Submission" (
    id              SERIAL PRIMARY KEY,
    created         TIMESTAMP WITHOUT TIME ZONE,
    "user"          INTEGER REFERENCES "User" (id),
    problem         INTEGER REFERENCES Problem (id),
    language        INTEGER REFERENCES Language (id),
    code            TEXT,
    time            INTEGER,        -- seconds
    status          INTEGER,        -- wait, run, error, pass, fail, constraint
    description     TEXT            -- compiler message, ...
);
