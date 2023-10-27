### Accessing PostgreSQL inside Docker Container

```
docker exec -it project_postgres_1 psql -U postgres
```

### Kill Port on MacOS

```
sudo lsof -i :<port_number>
sudo kill <PID>
```