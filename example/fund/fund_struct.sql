CREATE TABLE "fund" ("code" VARCHAR PRIMARY KEY UNIQUE, "name" VARCHAR, "type" VARCHAR, "create_date" DATE, "create_scale" FLOAT, "latest_scale" FLOAT, "update_date" DATE, "company_code" VARCHAR, "company_name" VARCHAR, "manager_id" VARCHAR, "manager_name" VARCHAR, "manage_exp" FLOAT, "trust_exp" FLOAT, "is_fav" BOOLEAN, "sort_id" INTEGER, "remark" VARCHAR, "tags" VARCHAR);
CREATE TABLE "fund_group" ("fund_code" VARCHAR, "group" VARCHAR);
CREATE TABLE "group" ("name" VARCHAR PRIMARY KEY);
CREATE TABLE "holding_stock" ("fund_code" VARCHAR, "fund_name" VARCHAR, "season" VARCHAR, "date" VARCHAR, "stock_code" VARCHAR, "stock_name" VARCHAR, "stock_value" FLOAT, "stock_amount" FLOAT, "stock_percent" FLOAT);
CREATE TABLE "manager" ("id" TEXT PRIMARY KEY, "name" TEXT, "start_work_date" TEXT, "work_days" INTEGER, "max_growth" FLOAT, "min_growth" FLOAT, "ave_growth" FLOAT, "holding_funds" INTEGER, "education" VARCHAR, "resume" TEXT);
CREATE TABLE "manager_experience" ("manager_id" VARCHAR, "manager_name" VARCHAR, "fund_code" VARCHAR, "fund_name" VARCHAR, "from_date" DATE, "to_date" DATE, "growth" FLOAT);
CREATE TABLE "net_value" ("code" VARCHAR, "date" DATE, "net_value" FLOAT, "total_net_value" FLOAT, "growth" FLOAT);
CREATE VIEW "latest_holding_stock" AS select fund_code, season, max(date) as date from holding_stock group by fund_code;
CREATE UNIQUE INDEX "uniq_holding_stock" ON "holding_stock" ("fund_code","season","stock_code");
CREATE UNIQUE INDEX "uniq_mng_exp" ON "manager_experience" ("manager_id","fund_code","from_date");
CREATE UNIQUE INDEX "uniq_net_value" ON "net_value" ("code","date");