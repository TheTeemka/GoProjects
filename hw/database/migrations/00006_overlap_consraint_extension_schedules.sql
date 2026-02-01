-- +goose Up
-- +goose StatementBegin
-- enable required index support
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- add generated minute-range columns (stored ints)
ALTER TABLE schedules
  ADD COLUMN start_min int GENERATED ALWAYS AS ((EXTRACT(EPOCH FROM start_time)::int) / 60) STORED,
  ADD COLUMN end_min   int GENERATED ALWAYS AS ((EXTRACT(EPOCH FROM end_time)::int) / 60) STORED;

ALTER TABLE schedules
  ADD CONSTRAINT schedules_no_overlap
  EXCLUDE USING GIST (
    group_id WITH =,
    day_of_week WITH =,
    int4range(start_min, end_min) WITH &&
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE schedules DROP CONSTRAINT IF EXISTS schedules_no_overlap;
ALTER TABLE schedules DROP COLUMN IF EXISTS start_min;
ALTER TABLE schedules DROP COLUMN IF EXISTS end_min;
-- +goose StatementEnd
