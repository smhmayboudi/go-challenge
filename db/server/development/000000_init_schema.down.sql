CREATE TABLE Instrument('Id' int, 'Name' varchar(255));

CREATE TABLE Trade('Id' int, 'InstrumentId' int, 'DateEn' timestamptz, 'Open' decimal, 'High' decimal, 'Low' decimal, 'Close' decimal);

INSERT INTO Instrument values(1,'AAPL'), (2,'GOOGL');

INSERT INTO Trade VALUES
(401 ,301 ,2001 ,1001 ,'2020-01-01' ,1 ,1),
(402 ,302 ,2002 ,1002 ,'2020-01-02' ,1 ,2),
(403 ,303 ,2003 ,1003 ,'2020-01-03' ,1 ,3),
(404 ,304 ,2004 ,1004 ,'2020-01-01' ,2 ,4),
(405 ,305 ,2005 ,1005 ,'2020-01-03' ,2 ,5),
(406 ,306 ,2006 ,1006 ,'2020-01-01' ,5 ,6),
(407 ,307 ,2007 ,1007 ,'2021-01-01' ,1 ,7);

-- QUERY
-- SELECT 'Instrument'.'Name', 'Trade'.'DateEn', 'Trade'.'Open', 'Trade'.'High', 'Trade'.'Low', 'Trade'.'Close' FROM Instrument RIGHT JOIN Trade ON 'Instrument'.'Id' = 'Trade'.'InstrumentId'
