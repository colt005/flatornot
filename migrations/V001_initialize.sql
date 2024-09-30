CREATE TABLE
  public.votes (
    id serial NOT NULL,
    vote integer NULL,
    metadata json NULL,
    created_at date NULL,
    session_id text NULL
  );

ALTER TABLE
  public.votes
ADD
  CONSTRAINT votes_pkey PRIMARY KEY (id)