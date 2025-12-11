PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  display_name TEXT,
  last_seen INTEGER
);

CREATE TABLE IF NOT EXISTS positions (
  user_id TEXT,
  lat REAL,
  lon REAL,
  alt REAL,
  heading REAL,
  speed REAL,
  timestamp INTEGER,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS markers (
  id TEXT PRIMARY KEY,
  owner_id TEXT,
  lat REAL,
  lon REAL,
  type TEXT,
  label TEXT,
  created_at INTEGER,
  updated_at INTEGER,
  FOREIGN KEY(owner_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS tile_versions (
  zoom INTEGER,
  x INTEGER,
  y INTEGER,
  version INTEGER,
  updated_at INTEGER,
  PRIMARY KEY(zoom, x, y)
);
