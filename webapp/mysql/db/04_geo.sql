ALTER TABLE isuumo.estate ADD geo GEOMETRY;
UPDATE isuumo.estate SET estate.`geo` = GeomFromText(Concat('POINT(',estate.latitude,' ',estate.longitude,')'));
ALTER TABLE isuumo.estate MODIFY geo GEOMETRY NOT NULL;
CREATE SPATIAL INDEX estate_door ON isuumo.estate (geo);
