
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

-- Create a trigger function to update views_count in the stories table only once per viewer
CREATE OR REPLACE FUNCTION update_story_views_count()
    RETURNS TRIGGER AS $$
BEGIN
    -- Check if the viewer has already viewed the story
    IF NOT EXISTS (
        SELECT 1 FROM story_views
        WHERE viewer_id = NEW.viewer_id AND story_id = NEW.story_id
    ) THEN
        -- If this is the first time the viewer is viewing the story, increment the views_count
        UPDATE stories
        SET views_count = views_count + 1, updated_at = NOW()
        WHERE id = NEW.story_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to call the function after inserting into story_views
CREATE TRIGGER story_views_count_trigger
    AFTER INSERT ON story_views
    FOR EACH ROW
EXECUTE FUNCTION update_story_views_count();

