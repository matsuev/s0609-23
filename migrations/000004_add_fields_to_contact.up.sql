ALTER TABLE IF EXISTS public.contact
   ADD COLUMN IF NOT EXISTS mother_name character variyng(50) NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS public.contact
   ADD COLUMN IF NOT EXISTS father_name character variyng(50) NOT NULL DEFAULT '';
