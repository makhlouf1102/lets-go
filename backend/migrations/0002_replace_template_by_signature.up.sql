DO $$
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name = 'problems' 
          AND column_name = 'template'
    ) THEN
        ALTER TABLE problems RENAME COLUMN template TO signature;
    END IF;
END $$;