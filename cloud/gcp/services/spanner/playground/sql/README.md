# spanner-cli SQL playground


use spanner-cli to run SQL queries

```bash
spanner-cli sql --project test-project --instance test-instance --database example-db < create-self-taught-table.sql
```

gcloud

```bash
# create an instance
gcloud spanner instances create test-instance --config=emulator-config --display-name="test-instance" --nodes=1

# create a database
gcloud spanner databases create example-db --instance test-instance  --ddl-file create-self-taught-table.sql

# update ddl
gcloud spanner databases ddl update --ddl-file update-self-taught-table.sql example-db

# drop database
gcloud spanner databases delete example-db --instance test-instance
```