
- Crear contenedor para la base de datos (POSTGRES)
docker run -d --name start_wars \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_DB=start_wars \
    -p 5432:5432 \
    -v start_wars:/var/lib/postgresql/data \
    postgres:alpine

- Create function manually
``` sql
CREATE OR REPLACE FUNCTION public.uuid_generate_v4()
RETURNS uuid
LANGUAGE sql
AS $$
    SELECT gen_random_uuid()::uuid;
$$;
``` 