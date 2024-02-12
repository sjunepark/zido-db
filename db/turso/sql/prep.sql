CREATE TABLE IF NOT EXISTS locations (
                                         Id                 TEXT PRIMARY KEY,
                                         BJDNumber          TEXT NOT NULL,
                                         SGGNumber          TEXT NOT NULL,
                                         EMDNumber          TEXT NOT NULL,
                                         RoadNumber         TEXT NOT NULL,
                                         UndergroundFlag    TEXT,
                                         BuildingMainNumber TEXT NOT NULL,
                                         BuildingSubNumber  TEXT,
                                         SDName             TEXT NOT NULL,
                                         SGGName            TEXT,
                                         EMDName            TEXT NOT NULL,
                                         RoadName           TEXT NOT NULL,
                                         BuildingName       TEXT,
                                         PostalNumber       TEXT NOT NULL,
                                         Long               REAL,
                                         Lat                REAL,
                                         Crs                TEXT NOT NULL,
                                         X                  REAL,
                                         Y                  REAL,
                                         ValidPosition      BOOLEAN,
                                         BaseDate           DATETIME NOT NULL,
                                         DatetimeAdded      DATETIME NOT NULL
);

-- DROP TABLE IF EXISTS locations;