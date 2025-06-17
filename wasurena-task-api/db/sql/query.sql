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
-- name: SelectTaskCategories :many
select
	*
from
	task_category cate
where
	cate.owner_user_id = $1
order by
	cate.display_order nulls last
limit 200;
-- name: UpdateTaskCategory :one
update
	task_category
set
	name = $3,
	display_order = $4
where
	id = $1
	and
	owner_user_id = $2
returning *;
-- name: SelectTaskCategoryByID :one
select
	*
from
	task_category cate
where
	cate.owner_user_id = $1
	and
	cate.id = $2;
-- name: DeleteTaskCategory :one
delete
from
	task_category cate
where
	cate.id = $1
	and
	cate.owner_user_id = $2
returning *;
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
-- name: UpdateTaskDefinition :one
update
	task_definition
set
	title = $3,
	display_flag = $4,
	notification_flag = $5,
	dead_line_check = $6,
	dead_line_check_sub_setting = $7,
	category_id = $8,
	detail = $9
where
	id = $1
	and
	owner_user_id = $2
 returning *;
-- name: UpdateAllTaskNotificationFlagByUser :many
update
	task_definition
set
	notification_flag = $2
where
	owner_user_id = $1
returning *;
-- name: UpdateAllTaskCategoryNull :many
update
	task_definition
set
	category_id = null
where
	category_id = $1
	and
	owner_user_id = $2
returning *;
-- name: SelectLatestTaskExecuteForNotify :many
select
	def.*,
	(case
		when exec.execute_date_time is null then '1999-12-31 15:00:00+00'::timestamptz
		else exec.execute_date_time::timestamptz
	end) as latest_exec_date_time
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
		def.owner_user_id,
		def.id;
-- name: SelectTaskDefinitionList :many
select
	def.*,
	task_category.name as category_name
from
		task_definition def
left outer join task_category on
		task_category.id = def.category_id
where
	def.owner_user_id = $1
order by
		task_category.display_order nulls last,
		def.id desc
limit 300;
-- name: SelectTaskDefinitionById :one
select
	def.*,
	task_category.name as category_name
from
	task_definition def
left outer join task_category on
	task_category.id = def.category_id
where
	def.owner_user_id = $1
	and
	def.id = $2;
-- name: SelectTaskCheckDisplayList :many
select
	def.*,
	(case
		when exec.execute_date_time is null then '1999-12-31 15:00:00+00'::timestamptz
		else exec.execute_date_time::timestamptz
	end) as latest_exec_date_time,
	cate.name as category_name
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
left outer join task_category cate on
		cate.id = def.category_id
where
		def.display_flag = true
	and
		def.owner_user_id = $1
order by
	case
		when def.dead_line_check is null then 1
		else 0
	end,
		cate.display_order nulls last,
		exec.execute_date_time;
-- name: DeleteTaskDefinition :one
delete
from
	task_definition def
where
	def.id = $1
	and
	def.owner_user_id = $2
returning *;
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
-- name: DeleteTaskExecuteByDefinitionId :many
delete
from
	task_execute exe
		using
	 task_definition def
where
	def.id = exe.task_definition_id
	and
	exe.task_definition_id = $1
	and
	def.owner_user_id = $2
returning *;
-- name: SelectTaskExecuteByDefinitionId :many
select
	*
from
	task_execute exec
where
	exec.execute_user_id = $1
	and exec.task_definition_id = $2
order by
	exec.execute_date_time desc;
-- name: DeleteTaskExecuteById :one
delete
from
	task_execute exe
where
	exe.id = $1
	and
	exe.execute_user_id = $2
returning *;
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
-- name: UpdateUserAccountLineBotFollow :one
update
	user_accounts
set
	is_line_bot_follow = $2
where
	id = $1
returning *;
-- name: UpdateUserImageUrl :one
update
	user_accounts
set
	image_url = $2
where
	id = $1
returning *;
-- name: UpdateUserAccountInfo :one
update
	user_accounts
set
	user_setting_id = $2,
	user_name = $3
where
	id = $1
returning *;
