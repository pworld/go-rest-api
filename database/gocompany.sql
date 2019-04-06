-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
"id" serial4 NOT NULL,
"username" varchar(64) COLLATE "default",
"email" varchar(64) COLLATE "default" NOT NULL,
"password" varchar(100) COLLATE "default" NOT NULL,
"verify_at" timestamp(6),
"last_login_at" timestamp(6)
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD UNIQUE ("email");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Table structure for employee_status
-- ----------------------------
DROP TABLE IF EXISTS "public"."employee_status";
CREATE TABLE "public"."employee_status" (
"id" serial4 NOT NULL,
"name" varchar(100) COLLATE "default" NOT NULL,
"description" varchar(255) COLLATE "default"
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Uniques structure for table employee_status
-- ----------------------------
ALTER TABLE "public"."employee_status" ADD UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table employee_status
-- ----------------------------
ALTER TABLE "public"."employee_status" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Table structure for companies
-- ----------------------------
DROP TABLE IF EXISTS "public"."companies";
CREATE TABLE "public"."companies" (
"id" serial4 NOT NULL,
"name" varchar(100) COLLATE "default" NOT NULL,
"tdp" varchar(64) COLLATE "default" NOT NULL,
"email" varchar(64) COLLATE "default" NOT NULL,
"primary_contact" varchar(100) COLLATE "default" NOT NULL,
"phone" varchar(32) COLLATE "default",
"address" text COLLATE "default",
"emp_count" int4 DEFAULT 0
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Uniques structure for table companies
-- ----------------------------
ALTER TABLE "public"."companies" ADD UNIQUE ("tdp", "name");

-- ----------------------------
-- Primary Key structure for table companies
-- ----------------------------
ALTER TABLE "public"."companies" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Table structure for employees
-- ----------------------------
DROP TABLE IF EXISTS "public"."employees";
CREATE TABLE "public"."employees" (
"id" serial4 NOT NULL,
"ktp" int8 NOT NULL,
"first_name" varchar(64) COLLATE "default" NOT NULL,
"last_name" varchar(64) COLLATE "default",
"handphone" varchar(32) COLLATE "default" NOT NULL,
"marital_status" varchar(64) COLLATE "default",
"last_position" varchar(100) COLLATE "default" NOT NULL,
"company_id" int4 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Uniques structure for table employees
-- ----------------------------
ALTER TABLE "public"."employees" ADD UNIQUE ("ktp");

-- ----------------------------
-- Primary Key structure for table employees
-- ----------------------------
ALTER TABLE "public"."employees" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Key structure for table "public"."employees"
-- ----------------------------
-- ALTER TABLE "public"."employees" ADD FOREIGN KEY ("company_id") REFERENCES "public"."companies" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Table structure for company_employees
-- ----------------------------
DROP TABLE IF EXISTS "public"."company_employees";
CREATE TABLE "public"."company_employees" (
"id" serial4 NOT NULL,
"company_id" int4 NOT NULL,
"employee_id" int4 NOT NULL,
"status_id" int4 NOT NULL,
"start_date" date NOT NULL,
"end_date" date,
"position" varchar(100) COLLATE "default" NOT NULL,
"position_desc" varchar(255) COLLATE "default"
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Primary Key structure for table company_employees
-- ----------------------------
ALTER TABLE "public"."company_employees" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Key structure for table "public"."company_employees"
-- ----------------------------
-- ALTER TABLE "public"."company_employees" ADD FOREIGN KEY ("status_id") REFERENCES "public"."employee_status" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ALTER TABLE "public"."company_employees" ADD FOREIGN KEY ("employee_id") REFERENCES "public"."employees" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ALTER TABLE "public"."company_employees" ADD FOREIGN KEY ("company_id") REFERENCES "public"."companies" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Table structure for employee_friends
-- ----------------------------
DROP TABLE IF EXISTS "public"."employee_friends";
CREATE TABLE "public"."employee_friends" (
"id" serial4 NOT NULL,
"employee_ktp" int8 NOT NULL,
"friend_ktp" int8 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Uniques structure for table employee_friends
-- ----------------------------
ALTER TABLE "public"."employee_friends" ADD UNIQUE ("employee_ktp", "friend_ktp");

-- ----------------------------
-- Primary Key structure for table employee_friends
-- ----------------------------
ALTER TABLE "public"."employee_friends" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Key structure for table "public"."employee_friends"
-- -- ----------------------------
-- ALTER TABLE "public"."employee_friends" ADD FOREIGN KEY ("employee_ktp") REFERENCES "public"."employees" ("ktp") ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ALTER TABLE "public"."employee_friends" ADD FOREIGN KEY ("friend_ktp") REFERENCES "public"."employees" ("ktp") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Insert Master Data for table employee_status
-- ----------------------------

INSERT INTO "public"."employee_status"(id, name, description) VALUES (1, 'active', 'active'), (2, 'unactive','not active');
