
-- +migrate Up
CREATE TABLE public.vehicles (
	id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	created_by text NULL,
	updated_by text NULL,
	"name" text NULL,
	model text NULL,
	count int8 NULL,
	CONSTRAINT vehicles_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_public_vehicles_deleted_at ON public.vehicles USING btree (deleted_at);

-- +migrate Down
DROP TABLE public.vehicles;
