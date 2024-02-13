CREATE TABLE IF NOT EXISTS locations
(
    BJDNumber          TEXT     NOT NULL CHECK (length(BJDNumber) = 10),
    SGGNumber          TEXT     NOT NULL CHECK (length(SGGNumber) = 5),
    EMDNumber          TEXT     NOT NULL CHECK (length(EMDNumber) = 3),
    RoadNumber         TEXT     NOT NULL CHECK (length(RoadNumber) = 7),
    UndergroundFlag    INTEGER  NOT NULL CHECK (UndergroundFlag = 0 OR UndergroundFlag = 1 OR UndergroundFlag = 2),
    BuildingMainNumber INTEGER  NOT NULL CHECK (BuildingMainNumber >= 0 AND BuildingMainNumber <= 99999),
    BuildingSubNumber  INTEGER  NOT NULL CHECK (BuildingSubNumber >= 0 AND BuildingSubNumber <= 99999),
    SDName             TEXT     NOT NULL CHECK (length(SDName) <= 40),
    SGGName            TEXT CHECK (length(SGGName) <= 40),
    EMDName            TEXT     NOT NULL CHECK (length(EMDName) <= 40),
    RoadName           TEXT     NOT NULL CHECK (length(RoadName) <= 40),
    BuildingName       TEXT CHECK (length(BuildingName) <= 40),
    PostalNumber       TEXT     NOT NULL CHECK (length(PostalNumber) = 5),
    Long               REAL CHECK (Long >= -180 AND Long <= 180),
    Lat                REAL CHECK (Lat >= -90 AND Lat <= 90),
    Crs                TEXT     NOT NULL,
    X                  REAL,
    Y                  REAL,
    ValidPosition      BOOLEAN CHECK (ValidPosition = 0 OR ValidPosition = 1),
    BaseDate           DATETIME NOT NULL,
    DatetimeAdded      DATETIME NOT NULL,
    PRIMARY KEY (SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber)
);