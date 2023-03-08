CREATE TABLE IF NOT EXISTS campaigns(
    id serial PRIMARY KEY,
    name VARCHAR(255)
);

drop index if exists campaigns_id_idx;
create unique index if not exists campaigns_id_idx on campaigns (id);

insert into campaigns (id,name) values (1,'Первая запись')