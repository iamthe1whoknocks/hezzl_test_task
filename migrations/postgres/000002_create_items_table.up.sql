CREATE TABLE IF NOT EXISTS items(
    id serial PRIMARY KEY,
    campaign_id int
     constraint items_campaign_id_fk
            references campaigns(id)
            on delete no action,
    name VARCHAR(255),
    description VARCHAR(255),
    priority serial,
    removed bool,
    created_at timestamp
);

drop index if exists items_id_idx;
create unique index if not exists items_id_idx on items (id);

drop index if exists items_campaign_id_idx;
create  index if not exists items_campaign_id_idx on items (campaign_id);

drop index if exists items_name_idx;
create  index if not exists items_name_idx on items (name);