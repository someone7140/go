create table task_definition (
  id varchar(255) primary key,
  title varchar(255) not null,
  owner_user_id varchar(255) not null,
  display_flag boolean not null,
  notification_flag boolean not null,
  category_id varchar(255),
  detail text
);

create index on
task_definition (title);

create index on
task_definition (owner_user_id);

create index on
task_definition (display_flag);

create index on
task_definition (notification_flag);

create index on
task_definition (category_id);
