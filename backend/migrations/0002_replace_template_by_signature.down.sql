DO $$
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name = 'problems' 
          AND column_name = 'signature'
    ) THEN
        -- Extra safety: only rename if template does NOT already exist
        IF NOT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = 'problems' 
              AND column_name = 'template'
        ) THEN
            ALTER TABLE problems RENAME COLUMN signature TO template;
        END IF;
    END IF;
END $$;