
CREATE TABLE IF NOT EXISTS story_views (
   id                BIGSERIAL PRIMARY KEY,
   viewer_id          BIGINT,
   story_id           BIGINT,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at  TIMESTAMP WITH TIME ZONE,

     FOREIGN KEY (viewer_id) REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (story_id) REFERENCES stories(id) ON DELETE CASCADE
    );

CREATE OR REPLACE FUNCTION update_story_views_count()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE stories
    SET views_count = views_count + 1, updated_at = NOW()
    WHERE id = NEW.story_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER story_views_count_trigger
    AFTER INSERT ON story_views
    FOR EACH ROW
EXECUTE FUNCTION update_story_views_count();


