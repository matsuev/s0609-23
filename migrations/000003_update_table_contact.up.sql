ALTER TABLE IF EXISTS public.contact DROP COLUMN IF EXISTS full_name;

ALTER TABLE IF EXISTS public.contact
    ADD COLUMN bithdate date;