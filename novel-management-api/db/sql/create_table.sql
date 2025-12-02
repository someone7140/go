CREATE TABLE user_accounts (
	id varchar NOT NULL,
	user_setting_id varchar NOT NULL,
	"name" varchar NOT NULL,
	gmail varchar NOT NULL,
	image_url varchar,
	created_at timestamptz NOT NULL,
	CONSTRAINT user_accounts_gmail_key UNIQUE (gmail),
	CONSTRAINT user_accounts_pkey PRIMARY KEY (id),
	CONSTRAINT user_accounts_user_setting_id_key UNIQUE (user_setting_id)
);

CREATE TABLE novels (
	id varchar NOT NULL,
	title varchar NOT NULL,
	description text,
	owner_user_account_id varchar NOT NULL,
	created_at timestamptz NOT NULL,
	CONSTRAINT novels_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_novels_owner_user_account_id ON novels USING btree (owner_user_account_id);

CREATE TABLE novel_settings (
	id varchar NOT NULL,
	"name" varchar NOT NULL,
	novel_id varchar NOT NULL,
	owner_user_account_id varchar NOT NULL,
	parent_setting_id varchar,
	display_order int4,
	"attributes" varchar[] NOT NULL,
	description text,
	CONSTRAINT novel_settings_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_novel_settings_novel_id ON novel_settings USING btree (novel_id);
CREATE INDEX idx_novel_settings_owner_user_account_id ON novel_settings USING btree (owner_user_account_id);

CREATE TABLE novel_contents (
	id varchar NOT NULL,
	chapter_name varchar NOT NULL,
	novel_id varchar NOT NULL,
	owner_user_account_id varchar NOT NULL,
	parent_contents_id varchar,
	display_order int4,
	contents text,
	description text,
	CONSTRAINT novel_contents_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_novel_contents_novel_id ON novel_contents USING btree (novel_id);
CREATE INDEX idx_novel_contents_owner_user_account_id ON novel_contents USING btree (owner_user_account_id);
