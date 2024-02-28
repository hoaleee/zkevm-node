-- +migrate Up
CREATE TABLE IF NOT EXISTS state.blobs
(
    batch_num BIGINT PRIMARY KEY,
    blob_id BYTEA
);

-- +migrate Down

DROP TABLE IF EXISTS state.blobs;