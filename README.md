To run a new postgres container run :
```bash
make postgres_run password="<password>"
```

To start a postgres container run :
```bash
make postgres_start
```

To create a migration run :
```bash
make create_migration name="<name>"
```

To migrate up run :
```bash
make migrate_up password="<password>"
```

To migrate down run :
```bash
make migrate_down password="<password>"
```
