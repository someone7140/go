create table task_category (
  id varchar(255) primary key,
  name varchar(255) not null,
  owner_user_id varchar(255) not null,
  display_order integer
);

create index on
task_category (owner_user_id);

create type dead_line_check_enum as enum (
  'DailyOnce',
  'DailyHour',
  'WeeklyDay',
  'WeeklyDayInterval',
  'MonthOnce',
  'MonthDate',
  'YearOnceDate'
  );

create table task_definition (
  id varchar(255) primary key,
  title varchar(255) not null,
  owner_user_id varchar(255) not null,
  display_flag boolean not null,
  notification_flag boolean not null,
  category_id varchar(255),
  dead_line_check dead_line_check_enum,
  dead_line_check_sub_setting jsonb,
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

create table task_execute (
  id varchar(255) primary key,
  task_definition_id varchar(255) not null references task_definition(id),
  execute_user_id varchar(255) not null,
  execute_date_time timestamptz not null,
  memo text
);

create index on
task_execute (task_definition_id);

create index on
task_execute (execute_user_id);
