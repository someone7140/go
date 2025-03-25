-- name: CreateTaskCategory :one
insert
	into
	task_category (
    id,
	name,
	owner_user_id,
	display_order
)
values (
    $1, 
    $2, 
    $3, 
    $4
) returning *;
-- name: CreateTaskDefinition :one
insert
	into
	task_definition (
    id,
	title,
	owner_user_id,
	display_flag,
	notification_flag,
	dead_line_check,
	dead_line_check_sub_setting,
	category_id,
	detail
)
values (
    $1, 
    $2, 
    $3, 
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) returning *;
-- name: SelectLatestTaskExecuteForNotify :many
select
	def.*,
	(case
		when exec.execute_date_time is null then '1999-12-31 15:00:00+00'::timestamptz
		else exec.execute_date_time::timestamptz
	end) as latest_date_time
from
		task_definition def
left outer join 
	(
	select
			task_definition_id,
			max(execute_date_time) as execute_date_time
	from
			task_execute
	group by
			task_definition_id) exec on
		def.id = exec.task_definition_id
where
		def.notification_flag = true
	and def.dead_line_check is not null
order by
		owner_user_id,
		id;
-- name: CreateTaskExecute :one
insert
	into
	task_execute (
    id,
	task_definition_id,
	execute_user_id,
	execute_date_time,
	memo
)
values (
    $1, 
    $2, 
    $3, 
    $4,
    $5
) returning *;
-- name: SelectUserAccountByUserSettingId :one
select
	*
from
	user_accounts
where
	user_setting_id = $1;
-- name: SelectUserAccountByLineId :one
select
	*
from
	user_accounts
where
	line_id = $1;
-- name: SelectUserAccountById :one
select
	*
from
	user_accounts
where
	id = $1;
-- name: CreateUserAccount :one
insert
	into
	user_accounts (
    id,
	user_setting_id,
	line_id,
	user_name,
	image_url
)
values (
    $1, 
    $2, 
    $3, 
    $4,
    $5
) returning *;
