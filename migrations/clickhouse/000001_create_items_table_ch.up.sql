CREATE TABLE IF NOT EXISTS items(
    id UInt32,
    campaign_id UInt32
    name string
    description string,
    priority UInt32,
    removed bool,
    event_time DateTime
) Engine=MergeTree;

DROP INDEX IF EXISTS items_id_idx;
ALTER TABLE items ADD INDEX items_id_idx(id) TYPE minmax;

DROP INDEX IF EXISTS items_campaign_id_idx;
ALTER TABLE items ADD INDEX items_campaign_id_idx(campaign_id) TYPE minmax;

DROP INDEX IF EXISTS items_name_idx;
ALTER TABLE items ADD INDEX items_name_idx(name) TYPE minmax;