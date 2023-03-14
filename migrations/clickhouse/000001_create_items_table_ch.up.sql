CREATE TABLE IF NOT EXISTS items(
    id UInt32,
    campaign_id UInt32,
    name String,
    description String,
    priority UInt32,
    removed Bool,
    event_time DateTime
) 
ENGINE=MergeTree()
PRIMARY KEY (id,event_time);


