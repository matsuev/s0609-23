ALTER TABLE IF EXISTS public.contact
   DROP COLUMN IF EXISTS birthdate;

ALTER TABLE IF EXISTS public.contact
   ADD COLUMN full_name text NOT NULL DEFAULT '';
