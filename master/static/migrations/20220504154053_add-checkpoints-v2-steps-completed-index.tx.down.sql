-- This change wasn't caught on a version boundary; if a version downgrade is run, the down
-- migration prior to us will also run, and it runs a delete cascade that will drop all our stuff.