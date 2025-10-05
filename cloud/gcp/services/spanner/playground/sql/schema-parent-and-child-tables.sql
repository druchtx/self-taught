-- example 1

CREATE TABLE Singers (
 SingerId   INT64 NOT NULL PRIMARY KEY,
 FirstName  STRING(1024),
 LastName   STRING(1024),
 SingerInfo BYTES(MAX),
 );

CREATE TABLE Albums (
 SingerId     INT64 NOT NULL,
 AlbumId      INT64 NOT NULL,
 AlbumTitle   STRING(MAX),
 ) PRIMARY KEY (SingerId, AlbumId),
-- `ON DELETE CASCADE` will cause the entire row to be deleted when the parent row is deleted
-- `ON DELETE NO ACTION` will take no action when the parent row is deleted
INTERLEAVE IN PARENT Singers ON DELETE CASCADE;

CREATE TABLE Songs (
 SingerId     INT64 NOT NULL,
 AlbumId      INT64 NOT NULL,
 TrackId      INT64 NOT NULL,
 SongName     STRING(MAX),
) PRIMARY KEY (SingerId, AlbumId, TrackId),
 INTERLEAVE IN PARENT Albums ON DELETE CASCADE;

-- example 2

CREATE TABLE Projects (
  ProjectId   INT64 NOT NULL,
  ProjectName STRING(1024),
) PRIMARY KEY (ProjectId);

CREATE TABLE Resources (
  ProjectId    INT64 NOT NULL,
  ResourceId   INT64 NOT NULL,
  ResourceName STRING(1024),
) PRIMARY KEY (ProjectId, ResourceId),
-- Resources can be inserted even if the project does not exist
  INTERLEAVE IN Projects;