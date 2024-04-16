CREATE TABLE IF NOT EXISTS Job (
"id" TEXT NOT NULL,
"source" TEXT,
"name" TEXT,
"type" INT,
"source_url" TEXT,
"hardware" TEXT,
"deploy_status" TEXT,
"status" TEXT,
"result_url" TEXT,
"container_log" TEXT,
"name_space" TEXT,
"k8s_deploy_name" TEXT,
"k8s_resource_type" TEXT,
"image_name" TEXT,
"build_log" TEXT,
"expire_time" INT,
"create_at" INT,
PRIMARY KEY ("id")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS index_job_id on Job(id);

CREATE TABLE IF NOT EXISTS Task (
"id" INT NOT NULL,
"zk_type" TEXT,
"name" TEXT,
"resource_type" INT,
"input_param" TEXT,
"tx_hash" TEXT,
"status" INT,
"reward" TEXT,
"create_time" INT,
"error" TEXT,
PRIMARY KEY ("id")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS index_task_id on Job(id);