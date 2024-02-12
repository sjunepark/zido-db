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

-- BJDNumber = {string} '3611010100'
-- SGGNumber = {string} '36110'
-- EMDNumber = {string} '101'
-- RoadNumber = {string} '2000002'
-- UndergroundFlag = {int} 0
-- BuildingMainNumber = {int} 1811
-- BuildingSubNumber = {int} 0
-- SDName = {string} '세종특별자치시'
-- SGGName = {string} ''
-- EMDName = {string} '반곡동'
-- RoadName = {string} '한누리대로'
-- BuildingName = {string} '수루배마을5단지 상가동'
-- PostalNumber = {string} '30145'
-- Long = {float64} 127.31348634049063
-- Lat = {float64} 36.4974913911321
-- Crs = {string} 'EPSG:5179'
-- X = {float64} 983296.172464
-- Y = {float64} 1833330.968984
-- ValidPosition = {int} 1
-- BaseDate = {time.Time} 2023-12-31T15:00:00Z
-- DatetimeAdded = {time.Time} 2024-02-13 00:51:35.589909 +0900 m=+0.005833793

INSERT INTO locations (BJDNumber,
                       SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber, SDName,
                       SGGName, EMDName, RoadName, BuildingName, PostalNumber, Long, Lat, Crs, X, Y, ValidPosition,
                       BaseDate, DatetimeAdded)
VALUES ('3611010100', '36110', '101', '2000002', 0, 1811, 0, '세종특별자치시', '', '반곡동', '한누리대로', '수루배마을5단지 상가동', '30145',
        127.31348634049063, 36.4974913911321, 'EPSG:5179', 983296.172464, 1833330.968984, 1, '2023-12-31T15:00:00Z',
        '2024-02-13 00:51:35.589909 +0900 m=+0.005833793');


SELECT sqlite_version();

DROP TABLE IF EXISTS locations;

-- Same table without CHECK constraints
CREATE TABLE IF NOT EXISTS locations
(
    BJDNumber          TEXT     NOT NULL,
    SGGNumber          TEXT     NOT NULL,
    EMDNumber          TEXT     NOT NULL,
    RoadNumber         TEXT     NOT NULL,
    UndergroundFlag    INTEGER  NOT NULL,
    BuildingMainNumber INTEGER  NOT NULL,
    BuildingSubNumber  INTEGER  NOT NULL,
    SDName             TEXT     NOT NULL,
    SGGName            TEXT,
    EMDName            TEXT     NOT NULL,
    RoadName           TEXT     NOT NULL,
    BuildingName       TEXT,
    PostalNumber       TEXT     NOT NULL,
    Long               REAL,
    Lat                REAL,
    Crs                TEXT     NOT NULL,
    X                  REAL,
    Y                  REAL,
    ValidPosition      BOOLEAN,
    BaseDate           DATETIME NOT NULL,
    DatetimeAdded      DATETIME NOT NULL,
    PRIMARY KEY (SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber)
);
