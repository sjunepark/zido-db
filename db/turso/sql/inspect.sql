SELECT * FROM locations;

-- summarize values of locations.UndergroundFlag
SELECT UndergroundFlag, COUNT(*) FROM locations GROUP BY UndergroundFlag;

-- view available tables
SELECT * FROM sqlite_master where type='table';