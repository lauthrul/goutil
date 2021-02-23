CREATE TABLE "basic" ("code" VARCHAR PRIMARY KEY UNIQUE, "name" VARCHAR, "create_date" DATE, "scale" INTEGER, "type" SMALLINT, "is_fav" BOOLEAN, "sort_id" INTEGER, "remark" VARCHAR, "tags" VARCHAR, "update_date" DATE);
CREATE TABLE "holding_stock" ("fund_code" VARCHAR, "season" VARCHAR, "date" VARCHAR, "stock_code" VARCHAR, "stock_name" VARCHAR, "stock_value" FLOAT, "stock_amount" FLOAT, "stock_percent" FLOAT);
CREATE TABLE "manager" ("fund_code" VARCHAR, "from_date" DATE, "to_date" DATE, "manager_name" VARCHAR, "manager_company" VARCHAR, "manager_work_date" DATE, "roi" FLOAT);
CREATE TABLE "net_value" ("code" VARCHAR, "date" DATE, "net_value" FLOAT, "total_net_value" FLOAT);
