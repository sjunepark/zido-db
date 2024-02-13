-- name: GetLocations :one
SELECT *
FROM locations;

-- name: InsertLocation :execresult
INSERT INTO locations (BJDNumber, SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber,
                       BuildingSubNumber, SDName, SGGName, EMDName, RoadName, BuildingName, PostalNumber, Long, Lat,
                       Crs, X, Y, ValidPosition, BaseDate, DatetimeAdded)
VALUES ($BJDNumber, $SGGNumber, $EMDNumber, $RoadNumber, $UndergroundFlag, $BuildingMainNumber,
        $BuildingSubNumber, $SDName, $SGGName, $EMDName, $RoadName, $BuildingName, $PostalNumber, $Long, $Lat,
        $Crs, $X, $Y, $ValidPosition, $BaseDate, $DatetimeAdded);
